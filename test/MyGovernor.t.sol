// SPDX-License-Identifier: MIT
pragma solidity >=0.8.19 <0.9.0;

import "forge-std/Test.sol";
import "../src/my_governor.sol";
import "../src/my_time_lock.sol";
import "../src/govtoken.sol";

/// @dev Simple target contract to be governed.
contract Box {
    uint256 public value;
    event ValueChanged(uint256 newValue);
    function set(uint256 newValue) external {
        value = newValue;
        emit ValueChanged(newValue);
    }
}

contract MyGovernorTest is Test {
    MyGovernor public governor;
    MyTimelock public timelock;
    GovToken public token;
    Box public box;

    address public alice = address(0xA11CE);
    address public bob   = address(0xB0B);

    function setUp() public {
        // Deploy token (mints 1,000,000 tokens to this contract)
        token = new GovToken();

        // Give voters enough voting power to pass quorum (4% of 1,000,000 = 40,000)
        token.transfer(alice, 50_000 ether);
        token.transfer(bob,   10_000 ether);

        // Delegation (must occur before the snapshot block)
        vm.prank(alice);
        token.delegate(alice);
        vm.prank(bob);
        token.delegate(bob);

        // Timelock setup
        address[] memory proposers = new address[](0); // none initially; Governor gets role below

        address[] memory executors = new address[](1);
        executors[0] = address(0); // allow anyone to execute

        timelock = new MyTimelock(2 days, proposers, executors);
        governor = new MyGovernor(token, timelock);

        // Grant roles to Governor, then revoke admin from this test
        timelock.grantRole(timelock.PROPOSER_ROLE(), address(governor));
        timelock.grantRole(timelock.EXECUTOR_ROLE(), address(0));
        timelock.revokeRole(timelock.DEFAULT_ADMIN_ROLE(), address(this));

        // Target to call via governance
        box = new Box();
    }

    function test_QuorumIsFourPercent() public {
        // Advance one block so quorum() at clock()-1 references the block with initial mint
        vm.roll(block.number + 1);
        // OpenZeppelin Governor uses timepoint = block number in block mode
        uint256 q = governor.quorum(governor.clock() - 1);
        // 4% of 1,000,000 * 1e18 = 40,000e18
        assertEq(q, 40_000 ether, "quorum should be 4% of total supply");
    }

    function test_DelegationGivesVotingPower() public {
        assertEq(token.getVotes(alice), 50_000 ether, "Alice voting power after delegate");
        assertEq(token.getVotes(bob),   10_000 ether, "Bob voting power after delegate");
    }

    function test_FullProposalLifecycle_PassAndExecute() public {
        // Prepare proposal to set value = 42 on Box
        address[] memory targets = new address[](1);
        targets[0] = address(box);

        uint256[] memory values = new uint256[](1);
        values[0] = 0;

        bytes[] memory calldatas = new bytes[](1);
        calldatas[0] = abi.encodeWithSignature("set(uint256)", 42);

        string memory description = "Set Box value to 42";
        uint256 proposalId = governor.propose(targets, values, calldatas, description);

        // Move past voting delay
        vm.roll(block.number + governor.votingDelay() + 1);

        // Vote: Alice votes For (1), Bob votes For (1)
        vm.prank(alice);
        governor.castVote(proposalId, 1);
        vm.prank(bob);
        governor.castVote(proposalId, 1);

        // Move past voting period
        vm.roll(block.number + governor.votingPeriod() + 1);

        // Queue the proposal
        bytes32 descriptionHash = keccak256(bytes(description));
        governor.queue(targets, values, calldatas, descriptionHash);

        // Wait out the timelock delay
        vm.warp(block.timestamp + timelock.getMinDelay() + 1);

        // Execute
        governor.execute(targets, values, calldatas, descriptionHash);

        // Verify side effect
        assertEq(box.value(), 42, "Box value should be set via governance");
    }
}
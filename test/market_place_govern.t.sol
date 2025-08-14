// SPDX-License-Identifier: MIT
pragma solidity >=0.8.19 <0.9.0;

import "forge-std/Test.sol";
import {MarketplaceGoverned, IToken} from "src/MarketplaceGoverned.sol";
import {MyGovernor} from "src/my_governor.sol";
import {MyTimelock} from "src/my_time_lock.sol";
import {GovToken} from "src/govtoken.sol";

contract MarketplaceGovernedTest is Test {
    GovToken public token;
    MyTimelock public timelock;
    MyGovernor public governor;
    MarketplaceGoverned public marketplace;

    address public deployer = makeAddr("deployer");
    address public alice = makeAddr("alice");
    address public bob = makeAddr("bob");
    address public treasury = makeAddr("treasury");

    uint256 public constant VOTING_DELAY = 1;
    uint256 public constant VOTING_PERIOD = 45818;
    uint256 public constant TIMELOCK_DELAY = 2 days;

    function setUp() public {
        vm.startPrank(deployer);
        token = new GovToken();

    // ✅ RIGHT — declare and initialize memory arrays
        address[] memory proposers = new address[](0);            // none at deploy time
        address[] memory executors = new address[](1);            // anyone can execute
        executors[0] = address(0);

    // Deploy timelock with arrays
        timelock = new MyTimelock(TIMELOCK_DELAY, proposers, executors);

    // (Optional but typical) wire roles post‑deploy
        timelock.grantRole(timelock.PROPOSER_ROLE(), address(governor));
        timelock.grantRole(timelock.EXECUTOR_ROLE(), address(0));
        address deployer = makeAddr("deployer");
        vm.startPrank(deployer);
        marketplace = new MarketplaceGoverned(IToken(address(token)), payable(treasury), 200, address(governor));

        // any owner‑only initialization here
        vm.stopPrank();
        marketplace.transferOwnership(address(timelock));

        // give alice voting power
        token.transfer(alice, 1_000_000 ether);

        vm.stopPrank();
    }

    function _proposeAndExecute(address voter, address target, bytes memory data) internal {
        // Build single‑target proposal arrays
        address[] memory targets = new address[](1);
        targets[0] = target;

        uint256[] memory values = new uint256[](1);
        values[0] = 0;

        bytes[] memory calldatas = new bytes[](1);
        calldatas[0] = data;

        // Governor requires the same description hash on queue/execute as propose
        string memory description = "test proposal";
        bytes32 descHash = keccak256(bytes(description));

        vm.startPrank(voter);
        uint256 proposalId = governor.propose(targets, values, calldatas, description);

        vm.roll(block.number + governor.votingDelay());
        governor.castVote(proposalId, 1); // 0=Against, 1=For, 2=Abstain

        vm.roll(block.number + governor.votingPeriod());
        governor.queue(targets, values, calldatas, descHash);

        vm.warp(block.timestamp + timelock.getMinDelay());
        governor.execute(targets, values, calldatas, descHash);
        vm.stopPrank();
    }

    function testListAndBuyHappyPath() public {
        vm.startPrank(alice);
        uint256 id = marketplace.list(1 ether);
        vm.stopPrank();

        vm.startPrank(bob);
        vm.deal(bob, 1 ether);
        marketplace.buy{value: 1 ether}(id);
        vm.stopPrank();
    }

    function testCannotListWhenPaused() public {
        _proposeAndExecute(alice, address(marketplace), abi.encodeWithSignature("pause()"));
        vm.startPrank(alice);
        vm.expectRevert();
        marketplace.list(1 ether);
        vm.stopPrank();
    }

    function testCannotBuyWhenPaused() public {
        vm.startPrank(alice);
        uint256 id = marketplace.list(1 ether);
        vm.stopPrank();
        _proposeAndExecute(alice, address(marketplace), abi.encodeWithSignature("pause()"));
        vm.startPrank(bob);
        vm.expectRevert();
        marketplace.buy{value: 1 ether}(id);
        vm.stopPrank();
    }

    function testUpdateFeeViaGovernance() public {
        _proposeAndExecute(alice, address(marketplace), abi.encodeWithSignature("setFeeBps(uint16)", 500));
        assertEq(marketplace.feeBps(), 500);
    }

    function testUpdateTreasuryViaGovernance() public {
        address newTreasury = makeAddr("newTreasury");
        _proposeAndExecute(alice, address(marketplace), abi.encodeWithSignature("setTreasury(address)", newTreasury));
        assertEq(marketplace.treasury(), newTreasury);
    }
}

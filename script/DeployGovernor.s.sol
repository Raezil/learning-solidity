// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";

import "../src/govtoken.sol";
import "../src/my_time_lock.sol";
import "../src/my_governor.sol";

contract DeployGovernor is Script {
    event Deployed(string name, address addr);

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        // 1) Deploy governance token
        GovToken token = new GovToken();
        emit Deployed("GovToken", address(token));

        // 2) Prepare timelock parameters
        address[] memory proposers = new address[](0);
        address[] memory executors = new address[](1);
        executors[0] = address(0); // anyone can execute

        // 3) Deploy timelock
        MyTimelock timelock = new MyTimelock(2 days, proposers, executors);
        emit Deployed("MyTimelock", address(timelock));

        // 4) Deploy governor
        MyGovernor governor = new MyGovernor(token, timelock);
        emit Deployed("MyGovernor", address(governor));

        // 5) Grant roles
        bytes32 proposerRole = timelock.PROPOSER_ROLE();
        bytes32 executorRole = timelock.EXECUTOR_ROLE();
        bytes32 adminRole = timelock.DEFAULT_ADMIN_ROLE();

        timelock.grantRole(proposerRole, address(governor));
        timelock.grantRole(executorRole, address(0));
        timelock.revokeRole(adminRole, msg.sender); // remove deployer control

        vm.stopBroadcast();
    }
}
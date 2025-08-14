// SPDX-License-Identifier: MIT
pragma solidity >=0.8.19 <0.9.0;

import "forge-std/Script.sol";
import {GovToken} from "src/govtoken.sol";
import {MyTimelock} from "src/my_time_lock.sol";
import {MyGovernor} from "src/my_governor.sol";
import {MarketplaceGoverned, IToken} from "src/MarketplaceGoverned.sol";
import {Helpers} from "script/Helpers.sol";

contract DeployAll is Script {
    function run() external {
        uint256 pk = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(pk);
        
        console2.log("Deploying with address:", deployer);
        console2.log("Starting deployment...");
        
        vm.startBroadcast(pk);
        
        // 1) Deploy governance token with voting power
        console2.log("1. Deploying GovToken...");
        GovToken token = new GovToken();
        console2.log("   GovToken deployed at:", address(token));
        
        // 2) Deploy Timelock (2 days delay, no initial proposers, anyone can execute)
        console2.log("2. Deploying Timelock...");
        address[] memory proposers = new address[](0); // Empty initially, will add Governor
        address[] memory executors = Helpers.arr(address(0)); // Anyone can execute
        
        MyTimelock timelock = new MyTimelock(2 days, proposers, executors);
        console2.log("   Timelock deployed at:", address(timelock));
        
        // 3) Deploy Governor bound to Timelock
        console2.log("3. Deploying Governor...");
        MyGovernor gov = new MyGovernor(token, timelock);
        console2.log("   Governor deployed at:", address(gov));
        
        // 4) Configure timelock roles
        console2.log("4. Configuring Timelock roles...");
        bytes32 PROPOSER_ROLE = timelock.PROPOSER_ROLE();
        bytes32 EXECUTOR_ROLE = timelock.EXECUTOR_ROLE();
        
        // Grant proposer role to Governor
        timelock.grantRole(PROPOSER_ROLE, address(gov));
        console2.log("   Granted PROPOSER_ROLE to Governor");
        
        // Grant executor role to everyone (address(0) means open execution)
        timelock.grantRole(EXECUTOR_ROLE, address(0));
        console2.log("   Granted EXECUTOR_ROLE to everyone");

        // 5) Deploy Marketplace owned by Timelock
        console2.log("5. Deploying MarketplaceGoverned...");
        MarketplaceGoverned market = new MarketplaceGoverned(
            IToken(address(0)),     // No token integration initially (can be changed later)
            payable(deployer),      // Initial treasury (can be changed via governance)
            200,                    // 2% fee (can be changed via governance)
            address(timelock)       // Timelock as owner (enables DAO control)
        );
        console2.log("   MarketplaceGoverned deployed at:", address(market));
        
        // 6) Delegate voting power to deployer for initial proposals
        console2.log("6. Delegating voting power...");
        token.delegate(deployer);
        console2.log("   Delegated voting power to deployer");
        
        vm.stopBroadcast();
        
        // Print deployment summary
        console2.log("\n=== DEPLOYMENT COMPLETE ===");
        console2.log("GovToken:     ", address(token));
        console2.log("Timelock:     ", address(timelock));
        console2.log("Governor:     ", address(gov));
        console2.log("Marketplace:  ", address(market));
        console2.log("\n=== GOVERNANCE PARAMETERS ===");
        console2.log("Timelock delay:     2 days");
        console2.log("Voting delay:       1 block");
        console2.log("Voting period:      45818 blocks (~1 week)");
        console2.log("Quorum:            4% of total supply");
        console2.log("Marketplace fee:    2% (200 bps)");
        console2.log("Initial treasury:   ", deployer);
        
        console2.log("\n=== NEXT STEPS ===");
        console2.log("1. Create governance proposals to:");
        console2.log("   - Change marketplace fee");
        console2.log("   - Update treasury address");
        console2.log("   - Pause/unpause marketplace");
        console2.log("2. Distribute governance tokens to community");
        console2.log("3. Test governance proposals on testnet first");
    }
}
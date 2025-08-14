// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.19 <0.9.0;
import "forge-std/Script.sol";
import {Marketplace}    from "../src/marketplace.sol";
import {Interactor}     from "../src/interactor.sol";
import {Token}          from "../src/contract.sol";

contract DeployAll is Script {
    function run() external {
        // begin broadcasting to the network
        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));

        // 1. Deploy Token
        Token token = new Token();

        // 2. Deploy Marketplace with the Token address
        Marketplace marketplace = new Marketplace(address(token));

        // 3. Deploy Interactor with the Marketplace address
        Interactor interactor = new Interactor(address(marketplace));
        console.log("Interactor deployed at:", address(interactor));
        // stop broadcasting
        vm.stopBroadcast();
    }
}

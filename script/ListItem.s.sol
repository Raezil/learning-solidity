// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.19 <0.9.0;
import "forge-std/Script.sol";
import {Interactor} from "../src/interactor.sol";

/// @notice Script for a seller to approve the marketplace and list an item
contract ListItemScript is Script {
    function run() external {
        uint256 sellerKey = vm.envUint("PRIVATE_KEY");
        address interactorAddr = vm.envAddress("INTERACTOR_ADDRESS");
        uint256 approveAmount = vm.envUint("APPROVE_AMOUNT");
        uint256 price = 0.5 ether;

        vm.startBroadcast(sellerKey);
        Interactor interactor = Interactor(interactorAddr);

        interactor.approveMarketplace(approveAmount);
        interactor.listItem(price);
        vm.stopBroadcast();
    }
}


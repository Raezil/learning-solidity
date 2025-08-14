// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.19 <0.9.0;
import "forge-std/Script.sol";
import {Interactor} from "../src/interactor.sol";

/// @notice Script for a buyer to purchase a listed item
contract BuyItemScript is Script {
    function run() external {
        uint256 buyerKey = vm.envUint("PRIVATE_KEY");
        address interactorAddr = vm.envAddress("INTERACTOR_ADDRESS");
        uint256 listingId = vm.envUint("LISTING_ID");
        uint256 value = vm.envUint("VALUE");

        vm.startBroadcast(buyerKey);
        Interactor interactor = Interactor(interactorAddr);
        interactor.buyItem{value: value}(listingId);
        vm.stopBroadcast();
    }
}


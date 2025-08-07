// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import {Interactor} from "../src/interactor.sol";

/// @notice Script to check listing details
contract CheckListingScript is Script {
    function run() external {
        address interactorAddr = vm.envAddress("INTERACTOR_ADDRESS");
        uint256 listingId = vm.envUint("LISTING_ID");

        Interactor interactor = Interactor(interactorAddr);
        
        try interactor.getListing(listingId) returns (
            address seller, 
            uint256 price, 
            bool active
        ) {
            console.log("Listing ID:", listingId);
            console.log("Seller:", seller);
            console.log("Price:", price);
            console.log("Active:", active);
            
            if (seller == address(0)) {
                console.log("ERROR: Listing does not exist!");
            } else if (!active) {
                console.log("ERROR: Listing exists but is inactive (already sold)");
            } else {
                console.log("SUCCESS: Listing is valid and active");
            }
        } catch {
            console.log("ERROR: Failed to get listing - may not exist");
        }
    }
}
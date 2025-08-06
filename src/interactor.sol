// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./marketplace.sol";
import "./contract.sol";

contract Interactor {
    Marketplace public marketplace;
    IToken public tokenContract;

    // Initialize with address of the deployed Marketplace
    constructor(address _marketplaceAddr) {
        marketplace = Marketplace(_marketplaceAddr);
        tokenContract = marketplace.token();
    }

    /// @notice Approve the Marketplace to pull tokens on behalf of this contract
    function approveMarketplace(uint256 amount) external {
        marketplace.approveMarketplace(amount);
    }

    /// @notice List an item for sale with a price in tokens
    function listItem(uint256 price) external {
        marketplace.listItem(price);
    }

    /// @notice Buy an item listing by ID, sending ETH equal to your offer
    function buyItem(uint256 listingId) external payable {
        require(msg.value > 0, "Send ETH to buy");
        marketplace.buy{value: msg.value}(listingId);
    }

    /// @notice Fetch listing details: seller, price, active flag
    /// @dev Assumes Marketplace has a getListing(listingId) view function
    function getListing(uint256 listingId) external view returns (address seller, uint256 price, bool active) {
        (seller, price, active) = marketplace.getListing(listingId);
    }

    /// @notice Retrieve on-chain profile for a given user
    function getProfile(address user) external view returns (uint256 balance, address to) {
        (balance, to) = tokenContract.profiles(user);
    }

    /// @notice Read how much allowance a seller has granted to this contract
    function getAllowance(address seller) external view returns (uint256) {
        return marketplace.allowance(seller, address(this));
    }
}
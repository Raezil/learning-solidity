// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./contract.sol";

contract Marketplace {

    struct Listing {
        address seller;
        uint256 price;
        bool active;
    }

    IToken public token;
    uint256 public nextListingId;
    mapping(uint256 => Listing) public listings;
    IToken public immutable token;

    // Allowance mapping to let sellers pre-approve Marketplace
    mapping(address => mapping(address => uint256)) public allowance;

    event Listed(uint256 indexed listingId, address indexed seller, uint256 price);
    event Bought(uint256 indexed listingId, address indexed buyer, uint256 price);
    event Approval(address indexed owner, address indexed spender, uint256 amount);

    constructor(address _token) {
        token = IToken(_token);
    }

    /// @notice Seller calls this to allow the Marketplace to pull tokens on sale.
    function approveMarketplace(uint256 amount) external {
        allowance[msg.sender][address(this)] = amount;
        emit Approval(msg.sender, address(this), amount);
    }

    /// @notice Seller lists an item, sets a price in tokens.
    function listItem(uint256 price) external {
        require(price > 0, "Price>0");
        listings[nextListingId] = Listing(msg.sender, price, true);
        emit Listed(nextListingId, msg.sender, price);
        nextListingId++;
    }

    /// @notice Buyer purchases: transfers tokens and ETH as desired.
    function buy(uint256 listingId) external payable {
        Listing storage item = listings[listingId];
        require(item.active, "Not active");
        require(allowance[item.seller][address(this)] >= item.price,
               "Marketplace not approved for enough tokens");

        // 1) Pull tokens from seller to buyer
        //    (Assumes RaezilToken.transfer does an internal balance update)
        token.transfer(item.price, msg.sender);

        // 2) Send ETH payment back to seller
        //    (You could also price in ETH and pull tokens – swap as you like!)
        payable(item.seller).transfer(msg.value);

        item.active = false;
        emit Bought(listingId, msg.sender, item.price);
    }
    /// @notice Return listing details
    function getListing(uint256 listingId) external view returns (address seller, uint256 price, bool active) {
        Listing memory item = listings[listingId];
        return (item.seller, item.price, item.active);
    }
}


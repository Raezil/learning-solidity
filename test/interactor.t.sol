// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../src/interactor.sol";
import "../src/contract.sol"; // brings IToken interface

/// @notice Mock implementation of the token interface for testing
contract MockToken is IToken {
    struct ProfileData { uint256 balance; address to; }
    mapping(address => ProfileData) private _profiles;

    /// @notice Seed a profile for testing
    function setProfile(address user, uint256 balance, address to) external {
        _profiles[user] = ProfileData(balance, to);
    }

    /// @notice IToken.profiles
    function profiles(address user) external view override returns (uint256, address) {
        ProfileData memory p = _profiles[user];
        return (p.balance, p.to);
    }

    /// @notice IToken.transfer stub
    function transfer(uint256 amount, address to) external override {
        // no-op
    }
}

/// @notice Mock implementation of the Marketplace contract
contract MockMarketplace {
    IToken public token;

    // Track forwarded parameters
    uint256 public lastApproveAmount;
    uint256 public lastListPrice;
    uint256 public lastBuyId;
    uint256 public lastBuyValue;

    // Storage for mocked data
    struct Listing { address seller; uint256 price; bool active; }
    mapping(uint256 => Listing) private _listings;
    mapping(address => mapping(address => uint256)) public allowances;

    constructor(IToken tokenAddr) {
        token = tokenAddr;
    }

    function approveMarketplace(uint256 amount) external {
        lastApproveAmount = amount;
    }

    function listItem(uint256 price) external {
        lastListPrice = price;
    }

    function buy(uint256 listingId) external payable {
        require(msg.value > 0, "Send ETH");
        lastBuyId = listingId;
        lastBuyValue = msg.value;
    }

    /// @notice Seed a listing
    function setListing(uint256 id, address seller, uint256 price, bool active) external {
        _listings[id] = Listing(seller, price, active);
    }

    /// @notice getListing for Interactor
    function getListing(uint256 id) external view returns (address seller, uint256 price, bool active) {
        Listing memory l = _listings[id];
        return (l.seller, l.price, l.active);
    }

    /// @notice Seed allowance data
    function setAllowance(address seller, address interactor, uint256 amount) external {
        allowances[seller][interactor] = amount;
    }

    /// @notice allowance for Interactor
    function allowance(address seller, address interactor) external view returns (uint256) {
        return allowances[seller][interactor];
    }
}

contract InteractorTest is Test {
    MockToken      token;
    MockMarketplace marketplace;
    Interactor      interactor;
    address         alice = address(0x123);
    address         bob   = address(0x999);

    function setUp() public {
        token       = new MockToken();
        marketplace = new MockMarketplace(IToken(token));
        interactor  = new Interactor(address(marketplace));

        // Fund this contract with ETH
        vm.deal(address(this), 10 ether);
    }

    function testApproveMarketplace() public {
        uint256 amount = 1e18;
        interactor.approveMarketplace(amount);
        assertEq(marketplace.lastApproveAmount(), amount, "approveMarketplace forwarded wrong amount");
    }

    function testListItem() public {
        uint256 price = 500;
        interactor.listItem(price);
        assertEq(marketplace.lastListPrice(), price, "listItem forwarded wrong price");
    }

    function testBuyItem() public {
        uint256 id   = 42;
        uint256 cost = 2 ether;
        interactor.buyItem{ value: cost }(id);
        assertEq(marketplace.lastBuyId(), id,   "buyItem forwarded wrong listingId");
        assertEq(marketplace.lastBuyValue(), cost, "buyItem forwarded wrong ETH amount");
    }

    function testBuyItemRevertsWithoutEth() public {
        vm.expectRevert(bytes("Send ETH to buy"));
        interactor.buyItem(1);
    }

    function testGetListingPopulated() public {
        uint256 id = 7;
        marketplace.setListing(id, alice, 750, true);
        (address seller, uint256 price, bool active) = interactor.getListing(id);
        assertEq(seller, alice,   "getListing returned wrong seller");
        assertEq(price, 750,      "getListing returned wrong price");
        assertTrue(active,         "getListing returned inactive flag");
    }

    function testGetListingDefault() public {
        (address seller, uint256 price, bool active) = interactor.getListing(999);
        assertEq(seller, address(0),   "default seller should be zero");
        assertEq(price, 0,             "default price should be zero");
        assertFalse(active,            "default active should be false");
    }

    function testGetProfile() public {
        uint256 bal = 1234;
        address to  = address(0x456);
        token.setProfile(alice, bal, to);
        (uint256 balance, address receivable) = interactor.getProfile(alice);
        assertEq(balance, bal,   "getProfile returned wrong balance");
        assertEq(receivable, to, "getProfile returned wrong address");
    }

    function testGetAllowancePopulated() public {
        uint256 allowed = 888;
        marketplace.setAllowance(alice, address(interactor), allowed);
        uint256 got = interactor.getAllowance(alice);
        assertEq(got, allowed, "getAllowance returned wrong allowance");
    }

    function testGetAllowanceDefault() public {
        uint256 got = interactor.getAllowance(bob);
        assertEq(got, 0, "default allowance should be zero");
    }
}

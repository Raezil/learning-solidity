// SPDX-License-Identifier: MIT
pragma solidity >=0.8.19 <0.9.0;

import "forge-std/Test.sol";
import {MarketplaceGoverned, IToken} from "src/MarketplaceGoverned.sol";

contract MockToken is IToken {
    mapping(address => uint256) private _bal;
    mapping(address => address) private _to;

    function transfer(uint256 /* amount */, address /* to_ */) external override {
        // no-op for tests
    }

    function setProfile(address user, uint256 bal, address to_) external {
        _bal[user] = bal;
        _to[user] = to_;
    }

    function profiles(address user) external view override returns (uint256 balance, address to_) {
        return (_bal[user], _to[user]);
    }
}

contract MarketplaceGovernedTest is Test {
    // Events mirrored from the contract (for expectEmit)
    event Listed(uint256 indexed id, address indexed seller, uint256 price);
    event Bought(uint256 indexed id, address indexed buyer, uint256 price, uint256 fee);
    event FeeUpdated(uint16 bps);
    event TreasuryUpdated(address indexed newTreasury);

    MockToken token;
    MarketplaceGoverned mkt;

    address owner    = makeAddr("owner");     // typically a Timelock
    address treasury = makeAddr("treasury");
    address seller   = makeAddr("seller");
    address buyer    = makeAddr("buyer");
    address rando    = makeAddr("rando");

    uint16 constant FEE_BPS = 200; // 2%

    function setUp() public {
        token = new MockToken();
        vm.prank(owner);
        mkt = new MarketplaceGoverned(IToken(address(token)), payable(treasury), FEE_BPS, owner);
    }

    function testInitialState() public {
        assertEq(mkt.owner(), owner);
        assertEq(address(mkt.treasury()), treasury);
        assertEq(mkt.feeBps(), FEE_BPS);
        assertEq(mkt.nextListingId(), 0);
    }

    function testSettersOnlyOwner() public {
        // Non-owner cannot set
        vm.prank(rando);
        vm.expectRevert(); // OwnableUnauthorizedAccount in OZ 5.x
        mkt.setFeeBps(100);

        vm.prank(rando);
        vm.expectRevert();
        mkt.setTreasury(payable(rando));

        // Owner can set; events are emitted
        vm.prank(owner);
        vm.expectEmit(true, false, false, true);
        emit FeeUpdated(500);
        mkt.setFeeBps(500);
        assertEq(mkt.feeBps(), 500);

        vm.prank(owner);
        vm.expectEmit(true, true, false, true);
        emit TreasuryUpdated(rando);
        mkt.setTreasury(payable(rando));
        assertEq(address(mkt.treasury()), rando);
    }

    function testSettersValidation() public {
        // fee > 10%
        vm.prank(owner);
        vm.expectRevert(bytes("fee>10%"));
        mkt.setFeeBps(1001);

        // treasury zero
        vm.prank(owner);
        vm.expectRevert(bytes("treasury=0"));
        mkt.setTreasury(payable(address(0)));
    }

    function testPauseUnpause() public {
        // non-owner cannot pause
        vm.prank(rando);
        vm.expectRevert();
        mkt.pause();

        // owner pauses
        vm.prank(owner);
        mkt.pause();

        // list should revert while paused
        vm.prank(seller);
        vm.expectRevert(); // EnforcedPause in OZ 5.x
        mkt.list(1 ether);

        // owner unpauses
        vm.prank(owner);
        mkt.unpause();
        // now list is fine
        vm.prank(seller);
        uint256 id = mkt.list(1 ether);
        assertEq(id, 1);
    }

    function testListRevertsOnZeroPrice() public {
        vm.prank(seller);
        vm.expectRevert(bytes("price=0"));
        mkt.list(0);
    }

    function testListEmitsAndStores() public {
        vm.prank(seller);
        vm.expectEmit(true, true, false, true);
        emit Listed(1, seller, 1 ether);
        uint256 id = mkt.list(1 ether);
        assertEq(id, 1);

        (address s, uint256 p, bool active) = mkt.getListing(id);
        assertEq(s, seller);
        assertEq(p, 1 ether);
        assertTrue(active);
        assertEq(mkt.nextListingId(), 1);
    }

    function testBuyHappyPath_DistributesFeeAndPayout() public {
        // Arrange: seller lists, buyer funded
        vm.prank(seller);
        uint256 id = mkt.list(5 ether);
        vm.deal(buyer, 10 ether);
        uint256 t0Treasury = treasury.balance;
        uint256 t0Seller   = seller.balance;
        uint256 feeBps = mkt.feeBps();
        uint256 fee = (5 ether * feeBps) / 10_000;
        uint256 payout = 5 ether - fee;

        // Expect event
        vm.expectEmit(true, true, false, true);
        emit Bought(id, buyer, 5 ether, fee);

        // Act
        vm.prank(buyer);
        mkt.buy{value: 5 ether}(id);

        // Assert balances
        assertEq(treasury.balance, t0Treasury + fee);
        assertEq(seller.balance,   t0Seller   + payout);

        // Listing is now inactive
        (, , bool active) = mkt.getListing(id);
        assertFalse(active);

        // Cannot buy again
        vm.prank(buyer);
        vm.expectRevert(bytes("inactive"));
        mkt.buy{value: 5 ether}(id);
    }

    function testBuyRevertsOnBadPaymentOrInactive() public {
        // nonexistent listing (inactive)
        vm.prank(buyer);
        vm.expectRevert(); // allow any revert, any depth
        mkt.buy{value: 1 ether}(999);

        // correct id but wrong msg.value
        vm.prank(seller);
        uint256 id = mkt.list(2 ether);

        vm.prank(buyer);
        vm.expectRevert(); // allow any revert, any depth
        mkt.buy{value: 1 ether}(id);
    }



    function testConstructorValidation() public {
        vm.expectRevert(bytes("treasury=0"));
        new MarketplaceGoverned(IToken(address(token)), payable(address(0)), 100, owner);

        vm.expectRevert(bytes("fee>10%"));
        new MarketplaceGoverned(IToken(address(token)), payable(treasury), 1001, owner);
    }
}

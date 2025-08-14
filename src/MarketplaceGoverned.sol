// SPDX-License-Identifier: MIT
pragma solidity >=0.8.19 <0.9.0;

/// Minimal token interface your existing Marketplace expects
interface IToken {
    function transfer(uint256 amount, address to) external;
    function profiles(address user) external view returns (uint256 balance, address to);
}

import "@openzeppelin/contracts/access/Ownable2Step.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/// @notice A simple DAO-governed marketplace with fees & pause control.
/// If you already have a Marketplace contract, you can adapt this pattern:
/// - add `Ownable2Step` and transfer ownership to the Timelock
/// - expose the same admin functions (setFeeBps, setTreasury, pause/unpause)
contract MarketplaceGoverned is Ownable2Step, Pausable, ReentrancyGuard {
    struct Listing { 
        address seller; 
        uint256 price; 
        bool active; 
    }

    IToken public token;                  // off-chain credit token used by your app
    address payable public treasury;      // fee receiver
    uint16 public feeBps;                 // e.g., 200 = 2%

    uint256 public nextListingId;
    mapping(uint256 => Listing) public listings;

    event Listed(uint256 indexed id, address indexed seller, uint256 price);
    event Bought(uint256 indexed id, address indexed buyer, uint256 price, uint256 fee);
    event FeeUpdated(uint16 bps);
    event TreasuryUpdated(address indexed newTreasury);

    constructor(IToken _token, address payable _treasury, uint16 _feeBps, address _owner) Ownable(msg.sender) {
        require(_treasury != address(0), "treasury=0");
        require(_feeBps <= 1000, "fee>10%");
        token = _token;
        treasury = _treasury;
        feeBps = _feeBps;
        
        // Transfer ownership to the specified owner (typically the Timelock)
        if (_owner != msg.sender) {
            _transferOwnership(_owner);
        }
    }

    // ------------------ DAO-controlled admin ------------------
    function setFeeBps(uint16 _feeBps) external onlyOwner { 
        require(_feeBps <= 1000, "fee>10%"); 
        feeBps = _feeBps; 
        emit FeeUpdated(_feeBps); 
    }
    
    function setTreasury(address payable _treasury) external onlyOwner { 
        require(_treasury != address(0), "treasury=0"); 
        treasury = _treasury; 
        emit TreasuryUpdated(_treasury); 
    }
    
    function pause() external onlyOwner { 
        _pause(); 
    }
    
    function unpause() external onlyOwner { 
        _unpause(); 
    }

    // ------------------ Core listing flows ------------------
    function list(uint256 price) external whenNotPaused returns (uint256 id) {
        require(price > 0, "price=0");
        id = ++nextListingId;
        listings[id] = Listing({ seller: msg.sender, price: price, active: true });
        emit Listed(id, msg.sender, price);
    }

    /// @notice Buyer pays in ETH; seller receives ETH minus DAO fee.
    /// You can swap to paying with your token if desired.
    function buy(uint256 id) external payable whenNotPaused nonReentrant {
        Listing storage item = listings[id];
        require(item.active, "inactive");
        require(msg.value == item.price, "bad payment");

        // Calculate fee
        uint256 fee = (msg.value * feeBps) / 10_000;
        uint256 payout = msg.value - fee;

        // Effects first
        item.active = false;
        emit Bought(id, msg.sender, item.price, fee);

        // Interactions
        if (fee > 0) {
            (bool success1,) = treasury.call{value: fee}("");
            require(success1, "fee xfer failed");
        }
        
        (bool success2,) = payable(item.seller).call{value: payout}("");
        require(success2, "seller xfer failed");

        // OPTIONAL: reflect purchase into your off-chain token profile
        // Example: credit 1 unit to buyer's profile destination address
        // (uint256 bal, address to) = token.profiles(msg.sender);
        // token.transfer(1, to);
    }

    function getListing(uint256 id) external view returns (address seller, uint256 price, bool active) {
        Listing memory item = listings[id];
        return (item.seller, item.price, item.active);
    }

    receive() external payable {}
}
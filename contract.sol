// SPDX-License-Identifier: MIT
pragma solidity >=0.4.0 <0.9.0;

interface IToken {
    // mirror only what the Marketplace needs:
    function transfer(uint256 amount, address to) external;
    function profiles(address user) external view returns (uint256 balance, address to);
}

contract Token {
    event TokenWasTransfered(address indexed from, address indexed to, uint256 amount);
    error NotEnoughFunds(uint256 requested, uint256 available);

    struct Profile {
        uint256 balance;
        address to;
    }

    mapping(address => Profile) public profiles;

    // Transfer tokens from sender to recipient using Profile
    function transfer(uint256 amount, address to) public {
        uint256 senderBalance = getBalance();

        if (senderBalance < amount) {
            revert NotEnoughFunds({ requested: amount, available: senderBalance });
        }

        // Deduct the amount from sender's profile balance
        deduct(amount);

        // Add the amount to the recipient's profile
        deposit(to, amount);

        // Emit transfer event
        emit TokenWasTransfered(msg.sender, to, amount);
    }

    function deduct(uint256 amount) public {
        profiles[msg.sender].balance = getBalance() - amount;
    }

    // Get the balance of the sender's profile
    function getBalance() public view returns (uint256) {
        return profiles[msg.sender].balance;
    }

    // Deposit tokens to a given address's profile
    function deposit(address to, uint256 amount) public {
        profiles[to].balance += amount;
    }

    // Create or update a profile
    function createOrUpdateProfile(address user, uint256 balance, address to) public {
        profiles[user] = Profile({ balance: balance, to: to });
    }

    // Get profile information
    function getProfile(address user) public view returns (uint256 balance, address to) {
        Profile storage profile = profiles[user];
        return (profile.balance, profile.to);
    }

    // Update the profile balance
    function updateProfileBalance(address user, uint256 newBalance) public {
        profiles[user].balance = newBalance;
    }

    // Update the profile 'to' address
    function updateProfileTo(address user, address newTo) public {
        profiles[user].to = newTo;
    }
}

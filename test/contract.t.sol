// SPDX-License-Identifier: MIT
pragma solidity >=0.4.0 <0.9.0;

import "forge-std/Test.sol";
import "../src/contract.sol";

contract TokenTest is Test {
    Token public token;
    
    address public alice = address(0x1);
    address public bob = address(0x2);
    address public charlie = address(0x3);
    
    event TokenWasTransfered(address indexed from, address indexed to, uint256 amount);
    
    function setUp() public {
        token = new Token();
    }
    
    // Test profile creation and updates
    function testCreateProfile() public {
        token.createOrUpdateProfile(alice, 1000, bob);
        
        (uint256 balance, address to) = token.getProfile(alice);
        assertEq(balance, 1000);
        assertEq(to, bob);
    }
    
    function testUpdateProfile() public {
        // Create initial profile
        token.createOrUpdateProfile(alice, 1000, bob);
        
        // Update the profile
        token.createOrUpdateProfile(alice, 2000, charlie);
        
        (uint256 balance, address to) = token.getProfile(alice);
        assertEq(balance, 2000);
        assertEq(to, charlie);
    }
    
    function testUpdateProfileBalance() public {
        token.createOrUpdateProfile(alice, 1000, bob);
        token.updateProfileBalance(alice, 1500);
        
        (uint256 balance,) = token.getProfile(alice);
        assertEq(balance, 1500);
    }
    
    function testUpdateProfileTo() public {
        token.createOrUpdateProfile(alice, 1000, bob);
        token.updateProfileTo(alice, charlie);
        
        (, address to) = token.getProfile(alice);
        assertEq(to, charlie);
    }
    
    // Test balance operations
    function testGetBalanceForNonExistentProfile() public {
        vm.prank(alice);
        uint256 balance = token.getBalance();
        assertEq(balance, 0);
    }
    
    function testGetBalanceForExistingProfile() public {
        token.createOrUpdateProfile(alice, 1000, bob);
        
        vm.prank(alice);
        uint256 balance = token.getBalance();
        assertEq(balance, 1000);
    }
    
    function testDeposit() public {
        token.deposit(alice, 500);
        
        (uint256 balance,) = token.getProfile(alice);
        assertEq(balance, 500);
    }
    
    function testMultipleDeposits() public {
        token.deposit(alice, 500);
        token.deposit(alice, 300);
        
        (uint256 balance,) = token.getProfile(alice);
        assertEq(balance, 800);
    }
    
    function testDeduct() public {
        token.createOrUpdateProfile(alice, 1000, bob);
        
        vm.prank(alice);
        token.deduct(200);
        
        (uint256 balance,) = token.getProfile(alice);
        assertEq(balance, 800);
    }
    
    // Test transfer functionality
    function testSuccessfulTransfer() public {
        // Setup: Give alice some tokens
        token.createOrUpdateProfile(alice, 1000, address(0));
        
        // Expect the transfer event
        vm.expectEmit(true, true, false, true);
        emit TokenWasTransfered(alice, bob, 300);
        
        // Execute transfer
        vm.prank(alice);
        token.transfer(300, bob);
        
        // Verify balances
        (uint256 aliceBalance,) = token.getProfile(alice);
        (uint256 bobBalance,) = token.getProfile(bob);
        
        assertEq(aliceBalance, 700);
        assertEq(bobBalance, 300);
    }
    
    function testTransferEntireBalance() public {
        token.createOrUpdateProfile(alice, 1000, address(0));
        
        vm.prank(alice);
        token.transfer(1000, bob);
        
        (uint256 aliceBalance,) = token.getProfile(alice);
        (uint256 bobBalance,) = token.getProfile(bob);
        
        assertEq(aliceBalance, 0);
        assertEq(bobBalance, 1000);
    }
    
    function testTransferToSelf() public {
        token.createOrUpdateProfile(alice, 1000, address(0));
        
        vm.prank(alice);
        token.transfer(500, alice);
        
        // Balance should remain the same
        (uint256 balance,) = token.getProfile(alice);
        assertEq(balance, 1000);
    }
    
    function testTransferInsufficientFunds() public {
        token.createOrUpdateProfile(alice, 100, address(0));
        
        vm.expectRevert(
            abi.encodeWithSelector(Token.NotEnoughFunds.selector, 500, 100)
        );
        
        vm.prank(alice);
        token.transfer(500, bob);
    }
    
    function testTransferFromZeroBalance() public {
        vm.expectRevert(
            abi.encodeWithSelector(Token.NotEnoughFunds.selector, 100, 0)
        );
        
        vm.prank(alice);
        token.transfer(100, bob);
    }
    
    function testTransferZeroAmount() public {
        token.createOrUpdateProfile(alice, 1000, address(0));
        
        vm.prank(alice);
        token.transfer(0, bob);
        
        // Balances should remain unchanged
        (uint256 aliceBalance,) = token.getProfile(alice);
        (uint256 bobBalance,) = token.getProfile(bob);
        
        assertEq(aliceBalance, 1000);
        assertEq(bobBalance, 0);
    }
    
    // Test multiple transfers
    function testMultipleTransfers() public {
        token.createOrUpdateProfile(alice, 1000, address(0));
        
        vm.startPrank(alice);
        token.transfer(200, bob);
        token.transfer(300, charlie);
        vm.stopPrank();
        
        (uint256 aliceBalance,) = token.getProfile(alice);
        (uint256 bobBalance,) = token.getProfile(bob);
        (uint256 charlieBalance,) = token.getProfile(charlie);
        
        assertEq(aliceBalance, 500);
        assertEq(bobBalance, 200);
        assertEq(charlieBalance, 300);
    }
    
    // Test profiles mapping directly
    function testProfilesMapping() public {
        token.createOrUpdateProfile(alice, 1000, bob);
        
        (uint256 balance, address to) = token.profiles(alice);
        assertEq(balance, 1000);
        assertEq(to, bob);
    }
    
    // Fuzz testing
    function testFuzzDeposit(address user, uint256 amount) public {
        vm.assume(user != address(0));
        vm.assume(amount <= type(uint128).max); // Prevent overflow
        
        token.deposit(user, amount);
        
        (uint256 balance,) = token.getProfile(user);
        assertEq(balance, amount);
    }
    
    function testFuzzTransfer(uint256 initialBalance, uint256 transferAmount) public {
        vm.assume(initialBalance <= type(uint128).max);
        vm.assume(transferAmount <= initialBalance);
        
        token.createOrUpdateProfile(alice, initialBalance, address(0));
        
        vm.prank(alice);
        token.transfer(transferAmount, bob);
        
        (uint256 aliceBalance,) = token.getProfile(alice);
        (uint256 bobBalance,) = token.getProfile(bob);
        
        assertEq(aliceBalance, initialBalance - transferAmount);
        assertEq(bobBalance, transferAmount);
    }
    
    function testFuzzTransferInsufficientFunds(uint256 balance, uint256 transferAmount) public {
        vm.assume(balance < transferAmount);
        vm.assume(transferAmount > 0);
        vm.assume(balance <= type(uint128).max);
        
        token.createOrUpdateProfile(alice, balance, address(0));
        
        vm.expectRevert(
            abi.encodeWithSelector(Token.NotEnoughFunds.selector, transferAmount, balance)
        );
        
        vm.prank(alice);
        token.transfer(transferAmount, bob);
    }
    
    // Edge cases
    function testMaxValueTransfer() public {
        uint256 maxValue = type(uint256).max;
        token.createOrUpdateProfile(alice, maxValue, address(0));
        
        vm.prank(alice);
        token.transfer(maxValue, bob);
        
        (uint256 aliceBalance,) = token.getProfile(alice);
        (uint256 bobBalance,) = token.getProfile(bob);
        
        assertEq(aliceBalance, 0);
        assertEq(bobBalance, maxValue);
    }
    
    function testTransferToZeroAddress() public {
        token.createOrUpdateProfile(alice, 1000, address(0));
        
        vm.prank(alice);
        token.transfer(500, address(0));
        
        (uint256 aliceBalance,) = token.getProfile(alice);
        (uint256 zeroBalance,) = token.getProfile(address(0));
        
        assertEq(aliceBalance, 500);
        assertEq(zeroBalance, 500);
    }
    
    // Test event emissions
    function testTransferEventEmission() public {
        token.createOrUpdateProfile(alice, 1000, address(0));
        
        vm.expectEmit(true, true, false, true);
        emit TokenWasTransfered(alice, bob, 250);
        
        vm.prank(alice);
        token.transfer(250, bob);
    }
    
    // Integration test
    function testComplexScenario() public {
        // Setup multiple users with different balances
        token.createOrUpdateProfile(alice, 1000, address(0));
        token.createOrUpdateProfile(bob, 500, address(0));
        
        // Alice transfers to Bob
        vm.prank(alice);
        token.transfer(200, bob);
        
        // Bob transfers to Charlie
        vm.prank(bob);
        token.transfer(300, charlie);
        
        // Charlie transfers back to Alice
        vm.prank(charlie);
        token.transfer(150, alice);
        
        // Verify final balances
        (uint256 aliceBalance,) = token.getProfile(alice);
        (uint256 bobBalance,) = token.getProfile(bob);
        (uint256 charlieBalance,) = token.getProfile(charlie);
        
        assertEq(aliceBalance, 950);  // 1000 - 200 + 150
        assertEq(bobBalance, 400);    // 500 + 200 - 300
        assertEq(charlieBalance, 150); // 0 + 300 - 150
    }
}
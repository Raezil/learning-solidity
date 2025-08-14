// SPDX-License-Identifier: MIT
pragma solidity >=0.8.19 <0.9.0;

library Helpers {
    /// @notice Helper function to create a single-element address array
    function arr(address single) internal pure returns (address[] memory) {
        address[] memory result = new address[](1);
        result[0] = single;
        return result;
    }
    
    /// @notice Helper function to create a two-element address array
    function arr(address first, address second) internal pure returns (address[] memory) {
        address[] memory result = new address[](2);
        result[0] = first;
        result[1] = second;
        return result;
    }
    
    /// @notice Helper function to create an empty address array
    function emptyAddressArray() internal pure returns (address[] memory) {
        return new address[](0);
    }
}
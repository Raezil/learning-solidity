// SPDX-License-Identifier: MIT
pragma solidity >=0.8.19 <0.9.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol";

contract GovToken is ERC20, ERC20Permit, ERC20Votes {
    constructor()
        ERC20("GovToken", "GOV")
        ERC20Permit("GovToken")
    {
        _mint(msg.sender, 1_000_000 ether);
    }

    // Override nonces function for multiple inheritance
    function nonces(address owner)
        public
        view
        override(Nonces, ERC20Permit)
        returns (uint256)
    {
        return super.nonces(owner);
    }
// Override _burn for ERC20Votes
// Override _update for ERC20Votes (required in newer OpenZeppelin versions)
    function _update(
        address from,
        address to,
        uint256 value
    ) internal override(ERC20, ERC20Votes) {
        super._update(from, to, value);
    }
}
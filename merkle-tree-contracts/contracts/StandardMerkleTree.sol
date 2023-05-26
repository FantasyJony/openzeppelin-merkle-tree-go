// SPDX-License-Identifier: MIT

pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

contract StandardMerkleTree {
    
    function verify(bytes32[] calldata proof, bytes32 root, address account, uint256 amount) external pure returns (bool) {
        return MerkleProof.verifyCalldata(proof, root, _leaf(account, amount));
    }

    function multiProofVerify(
        bytes32[] calldata proof,
        bool[] calldata proofFlags,
        bytes32 root,
        address[] calldata accounts,
        uint256[] calldata amounts
    ) external pure returns (bool) {
        bytes32[] memory leaves = new bytes32[](accounts.length);
        for(uint256 i = 0; i < accounts.length; i++) {
            leaves[i] = _leaf(accounts[i], amounts[i]);
        }
        return MerkleProof.multiProofVerifyCalldata(proof, proofFlags, root, leaves);
    }

    function _leaf(address account, uint256 amount) internal pure returns (bytes32) {
        return keccak256(bytes.concat(keccak256(abi.encode(account, amount))));
    }
}

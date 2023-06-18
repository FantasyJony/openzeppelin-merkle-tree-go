import { time, loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import { ethers } from "hardhat";
import { StandardMerkleTree } from "@openzeppelin/merkle-tree";

describe("StandardMerkleTree", function () {

  async function deployFixture() {
    const [owner, otherAccount] = await ethers.getSigners();
    const StandardMerkleTree = await ethers.getContractFactory("StandardMerkleTree");
    const standardMerkleTree = await StandardMerkleTree.deploy();

    return { standardMerkleTree,  owner, otherAccount };
  }

  describe("Merkle Tree", function () {

    it("verify", async function () {

      const { standardMerkleTree } = await loadFixture(deployFixture);

      const account = "0x1111111111111111111111111111111111111111";
      const amount = "5000000000000000000";
      const root = "0xd4dee0beab2d53f2cc83e567171bd2820e49898130a22622b10ead383e90bd77";
      const proof = [
        '0xb92c48e9d7abe27fd8dfd6b5dfdbfb1c9a463f80c712b66f3a5180a090cccafc',
      ];

      expect(await standardMerkleTree.verify(proof,root,account,amount)).to.equal(true);
    });

    it("verifyMultiProof", async function () {
      const { standardMerkleTree } = await loadFixture(deployFixture);

      const proof = [
        "0x8610c4ddba34d72ee1dabba4f1a813087579d4c6579c495c101530432969efa7",
      ];
      
      const proofFlags = [true, false];

      const root = "0xcef9852531f2476330b76131d5de322f616540e5668b46383dd26f96c50d8861";

      const accounts = [
        "0x1111111111111111111111111111111111111111",
        "0x2222222222222222222222222222222222222222"
      ];

      const amounts = [
        "5000000000000000000",
        "2500000000000000000"
      ];

      expect(await standardMerkleTree.multiProofVerify(proof, proofFlags ,root, accounts, amounts)).to.equal(true);
    });

    it("array",  async function () {
      // const { standardMerkleTree } = await loadFixture(deployFixture)
      const values = [
        [
          "0x1111111111111111111111111111111111111111",
           "1",
           "2",
           "3",
          //  ["1","2","3"]
          ],
        [
          "0x1111111111111111111111111111111111111111", 
          "0",
          "2",
          "1",
          // ["0","1","2"]
        ],
      ];
      const tree = StandardMerkleTree.of(values, [
        "address",
        "uint8",
        "uint88",
        "uint128",
        // "uint8[]"
      ])
      console.log('Merkle Root:', tree.root);
      console.log(tree.dump())
    })

  });
});

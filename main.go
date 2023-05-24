package main

import (
	"fmt"
	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {

	values := [][]interface{}{
		{
			smt.SolAddress("0x1111111111111111111111111111111111111111"),
			smt.SolNumber("5000000000000000000"),
		},
		{
			smt.SolAddress("0x2222222222222222222222222222222222222222"),
			smt.SolNumber("2500000000000000000"),
		},
	}

	leafEncodings := []string{
		smt.SOL_ADDRESS,
		smt.SOL_UINT256,
	}

	// (1) make tree
	tree, err := smt.Of(values, leafEncodings)
	handleError(err)
	fmt.Println("(1) Merkle Root: ", hexutil.Encode(tree.GetRoot()))

	// (2) dump
	dump, err := tree.TreeMarshal()
	handleError(err)
	fmt.Println("(2) Dump: ", string(dump))

	// (3) dump -> tree
	newTree, err := smt.TreeUnmarshal(dump)
	handleError(err)

	entries := newTree.Entries()
	for k, v := range entries {
		if v.Value[0] == smt.SolAddress("0x1111111111111111111111111111111111111111") {
			proof, err := newTree.GetProofWithIndex(k)
			handleError(err)
			fmt.Println("(3) ValueIndex:", v.ValueIndex, " TreeIndex:", v.TreeIndex, " Hash:", hexutil.Encode(v.Hash), " Value:", v.Value)
			for pk, pv := range proof {
				fmt.Println("(3) Proof k:", pk, " v:", hexutil.Encode(pv))
			}
		}
	}

	// (4) dump leaf proof
	leafProofDump, err := tree.TreeProofMarshal()
	handleError(err)
	fmt.Println("(4) Dump Leaf Proof: ", string(leafProofDump))

	// (5) verify proof
	firstProof, err := newTree.GetProofWithIndex(0)
	handleError(err)
	firstValue := entries[0].Value

	verified, err := newTree.Verify(firstProof, firstValue)
	handleError(err)
	fmt.Println("(5) Verify:", verified)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

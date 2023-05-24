# GO Openzeppelin Standard Tree

A Go library to generate merkle trees and merkle proofs.

## @openzeppelin/merkle-tree

fork from [@openzeppelin/merkle-tree](https://github.com/OpenZeppelin/merkle-tree)

contracts doc [MerkleProof](https://docs.openzeppelin.com/contracts/4.x/api/utils#MerkleProof)

## Installation
```
go get github.com/FantasyJony/openzeppelin-merkle-tree-go@v1.0.3
go mod tidy
```

## Example
```go
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
	tree, err := smt.Of(leafEncodings, values)
	handleError(err)
	fmt.Println("(1) Merkle Root: ", hexutil.Encode(tree.GetRoot()))

	// (2) dump
	dump, err := tree.TreeMarshal()
	handleError(err)
	fmt.Println("(2) Dump: ", string(dump))

	// (3) json -> tree
	newTree, err := smt.TreeUnmarshal(dump)
	handleError(err)
	for k, v := range newTree.Leaves {
		if v.Value[0] == smt.SolAddress("0x1111111111111111111111111111111111111111") {
			proof, err := newTree.GetProofWithIndex(k)
			handleError(err)
			fmt.Println("(3) ValueIndex:", v.ValueIndex, " TreeIndex:", v.TreeIndex, " Hash:", hexutil.Encode(v.Hash), " Value:", v.Value)
			
			for pk, pv := range proof {
				fmt.Println("(3) Proof k:", pk, " v:", hexutil.Encode(pv))
			}
		}
	}

	// (4) verify
	firstProof, err := newTree.GetProofWithIndex(0)
	handleError(err)
	firstValue := newTree.Leaves[0].Value

	result, err := newTree.Verify(firstProof, firstValue)
	handleError(err)
	fmt.Println("(4) Verify:", result)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
```
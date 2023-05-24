# GO Openzeppelin Standard Tree

A Go library to generate merkle trees and merkle proofs.

## @openzeppelin/merkle-tree

fork from [@openzeppelin/merkle-tree](https://github.com/OpenZeppelin/merkle-tree)

contracts doc [MerkleProof](https://docs.openzeppelin.com/contracts/4.x/api/utils#MerkleProof)

## Installation
```
go get github.com/FantasyJony/openzeppelin-merkle-tree-go@v1.0.3
```

```
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

	// (3) dump -> tree
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

	// (4) dump leaf proof
	leafProofDump, err := tree.TreeProofMarshal()
	handleError(err)
	fmt.Println("(4) Dump Leaf Proof: ", string(leafProofDump))

	// (5) verify proof
	firstProof, err := newTree.GetProofWithIndex(0)
	handleError(err)
	firstValue := newTree.Leaves[0].Value

	result, err := newTree.Verify(firstProof, firstValue)
	handleError(err)
	fmt.Println("(5) Verify:", result)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
```

### Dump

1. TreeMarshal

```markdown
{
    "format": "standard-v1",
    "tree": [
        "0xd4dee0beab2d53f2cc83e567171bd2820e49898130a22622b10ead383e90bd77",
        "0xeb02c421cfa48976e66dfb29120745909ea3a0f843456c263cf8f1253483e283",
        "0xb92c48e9d7abe27fd8dfd6b5dfdbfb1c9a463f80c712b66f3a5180a090cccafc"
    ],
    "values": [
        {
            "value": [
                "0x1111111111111111111111111111111111111111",
                "5000000000000000000"
            ],
            "treeIndex": 1
        },
        {
            "value": [
                "0x2222222222222222222222222222222222222222",
                "2500000000000000000"
            ],
            "treeIndex": 2
        }
    ],
    "leafEncoding": [
        "address",
        "uint256"
    ]
}
```

2. TreeProofMarshal

```markdown
{
    "root": "0xd4dee0beab2d53f2cc83e567171bd2820e49898130a22622b10ead383e90bd77",
    "leafEncoding": [
        "address",
        "uint256"
    ],
    "leafProof": [
        {
            "value": [
                "0x1111111111111111111111111111111111111111",
                "5000000000000000000"
            ],
            "proof": [
                "0xb92c48e9d7abe27fd8dfd6b5dfdbfb1c9a463f80c712b66f3a5180a090cccafc"
            ]
        },
        {
            "value": [
                "0x2222222222222222222222222222222222222222",
                "2500000000000000000"
            ],
            "proof": [
                "0xeb02c421cfa48976e66dfb29120745909ea3a0f843456c263cf8f1253483e283"
            ]
        }
    ]
}
```

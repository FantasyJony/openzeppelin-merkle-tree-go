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

	fmt.Println("=== (1) Of BEGIN ===")

	// (1) make tree
	t1, err := smt.Of(values, leafEncodings)
	handleError(err)
	fmt.Println("Root: ", hexutil.Encode(t1.GetRoot()))

	fmt.Println("=== (1) Of END ===")

	fmt.Println("=== (2) Dump BEGIN ===")
	// (2) dump
	dump, err := t1.TreeMarshal()
	handleError(err)
	fmt.Println(string(dump))

	fmt.Println("=== (2) Dump END ===")

	fmt.Println("=== (3) Load BEGIN ===")
	// (3) load
	t2, err := smt.Load(dump)
	handleError(err)

	entries := t2.Entries()
	for k, v := range entries {
		if v.Value[0] == smt.SolAddress("0x1111111111111111111111111111111111111111") {
			proof, err := t2.GetProofWithIndex(k)
			handleError(err)
			fmt.Println("ValueIndex:", v.ValueIndex, " TreeIndex:", v.TreeIndex, " Hash:", hexutil.Encode(v.Hash), " Value:", v.Value)
			for pk, pv := range proof {
				fmt.Println("Proof k:", pk, " v:", hexutil.Encode(pv))
			}
		}
	}
	fmt.Println("=== (3) Load END ===")

	fmt.Println("=== (4) DumpLeafProof BEGIN ===")

	// (4) dump leaf proof
	leafProofDump, err := t2.TreeProofMarshal()
	handleError(err)
	fmt.Println(string(leafProofDump))

	fmt.Println("=== (4) DumpLeafProof END ===")

	fmt.Println("=== (5) Verify BEGIN ===")
	// (5) verify proof
	firstProof, err := t2.GetProofWithIndex(0)
	handleError(err)
	firstValue := entries[0].Value

	verified, err := t2.Verify(firstProof, firstValue)
	handleError(err)
	fmt.Println("Verify:", verified)
	fmt.Println("=== (5) Verify END ===")

	//fmt.Println("=== (6) CreateTree BEGIN ===")

	// (6) create tree
	t3, err := smt.CreateTree(leafEncodings)

	//fmt.Println("=== (6) CreateTree END ===")

	fmt.Println("=== (7) AddLeaf BEGIN ===")

	// (7) add leaf
	leafHash0, err := t3.AddLeaf(values[0])
	fmt.Println("index: 0 , AddLeaf: ", hexutil.Encode(leafHash0))

	leafHash1, err := t3.AddLeaf(values[1])
	fmt.Println("index: 1 , AddLeaf: ", hexutil.Encode(leafHash1))

	fmt.Println("=== (7) AddLeaf END ===")

	fmt.Println("=== (8) MakeTree BEGIN ===")

	// (8) make tree
	r3, err := t3.MakeTree()
	fmt.Println("(8) MakeTree Root: ", hexutil.Encode(r3))

	fmt.Println("=== (8) MakeTree END ===")

}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

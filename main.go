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
		{
			smt.SolAddress("0x3333333333333333333333333333333333333333"),
			smt.SolNumber("1500000000000000000"),
		},
		{
			smt.SolAddress("0x4444444444444444444444444444444444444444"),
			smt.SolNumber("1000000000000000000"),
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

	fmt.Println()
	fmt.Println("=== (2) CreateTree BEGIN ===")

	// (2.1) create tree
	t2, err := smt.CreateTree(leafEncodings)

	// (2.2) add leaf
	fmt.Println("=== (2.2) AddLeaf BEGIN ===")
	for k, v := range values {
		leafHash0, err := t2.AddLeaf(v)
		handleError(err)
		fmt.Println(k, " AddLeaf: ", hexutil.Encode(leafHash0))
	}
	fmt.Println("=== (2.2) AddLeaf END ===")

	// (2.3) make tree
	fmt.Println("=== (2.3) MakeTree BEGIN ===")
	r3, err := t2.MakeTree()
	fmt.Println("MakeTree Root: ", hexutil.Encode(r3))
	fmt.Println("=== (2.3) MakeTree END ===")

	fmt.Println()
	// (3) dump
	fmt.Println("=== (3) Dump BEGIN ===")
	dump, err := t1.TreeMarshal()
	handleError(err)
	fmt.Println(string(dump))
	fmt.Println("=== (3) Dump END ===")

	fmt.Println()
	// (4) load
	fmt.Println("=== (4) Load BEGIN ===")
	t3, err := smt.Load(dump)
	handleError(err)

	entries := t3.Entries()
	for k, v := range entries {
		if v.Value[0] == smt.SolAddress("0x1111111111111111111111111111111111111111") {
			proof, err := t3.GetProofWithIndex(k)
			handleError(err)
			fmt.Println(fmt.Sprintf("ValueIndex: %v , TreeIndex: %v , Hash: %v ,  Value: %v", v.ValueIndex, v.TreeIndex, hexutil.Encode(v.Hash), v.Value))
			for pk, pv := range proof {
				fmt.Println(fmt.Sprintf("[%v] = %v", pk, hexutil.Encode(pv)))
			}
		}
	}
	fmt.Println("=== (4) Load END ===")

	fmt.Println()
	// (5) dump leaf proof
	fmt.Println("=== (5) DumpLeafProof BEGIN ===")

	leafProofDump, err := t3.TreeProofMarshal()
	handleError(err)
	fmt.Println(string(leafProofDump))
	fmt.Println("=== (5) DumpLeafProof END ===")

	fmt.Println()
	// (6) verify proof
	fmt.Println("=== (6) Verify BEGIN ===")
	firstProof, err := t3.GetProofWithIndex(0)
	handleError(err)
	firstValue := entries[0].Value

	verified, err := t3.Verify(firstProof, firstValue)
	handleError(err)
	fmt.Println("Verify:", verified)
	fmt.Println("=== (6) Verify END ===")

	fmt.Println()
	// (7)
	fmt.Println("=== (7) VerifyMultiProof BEGIN ===")
	mulitProof, err := t3.GetMultiProof([][]interface{}{values[0], values[1]})
	//t3.VerifyMultiProof()
	fmt.Println("Proof:")
	for k, v := range mulitProof.Proof {
		fmt.Println(fmt.Sprintf("[%v] = %v", k, hexutil.Encode(v)))
	}
	//fmt.Println("=== (7.1) Proof END ===")
	//fmt.Println()
	fmt.Println("ProofFlags:")
	//fmt.Println("=== (7.2) ProofFlags BEGIN ===")
	fmt.Println(mulitProof.ProofFlags)
	//fmt.Println("=== (7.2) ProofFlags END ===")
	//fmt.Println()

	fmt.Println("Values:")
	for _, v := range mulitProof.Values {
		fmt.Println(v)
	}

	verified, err = t3.VerifyMultiProof(mulitProof)
	handleError(err)
	fmt.Println("Verify:", verified)

	fmt.Println("=== (7) VerifyMultiProof END ===")
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

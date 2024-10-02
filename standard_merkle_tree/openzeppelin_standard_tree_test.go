package standard_merkle_tree

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"testing"
)

func TestSMTCreateTree(t *testing.T) {

	tree, err := CreateTree([]string{SOL_ADDRESS, SOL_UINT256})
	if err != nil {
		fmt.Println("CreateTree ERR: ", err)
	}

	leaf1 := []interface{}{
		SolAddress("0x1111111111111111111111111111111111111111"),
		SolNumber("5000000000000000000"),
	}

	hash1, err := tree.AddLeaf(leaf1)
	if err != nil {
		fmt.Println("AddLeaf ERR: ", err)
	}
	fmt.Println("01 AddLeaf Hash: ", hexutil.Encode(hash1))

	leaf2 := []interface{}{
		SolAddress("0x2222222222222222222222222222222222222222"),
		SolNumber("2500000000000000000"),
	}

	hash2, err := tree.AddLeaf(leaf2)
	if err != nil {
		fmt.Println("AddLeaf ERR: ", err)
	}
	fmt.Println("02 AddLeaf Hash: ", hexutil.Encode(hash2))

	root, err := tree.MakeTree()
	if err != nil {
		fmt.Println("MakeTree ERR: ", err)
	}
	fmt.Println("03 Merkle Root: ", hexutil.Encode(root))
}

func TestSMTOf(t *testing.T) {

	leaf1 := []interface{}{
		SolAddress("0x1111111111111111111111111111111111111111"),
		SolNumber("5000000000000000000"),
	}

	leaf2 := []interface{}{
		SolAddress("0x2222222222222222222222222222222222222222"),
		SolNumber("2500000000000000000"),
	}

	leaves := [][]interface{}{
		leaf1,
		leaf2,
	}

	tree, err := Of(
		leaves,
		[]string{
			SOL_ADDRESS,
			SOL_UINT256,
		})

	if err != nil {
		fmt.Println("Of ERR", err)
	}

	root := hexutil.Encode(tree.GetRoot())
	fmt.Println("Merkle Root: ", root)

	proof, err := tree.GetProof(leaf1)
	strProof := make([]string, len(proof))
	if err != nil {
		fmt.Println("GetProof ERR", err)
	}
	for _, v := range proof {
		strProof = append(strProof, hexutil.Encode(v))
	}
	fmt.Println("02 proof: ", strProof)
}

func TestSMTVerify(t *testing.T) {

	root, err := hexutil.Decode("0xd4dee0beab2d53f2cc83e567171bd2820e49898130a22622b10ead383e90bd77")
	if err != nil {
		fmt.Println(err)
	}

	leafEncodings := []string{
		SOL_ADDRESS,
		SOL_UINT256,
	}

	value := []interface{}{
		SolAddress("0x1111111111111111111111111111111111111111"),
		SolNumber("5000000000000000000"),
	}

	proofValue01, err := hexutil.Decode("0xb92c48e9d7abe27fd8dfd6b5dfdbfb1c9a463f80c712b66f3a5180a090cccafc")
	if err != nil {
		fmt.Println(err)
	}
	proof := [][]byte{
		proofValue01,
	}

	leaf, err := LeafHash(leafEncodings, value)
	if err != nil {
		fmt.Println(err)
	}

	verified, err := Verify(root, leaf, proof)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(verified)
}

func TestSMTVerifyMultiProof(t *testing.T) {

	root, err := hexutil.Decode("0xd4dee0beab2d53f2cc83e567171bd2820e49898130a22622b10ead383e90bd77")
	if err != nil {
		fmt.Println(err)
	}

	leafEncodings := []string{
		SOL_ADDRESS,
		SOL_UINT256,
	}

	value := []interface{}{
		SolAddress("0x1111111111111111111111111111111111111111"),
		SolNumber("5000000000000000000"),
	}

	leaf1, err := LeafHash(leafEncodings, value)
	if err != nil {
		fmt.Println(err)
	}

	proofValue01, err := hexutil.Decode("0xb92c48e9d7abe27fd8dfd6b5dfdbfb1c9a463f80c712b66f3a5180a090cccafc")
	if err != nil {
		fmt.Println(err)
	}
	proof := [][]byte{
		proofValue01,
	}

	proofFlags := []bool{
		false,
	}

	leaves := [][]byte{
		leaf1,
	}

	multiProof := &MultiProof{
		Proof:      proof,
		ProofFlags: proofFlags,
		Leaves:     leaves,
	}

	verified, err := VerifyMultiProof(root, multiProof)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(verified)
}

func TestGetProofAndVerify(t *testing.T) {

	leaf1 := []interface{}{
		SolAddress("0x1111111111111111111111111111111111111111"),
		SolNumber("5000000000000000000"),
	}

	leaf2 := []interface{}{
		SolAddress("0x2222222222222222222222222222222222222222"),
		SolNumber("2500000000000000000"),
	}

	leaves := [][]interface{}{
		leaf1,
		leaf2,
	}

	tree, err := Of(
		leaves,
		[]string{
			SOL_ADDRESS,
			SOL_UINT256,
		})

	if err != nil {
		fmt.Println("Of ERR", err)
	}

	root := hexutil.Encode(tree.GetRoot())
	fmt.Println("01 Merkle Root: ", root)

	proof, err := tree.GetProof(leaf1)
	strProof := make([]string, len(proof))
	if err != nil {
		fmt.Println("GetProof ERR", err)
	}

	for _, v := range proof {
		strProof = append(strProof, hexutil.Encode(v))
	}
	fmt.Println("02 proof: ", strProof)

	proof2, err := tree.GetProofWithIndex(0)
	strProof2 := make([]string, len(proof2))
	if err != nil {
		fmt.Println("GetProof ERR", err)
	}

	for _, v := range proof2 {
		strProof2 = append(strProof2, hexutil.Encode(v))
	}
	fmt.Println("03 proof index: ", strProof2)

	isVerify, err := tree.Verify(proof, leaf1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("04 verify: ", isVerify)

	isIndexVerify, err := tree.VerifyWithIndex(proof, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("05 verifyWithIndex: ", isIndexVerify)

}

func TestGetMultiProofAndVerify(t *testing.T) {

	leaf1 := []interface{}{
		SolAddress("0x1111111111111111111111111111111111111111"),
		SolNumber("5000000000000000000"),
	}

	leaf2 := []interface{}{
		SolAddress("0x2222222222222222222222222222222222222222"),
		SolNumber("2500000000000000000"),
	}

	leaves := [][]interface{}{
		leaf1,
		leaf2,
	}

	tree, err := Of(
		leaves,
		[]string{
			SOL_ADDRESS,
			SOL_UINT256,
		})

	if err != nil {
		fmt.Println("Of ERR", err)
	}

	root := hexutil.Encode(tree.GetRoot())
	fmt.Println("01 Merkle Root: ", root)

	proof, err := tree.GetMultiProof([][]interface{}{leaf1})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("02 leaf1 proof: ", proof)

	proof02, err := tree.GetMultiProofWithIndices([]int{0})
	fmt.Println("03 leaf1 proof: ", proof02)

	multiValue, err := tree.VerifyMultiProof(proof)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("04 VerifyMultiProof: ", multiValue)
}

func TestDumpOf(t *testing.T) {

	leaf1 := []interface{}{
		SolAddress("0x1111111111111111111111111111111111111111"),
		SolNumber("5000000000000000000"),
	}

	leaf2 := []interface{}{
		SolAddress("0x2222222222222222222222222222222222222222"),
		SolNumber("2500000000000000000"),
	}

	leaves := [][]interface{}{
		leaf1,
		leaf2,
	}

	tree, err := Of(
		leaves,
		[]string{
			SOL_ADDRESS,
			SOL_UINT256,
		})

	if err != nil {
		fmt.Println("Of ERR", err)
	}

	root := hexutil.Encode(tree.GetRoot())
	fmt.Println("01 Merkle Root: ", root)

	fmt.Println("02 TreeMarshal")
	value, err := tree.TreeMarshal()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(value))

	fmt.Println("03 TreeUnmarshal")
	tree2, err := TreeUnmarshal(value)
	value2, err := tree2.TreeMarshal()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(value2))

	fmt.Println("04 Load")
	tree3, err := Load([]byte(string(value2)))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hexutil.Encode(tree3.GetRoot()))
}

func TestDumpLeafProof(t *testing.T) {

	leaf1 := []interface{}{
		SolAddress("0x1111111111111111111111111111111111111111"),
		//SolNumber("5000000000000000000"),
		SolNumber("500"),
	}

	leaf2 := []interface{}{
		SolAddress("0x2222222222222222222222222222222222222222"),
		//SolNumber("2500000000000000000"),
		SolNumber("250"),
	}

	leaves := [][]interface{}{
		leaf1,
		leaf2,
	}

	tree, err := Of(
		leaves,
		[]string{
			SOL_ADDRESS,
			SOL_UINT8,
		})

	if err != nil {
		fmt.Println("Of ERR", err)
	}

	root := hexutil.Encode(tree.GetRoot())
	fmt.Println("01 Merkle Root: ", root)

	fmt.Println("02 DumpLeafProof")
	proof, err := tree.DumpLeafProof()
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range proof.Proofs {
		bProof := make([][]byte, len(v.Proof))
		for a, b := range v.Proof {
			bProof[a], _ = hexutil.Decode(b)
		}
		r, err := tree.Verify(bProof, v.getSolValueUnmarshal(proof.LeafEncoding))
		if err != nil {
			fmt.Println("Verify ERR: ", err)
			return
		}
		fmt.Println("03 Verify Proof ", k, " :", r)
	}

	fmt.Println("04 TreeProofMarshal")
	proofJson, err := tree.TreeProofMarshal()
	if err != nil {
		fmt.Println("TreeProofMarshal ERR: ", err)
	}
	fmt.Println(string(proofJson))
}

func TestArrayArg(t *testing.T) {

	values := [][]interface{}{
		{
			SolAddress("0x1111111111111111111111111111111111111111"),
			SolNumberArray([]interface{}{
				"5000000000000000000", "2500000000000000000",
			}),
			SolBoolArray([]interface{}{
				true,
				false,
			}),
		},
		{
			SolAddress("0x2222222222222222222222222222222222222222"),
			SolNumberArray([]interface{}{
				"2500000000000000000", "5000000000000000000",
			}),
			SolBoolArray([]interface{}{
				false,
				true,
			}),
		},
	}

	tree, err := Of(
		values,
		[]string{
			SOL_ADDRESS,
			SOL_UINT256_ARRAY,
			SOL_BOOL_ARRAY,
		})

	if err != nil {
		fmt.Println("Of ERR: ", err)
	}

	root := hexutil.Encode(tree.GetRoot())
	fmt.Println("01 Merkle Root: ", root)

	proof, err := tree.GetProof(values[0])
	strProof := make([]string, len(proof))
	if err != nil {
		fmt.Println("GetProof ERR", err)
	}
	for _, v := range proof {
		strProof = append(strProof, hexutil.Encode(v))
	}
	fmt.Println("02 proof: ", strProof)

	fmt.Println("03 TreeMarshal")
	value, err := tree.TreeMarshal()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(value))

	fmt.Println("04 TreeUnmarshal")
	tree2, err := TreeUnmarshal(value)
	value2, err := tree2.TreeMarshal()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(value2))
}

func TestAbiArg(t *testing.T) {
	values := [][]interface{}{
		{
			SolAddress("0x1111111111111111111111111111111111111111"),
			SolNumber("1"),
			SolNumber("2"),
			SolNumber("3"),
			//SolNumberArray([]interface{}{
			//	"1", "2", "3",
			//}),
		},
		{
			SolAddress("0x1111111111111111111111111111111111111111"),
			SolNumber("0"),
			SolNumber("2"),
			SolNumber("1"),
			//SolNumberArray([]interface{}{
			//	"0", "1", "2",
			//}),
		},
	}

	leafEncodings := []string{
		SOL_ADDRESS,
		SOL_UINT8,
		SOL_UINT88,
		SOL_UINT128,
		//SOL_UINT8_ARRAY,
	}

	tree, err := Of(
		values,
		leafEncodings)

	if err != nil {
		fmt.Println("Of ERR: ", err)
	}

	root := hexutil.Encode(tree.GetRoot())
	fmt.Println("01 Merkle Root: ", root)

	proof, err := tree.GetProof(values[0])
	strProof := make([]string, len(proof))
	if err != nil {
		fmt.Println("GetProof ERR", err)
	}
	for _, v := range proof {
		strProof = append(strProof, hexutil.Encode(v))
	}
	fmt.Println("02 proof: ", strProof)

	fmt.Println("03 TreeMarshal")
	value, err := tree.TreeMarshal()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(value))

	fmt.Println("04 TreeUnmarshal")
	tree2, err := TreeUnmarshal(value)
	value2, err := tree2.TreeMarshal()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(value2))
}

func TestStringArg(t *testing.T) {

	values := [][]interface{}{
		{
			SolString("a"),
		},
		{
			SolString("b"),
		},
		{
			SolString("c"),
		},
		{
			SolString("d"),
		},
	}

	leafEncodings := []string{
		SOL_STRING,
	}

	t1, err := Of(values, leafEncodings)
	if err != nil {
		println("error:", err.Error())
		return
	}
	fmt.Println("Root: ", hexutil.Encode(t1.GetRoot()))
}

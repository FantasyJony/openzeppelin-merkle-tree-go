package standard_merkle_tree

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"testing"
)

func TestCreateTree(t *testing.T) {

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

func TestOf(t *testing.T) {

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

	tree, err := Of([]string{
		SOL_ADDRESS,
		SOL_UINT256,
	}, leaves)

	if err != nil {
		fmt.Println("Of ERR", err)
	}

	root := hexutil.Encode(tree.GetRoot())
	fmt.Println("Merkle Root: ", root)
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

	tree, err := Of([]string{
		SOL_ADDRESS,
		SOL_UINT256,
	}, leaves)

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

	tree, err := Of([]string{
		SOL_ADDRESS,
		SOL_UINT256,
	}, leaves)

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

	multiValue, err := tree.VerifyMultiProof(proof)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("03 VerifyMultiProof: ", multiValue)
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

	tree, err := Of([]string{
		SOL_ADDRESS,
		SOL_UINT256,
	}, leaves)

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
}

func TestDumpLeafProof(t *testing.T) {

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

	tree, err := Of([]string{
		SOL_ADDRESS,
		SOL_UINT256,
	}, leaves)

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

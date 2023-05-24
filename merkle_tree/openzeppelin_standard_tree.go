package merkle_tree

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	json_iterator "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"sort"
)

var (
	json = json_iterator.ConfigCompatibleWithStandardLibrary
)

type StandardTree struct {
	leafEncodings []string
	leaves        []*LeafValue
	hashLookup    map[string]int
	tree          [][]byte
}

type LeafValue struct {
	ValueIndex int
	TreeIndex  int
	Hash       []byte
	Value      []interface{}
}

func (l *LeafValue) getSolValueMarshal(leafEncoding []string) []interface{} {
	values := make([]interface{}, len(leafEncoding))
	for k, v := range leafEncoding {
		values[k] = SolValueMarshal(l.Value[k], v)
	}
	return values
}

type ValueMultiProof struct {
	Proof      [][]byte
	ProofFlags []bool
	Values     [][]interface{}
}

type StandardValueData struct {
	Value     []interface{} `json:"value"`
	TreeIndex int           `json:"treeIndex"`
}

func (svd *StandardValueData) getSolValueUnmarshal(leafEncoding []string) []interface{} {
	values := make([]interface{}, len(svd.Value))
	for k, v := range svd.Value {
		values[k] = SolValueUnmarshal(v, leafEncoding[k])
	}
	return values
}

type StandardMerkleTreeData struct {
	Format       string               `json:"format"`
	Tree         []string             `json:"tree"`
	Values       []*StandardValueData `json:"values"`
	LeafEncoding []string             `json:"leafEncoding"`
}

type StandardMerkleLeafProofData struct {
	Value []interface{} `json:"value"`
	Proof []string      `json:"proof"`
}

func (lpd *StandardMerkleLeafProofData) getSolValueUnmarshal(leafEncoding []string) []interface{} {
	values := make([]interface{}, len(leafEncoding))
	for k, v := range leafEncoding {
		values[k] = SolValueUnmarshal(lpd.Value[k], v)
	}
	return values
}

type StandardMerkleTreeProofData struct {
	Root         string                         `json:"root"`
	LeafEncoding []string                       `json:"leafEncoding"`
	Proofs       []*StandardMerkleLeafProofData `json:"leafProof"`
}

func (l *LeafValue) ToString() string {
	return fmt.Sprintf("ValueIndex: %v, TreeIndex: %v , Hash:%v", l.ValueIndex, l.TreeIndex, hexutil.Encode(l.Hash))
}

func CreateTree(leafEncodings []string) (*StandardTree, error) {
	return &StandardTree{
		leafEncodings: leafEncodings,
	}, nil
}

func Of(leafEncodings []string, leaves [][]interface{}) (*StandardTree, error) {
	tree, err := CreateTree(leafEncodings)
	if err != nil {
		return nil, err
	}
	for _, v := range leaves {
		_, err := tree.AddLeaf(v)
		if err != nil {
			return nil, err
		}
	}
	_, err = tree.MakeTree()
	if err != nil {
		return nil, err
	}
	return tree, nil
}

func (st *StandardTree) checkLeafEncodings() error {
	if st.leafEncodings == nil || len(st.leafEncodings) == 0 {
		return errors.New("LeafEncodings is not in StandardTree")
	}
	return nil
}

func (st *StandardTree) validateValue(valueIndex int) ([]byte, error) {
	err := checkBounds(valueIndex, st.leaves)
	if err != nil {
		return nil, err
	}
	leafValue, err := st.getLeafValue(valueIndex)
	if err != nil {
		return nil, err
	}
	treeIndex := leafValue.TreeIndex
	err = checkBounds(treeIndex, st.tree)
	if err != nil {
		return nil, err
	}
	value := leafValue.Value
	leaf, err := st.LeafHash(value)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(leaf, st.tree[treeIndex]) == false {
		return nil, errors.New("Merkle tree does not contain the expected value")
	}
	return leaf, nil
}

func (st *StandardTree) getLeafValue(valueIndex int) (*LeafValue, error) {
	err := checkBounds(valueIndex, st.leaves)
	if err != nil {
		return nil, err
	}
	for _, v := range st.leaves {
		if v.ValueIndex == valueIndex {
			return v, nil
		}
	}
	return nil, errors.New("valueIndex is not in leaves")
}

func (st *StandardTree) getValueIndexWithLeafHash(leafHash []byte) (int, error) {
	value, ok := st.hashLookup[hexutil.Encode(leafHash)]
	if ok == false {
		return 0, errors.New("Leaf is not in tree")
	}
	return value, nil
}

func (st *StandardTree) leafLookup(values []interface{}) (int, error) {
	leafHash, err := st.LeafHash(values)
	if err != nil {
		return 0, err
	}
	return st.getValueIndexWithLeafHash(leafHash)
}

func (st *StandardTree) LeafHash(values []interface{}) ([]byte, error) {
	return abiPackLeafHash(st.leafEncodings, values...)
}

func (st *StandardTree) LeafHashWithIndex(index int) ([]byte, error) {
	return st.validateValue(index)
}

func (st *StandardTree) GetProof(leaf []interface{}) ([][]byte, error) {
	valueIndex, err := st.leafLookup(leaf)
	if err != nil {
		return nil, err
	}
	return st.GetProofWithIndex(valueIndex)
}

func (st *StandardTree) GetProofWithIndex(valueIndex int) ([][]byte, error) {
	_, err := st.validateValue(valueIndex)
	if err != nil {
		return nil, err
	}

	leafValue, err := st.getLeafValue(valueIndex)
	if err != nil {
		return nil, err
	}

	treeIndex := leafValue.TreeIndex
	proof, err := getProof(st.tree, treeIndex)

	if err != nil {
		return nil, err
	}
	verifyResult, err := verify(st.GetRoot(), st.tree[treeIndex], proof)
	if err != nil {
		return nil, err
	}
	if !verifyResult {
		return nil, errors.New("unable to prove value")
	}
	return proof, nil
}

func (st *StandardTree) Verify(proof [][]byte, leaf []interface{}) (bool, error) {
	leafHash, err := st.LeafHash(leaf)
	if err != nil {
		return false, err
	}
	return verify(st.GetRoot(), leafHash, proof)
}

func (st *StandardTree) VerifyWithIndex(proof [][]byte, index int) (bool, error) {
	leafHash, err := st.LeafHashWithIndex(index)
	if err != nil {
		return false, err
	}
	return verify(st.GetRoot(), leafHash, proof)
}

func (st *StandardTree) GetMultiProof(leaves [][]interface{}) (*ValueMultiProof, error) {
	valueIndices := make([]int, len(leaves))
	for k, v := range leaves {
		index, err := st.leafLookup(v)
		if err != nil {
			return nil, err
		}
		valueIndices[k] = index
	}
	return st.GetMultiProofWithIndices(valueIndices)
}

func (st *StandardTree) GetMultiProofWithIndices(indices []int) (*ValueMultiProof, error) {
	for _, v := range indices {
		_, err := st.validateValue(v)
		if err != nil {
			return nil, err
		}
	}
	treeIndices := make([]int, len(indices))
	for k, v := range indices {
		leaf, err := st.getLeafValue(v)
		if err != nil {
			return nil, err
		}
		treeIndices[k] = leaf.TreeIndex
	}
	proof, err := getMultiProof(st.tree, treeIndices)
	if err != nil {
		return nil, err
	}
	verifyResult, err := verifyMultiProof(st.GetRoot(), proof)
	if verifyResult == false {
		return nil, errors.New("Unable to prove values")
	}
	leaves := make([][]interface{}, len(proof.Leaves))
	for k, v := range proof.Leaves {
		valueIndex, err := st.getValueIndexWithLeafHash(v)
		if err != nil {
			return nil, err
		}
		leaf, err := st.getLeafValue(valueIndex)
		if err != nil {
			return nil, err
		}
		leaves[k] = leaf.Value
	}

	return &ValueMultiProof{
		Values:     leaves,
		Proof:      proof.Proof,
		ProofFlags: proof.ProofFlags,
	}, nil
}

func (st *StandardTree) VerifyMultiProof(multiProof *ValueMultiProof) (bool, error) {
	leaves := make([][]byte, len(multiProof.Values))
	for k, v := range multiProof.Values {
		leafHash, err := st.LeafHash(v)
		if err != nil {
			return false, err
		}
		leaves[k] = leafHash
	}
	proof := &MultiProof{
		Proof:      multiProof.Proof,
		ProofFlags: multiProof.ProofFlags,
		Leaves:     leaves,
	}
	return verifyMultiProof(st.GetRoot(), proof)
}

func (st *StandardTree) AddLeaf(values []interface{}) ([]byte, error) {

	err := st.checkLeafEncodings()
	if err != nil {
		return nil, err
	}

	if len(st.leafEncodings) != len(values) {
		return nil, errors.New("Length mismatch")
	}

	leafHash, err := st.LeafHash(values)
	if err != nil {
		return nil, err
	}

	leafValue := &LeafValue{
		Value:      values,
		ValueIndex: len(st.leaves),
		Hash:       leafHash,
	}

	st.leaves = append(st.leaves, leafValue)
	return leafHash, nil
}

func (st *StandardTree) MakeTree() ([]byte, error) {

	if len(st.leaves) == 0 {
		return nil, errors.New("Expected non-zero number of leaves")
	}

	// sort hash
	leafValues := make([]*LeafValue, len(st.leaves))
	for k, v := range st.leaves {
		leafValues[k] = v
	}
	sort.Slice(leafValues, func(i, j int) bool {
		return compareBytes(leafValues[i].Hash, leafValues[j].Hash)
	})
	leaves := make([][]byte, len(leafValues))
	for k, v := range leafValues {
		leaves[k] = v.Hash
	}

	// make tree
	tree, err := makeMerkleTree(leaves)
	if err != nil {
		return nil, err
	}

	hashLookup := make(map[string]int)
	for leafIndex, v := range leafValues {
		treeIndex := len(tree) - leafIndex - 1
		v.TreeIndex = treeIndex
		hashLookup[hexutil.Encode(v.Hash)] = v.ValueIndex
	}

	st.hashLookup = hashLookup
	st.tree = tree
	return st.GetRoot(), nil
}

func (st *StandardTree) GetRoot() []byte {
	return st.tree[0]
}

func (st *StandardTree) Dump() *StandardMerkleTreeData {
	tree := make([]string, len(st.tree))
	for k, v := range st.tree {
		tree[k] = hexutil.Encode(v)
	}
	valueData := make([]*StandardValueData, len(st.leaves))
	for k, v := range st.leaves {
		valueData[k] = &StandardValueData{
			TreeIndex: v.TreeIndex,
			Value:     v.getSolValueMarshal(st.leafEncodings),
		}
	}
	return &StandardMerkleTreeData{
		Format:       "standard-v1",
		LeafEncoding: st.leafEncodings,
		Tree:         tree,
		Values:       valueData,
	}
}

func (st *StandardTree) TreeMarshal() ([]byte, error) {
	treeData := st.Dump()
	return json.Marshal(treeData)
}

func TreeUnmarshal(value []byte) (*StandardTree, error) {
	var std StandardMerkleTreeData
	err := json.Unmarshal(value, &std)
	if err != nil {
		return nil, err
	}
	values := make([][]interface{}, len(std.Values))
	for k, v := range std.Values {
		values[k] = v.getSolValueUnmarshal(std.LeafEncoding)
	}
	tree, err := Of(std.LeafEncoding, values)
	return tree, err
}

func (st *StandardTree) DumpLeafProof() (*StandardMerkleTreeProofData, error) {
	leafProofData := make([]*StandardMerkleLeafProofData, len(st.leaves))
	for k, v := range st.leaves {
		values := v.getSolValueMarshal(st.leafEncodings)
		leafProof, err := st.GetProof(v.Value)
		if err != nil {
			return nil, err
		}
		proof := make([]string, len(leafProof))
		for a, b := range leafProof {
			proof[a] = hexutil.Encode(b)
		}
		leafProofData[k] = &StandardMerkleLeafProofData{
			Value: values,
			Proof: proof,
		}
	}
	return &StandardMerkleTreeProofData{
		Root:         hexutil.Encode(st.GetRoot()),
		Proofs:       leafProofData,
		LeafEncoding: st.leafEncodings,
	}, nil
}

func (st *StandardTree) TreeProofMarshal() ([]byte, error) {
	treeData, err := st.DumpLeafProof()
	if err != nil {
		return nil, err
	}
	return json.Marshal(treeData)
}

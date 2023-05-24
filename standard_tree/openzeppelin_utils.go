package standard_tree

import (
	"bytes"
	"github.com/pkg/errors"
	"math"
	"sort"
)

func abiPackLeafHash(leafEncodings []string, values ...interface{}) ([]byte, error) {
	data, err := AbiPack(leafEncodings, values...)
	if err != nil {
		return nil, err
	}
	hash, err := standardLeafHash(data)
	return hash, err
}

func standardLeafHash(value []byte) ([]byte, error) {
	k1, err := Keccak256(value)
	if err != nil {
		return nil, err
	}
	k2, err := Keccak256(k1)
	return k2, err
}

func hashPair(a, b []byte) ([]byte, error) {
	pairs := [][]byte{a, b}
	sort.Slice(pairs, func(i, j int) bool {
		return compareBytes(pairs[i], pairs[j])
	})
	var data []byte
	for _, v := range pairs {
		data = append(data, v...)
	}
	return Keccak256(data)
}

func makeMerkleTree(leaves [][]byte) ([][]byte, error) {
	if len(leaves) == 0 {
		return nil, errors.New("Expected non-zero number of leaves")
	}
	sort.Slice(leaves, func(i, j int) bool {
		return compareBytes(leaves[i], leaves[j])
	})
	for _, v := range leaves {
		err := checkValidMerkleNode(v)
		if err != nil {
			return nil, err
		}
	}
	lenTree := 2*len(leaves) - 1
	hashTree := make([][]byte, lenTree)
	for i, v := range leaves {
		hashTree[lenTree-1-i] = v
	}
	for i := lenTree - 1 - len(leaves); i >= 0; i-- {
		hash, err := hashPair(hashTree[leftChildIndex(i)], hashTree[rightChildIndex(i)])
		if err != nil {
			return nil, err
		}
		hashTree[i] = hash
	}
	return hashTree, nil
}

func compareBytes(a, b []byte) bool {
	n := len(a)
	if len(b) < len(a) {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func siblingIndex(i int) (int, error) {
	if i > 0 {
		return i - int(math.Pow(-1, float64(i%2))), nil
	} else {
		return 0, errors.New("Root has no siblings")
	}
}

func parentIndex(i int) (int, error) {
	if i > 0 {
		return int(math.Floor(float64(i-1) / 2)), nil
	} else {
		return 0, errors.New("Root has no siblings")
	}
}

func leftChildIndex(i int) int {
	return 2*i + 1
}

func rightChildIndex(i int) int {
	return 2*i + 2
}

func getProof(tree [][]byte, index int) ([][]byte, error) {
	err := checkLeafNode(tree, index)
	if err != nil {
		return nil, err
	}
	proof := make([][]byte, 0)
	for index > 0 {
		treeIndex, err := siblingIndex(index)
		if err != nil {
			return nil, err
		}
		proof = append(proof, tree[treeIndex])
		index, err = parentIndex(index)
		if err != nil {
			return nil, err
		}
	}
	return proof, nil
}

func verify(rootHash, leafHash []byte, proof [][]byte) (bool, error) {
	impliedRoot, err := processProof(leafHash, proof)
	if err != nil {
		return false, err
	}
	return bytes.Equal(rootHash, impliedRoot), nil
}

type MultiProof struct {
	Proof      [][]byte
	ProofFlags []bool
	Leaves     [][]byte
}

func getMultiProof(tree [][]byte, indices []int) (*MultiProof, error) {

	for _, v := range indices {
		err := checkLeafNode(tree, v)
		if err != nil {
			return nil, err
		}
	}

	if someValue(indices) {
		return nil, errors.New("Cannot prove duplicated index")
	}

	stack := make([]int, len(indices))
	copy(stack, indices)

	proofFlags := make([]bool, 0)
	proof := make([][]byte, 0)

	for len(stack) > 0 && stack[0] > 0 {

		j, err := shift(&stack)
		if err != nil {
			return nil, err
		}

		s, err := siblingIndex(j)
		if err != nil {
			return nil, err
		}

		p, err := parentIndex(j)
		if err != nil {
			return nil, err
		}

		if len(stack) > 0 && s == stack[0] {
			proofFlags = append(proofFlags, true)
			_, err := shift(&stack)
			if err != nil {
				return nil, err
			}
		} else {
			proofFlags = append(proofFlags, false)
			proof = append(proof, tree[s])
		}

		stack = append(stack, p)
	}

	if len(indices) == 0 {
		proof = append(proof, tree[0])
	}

	leaves := make([][]byte, len(indices))
	for k, v := range indices {
		leaves[k] = tree[v]
	}

	return &MultiProof{
		Proof:      proof,
		ProofFlags: proofFlags,
		Leaves:     leaves,
	}, nil
}

func verifyMultiProof(rootHash []byte, multiProof *MultiProof) (bool, error) {
	impliedRoot, err := processMultiProof(multiProof)
	if err != nil {
		return false, err
	}
	return bytes.Equal(rootHash, impliedRoot), nil
}

func processMultiProof(multiProof *MultiProof) ([]byte, error) {
	for _, v := range multiProof.Leaves {
		err := checkValidMerkleNode(v)
		if err != nil {
			return nil, err
		}
	}

	for _, v := range multiProof.Proof {
		err := checkValidMerkleNode(v)
		if err != nil {
			return nil, err
		}
	}

	unFlagCount := 0
	for _, v := range multiProof.ProofFlags {
		if v == false {
			unFlagCount++
		}
	}

	if len(multiProof.Proof) < unFlagCount {
		return nil, errors.New("Invalid multiProof format")
	}

	if len(multiProof.Leaves)+len(multiProof.Proof) != len(multiProof.ProofFlags)+1 {
		return nil, errors.New("Provided leaves and multiProof are not compatible")
	}

	stack := make([][]byte, len(multiProof.Leaves))
	copy(stack, multiProof.Leaves)

	proof := make([][]byte, len(multiProof.Proof))
	copy(proof, multiProof.Proof)

	for _, flag := range multiProof.ProofFlags {

		a, err := shift(&stack)
		if err != nil {
			return nil, err
		}

		var b []byte
		if flag {
			b, err = shift(&stack)
			if err != nil {
				return nil, err
			}
		} else {
			b, err = shift(&proof)
			if err != nil {
				return nil, err
			}
		}

		c, err := hashPair(a, b)
		if err != nil {
			return nil, err
		}

		stack = append(stack, c)
	}

	leafProof, err := pop(&stack)
	if err == nil {
		return leafProof, nil
	}

	leafProof, err = shift(&proof)
	if err != nil {
		return nil, err
	}

	return leafProof, nil
}

func processProof(leaf []byte, proof [][]byte) ([]byte, error) {
	err := checkValidMerkleNode(leaf)
	if err != nil {
		return nil, err
	}
	for _, v := range proof {
		err = checkValidMerkleNode(v)
		if err != nil {
			return nil, err
		}
	}

	leafProof := make([]byte, len(leaf))
	copy(leafProof, leaf)

	for _, v := range proof {
		hash, err := hashPair(leafProof, v)
		if err != nil {
			return nil, err
		}
		leafProof = hash
	}

	return leafProof, err
}

func isTreeNode(tree [][]byte, i int) bool {
	return i >= 0 && i < len(tree)
}

func isInternalNode(tree [][]byte, i int) bool {
	return isTreeNode(tree, leftChildIndex(i))
}

func isValidMerkleNode(hash []byte) bool {
	return len(hash) == 32
}

func isLeafNode(tree [][]byte, i int) bool {
	return isTreeNode(tree, i) && !isInternalNode(tree, i)
}

func checkValidMerkleNode(hash []byte) error {
	if isValidMerkleNode(hash) {
		return nil
	}
	return errors.New("Merkle tree nodes must length 32")
}

func checkLeafNode(tree [][]byte, i int) error {
	if isLeafNode(tree, i) {
		return nil
	}
	return errors.New("Index is not a leaf")
}

func shift[T any](values *[]T) (T, error) {
	if len(*values) == 0 {
		var empty T
		return empty, errors.New("Values len is zero")
	}
	firstElement := (*values)[0]
	*values = (*values)[1:]
	return firstElement, nil
}

func pop[T any](values *[]T) (T, error) {
	if len(*values) == 0 {
		var empty T
		return empty, errors.New("Values len is zero")
	}
	lastIdx := len(*values) - 1
	lastElement := (*values)[lastIdx]
	*values = (*values)[:lastIdx]
	return lastElement, nil
}

func checkBounds[T any](index int, values []T) error {
	if index < 0 || index >= len(values) {
		return errors.New("Index out of bounds")
	}
	return nil
}

func someValue[T comparable](indices []T) bool {
	for i := 1; i < len(indices); i++ {
		for p := 0; p < i; p++ {
			if indices[i] == indices[p] {
				return true
			}
		}
	}
	return false
}

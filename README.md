# Openzeppelin Standard Tree Go

A Go library to generate merkle trees and merkle proofs.

[![License](https://img.shields.io/github/license/FantasyJony/openzeppelin-merkle-tree-go)](https://opensource.org/licenses/Apache-2.0)
[![Tag](https://img.shields.io/github/v/tag/FantasyJony/openzeppelin-merkle-tree-go?sort=semver)](https://github.com/FantasyJony/openzeppelin-merkle-tree-go/tags)
[![Go version](https://img.shields.io/github/go-mod/go-version/FantasyJony/openzeppelin-merkle-tree-go)](https://golang.org/dl/#stable)
[![Go Reference](https://pkg.go.dev/badge/github.com/FantasyJony/openzeppelin-merkle-tree-go.svg)](https://pkg.go.dev/github.com/FantasyJony/openzeppelin-merkle-tree-go)

[![Hardhat](https://img.shields.io/node/v/hardhat)](https://hardhat.org/docs)
[![@openzeppelin/contracts](https://img.shields.io/github/package-json/dependency-version/FantasyJony/openzeppelin-merkle-tree-go/@openzeppelin/merkle-tree?filename=merkle-tree-contracts%2Fpackage.json)](https://github.com/OpenZeppelin/openzeppelin-contracts)
[![@openzeppelin/merkle-tree](https://img.shields.io/github/package-json/dependency-version/FantasyJony/openzeppelin-merkle-tree-go/@openzeppelin/contracts?filename=merkle-tree-contracts%2Fpackage.json)](https://github.com/OpenZeppelin/merkle-tree)

## @openzeppelin/merkle-tree

fork from [@openzeppelin/merkle-tree](https://github.com/OpenZeppelin/merkle-tree)

contracts doc [MerkleProof](https://docs.openzeppelin.com/contracts/4.x/api/utils#MerkleProof)

## Installation
```
go get github.com/FantasyJony/openzeppelin-merkle-tree-go@v1.1.0
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
	
    fmt.Println("Proof:")
    for k, v := range mulitProof.Proof {
        fmt.Println(fmt.Sprintf("[%v] = %v", k, hexutil.Encode(v)))
    }
	
    fmt.Println("ProofFlags:")
    fmt.Println(mulitProof.ProofFlags)
    
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
```

Run: [main_run.txt](https://github.com/FantasyJony/openzeppelin-merkle-tree-go/blob/main/docs/main_run.txt)

## Dump

TreeMarshal: [tree.json](https://github.com/FantasyJony/openzeppelin-merkle-tree-go/blob/main/docs/dump_tree.json)

TreeProofMarshal:[proof.json](https://github.com/FantasyJony/openzeppelin-merkle-tree-go/blob/main/docs/dump_proof.json)


## Solidity

### hardhat
```
cd merkle-tree-contracts
npm install
npx hardhat test
```

```
  StandardMerkleTree
    Merkle Tree
      ✔ verify (597ms)
      ✔ verifyMultiProof


  2 passing (613ms)
```

contracts: [StandardMerkleTree.sol](https://github.com/FantasyJony/openzeppelin-merkle-tree-go/blob/main/merkle-tree-contracts/contracts/StandardMerkleTree.sol)

test: [StandardMerkleTree.ts](https://github.com/FantasyJony/openzeppelin-merkle-tree-go/blob/main/merkle-tree-contracts/test/StandardMerkleTree.ts)

## API

### `import smt`

```go
import (
	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	"github.com/ethereum/go-ethereum/common/hexutil"
)
```

### `smt.Of`

```go
// Of
tree, err := smt.Of(
    [][]interface{}{
        {
            smt.SolAddress("0x1111111111111111111111111111111111111111"),
            smt.SolNumber("5000000000000000000"),
        },
        {
            smt.SolAddress("0x2222222222222222222222222222222222222222"),
            smt.SolNumber("2500000000000000000"),
        },
    },
    []string {
        smt.SOL_ADDRESS,
        smt.SOL_UINT256,
    },
)
```

### `smt.CreateTree`

```go
tree, err := smt.CreateTree([]string{smt.SOL_ADDRESS, smt.SOL_UINT256})
```

### `smt.Verify`

```go
// root
rootHash, err := hexutil.Decode("0xd4dee0beab2d53f2cc83e567171bd2820e49898130a22622b10ead383e90bd77")
if err != nil {
    fmt.Println(err)
}

// proof
pv, err := hexutil.Decode("0xb92c48e9d7abe27fd8dfd6b5dfdbfb1c9a463f80c712b66f3a5180a090cccafc")
proof := [][]byte{
    pv,
}


// leafEncodings
leafEncodings := []string {
    smt.SOL_ADDRESS,
    smt.SOL_UINT256,
}

// value
value := []interface{}{
    smt.SolAddress("0x1111111111111111111111111111111111111111"),
    smt.SolNumber("5000000000000000000"),
}

// leaf
leafHah, err := LeafHash(leafEncodings, value)
if err != nil {
    fmt.Println(err)
}

// Verify
verified, err := smt.Verify(rootHash, leafHah, proof)
```

### `smt.VerifyMultiProof`

```go
// root
root, err := hexutil.Decode("0xd4dee0beab2d53f2cc83e567171bd2820e49898130a22622b10ead383e90bd77")
if err != nil {
    fmt.Println(err)
}

leafEncodings := []string{
    smt.SOL_ADDRESS,
    smt.SOL_UINT256,
}

value := []interface{}{
    smt.SolAddress("0x1111111111111111111111111111111111111111"),
    smt.SolNumber("5000000000000000000"),
}

// leaf
leaf1, err := smt,LeafHash(leafEncodings, value)
if err != nil {
    fmt.Println(err)
}

leaves := [][]byte{
    leaf1,
}

// proof
proofValue01, err := hexutil.Decode("0xb92c48e9d7abe27fd8dfd6b5dfdbfb1c9a463f80c712b66f3a5180a090cccafc")
if err != nil {
    fmt.Println(err)
}
proof := [][]byte{
    proofValue01,
}

// proofFlags
proofFlags := []bool{
    false,
}

// multiProof
multiProof := &smt.MultiProof{
    Proof:      proof,
    ProofFlags: proofFlags,
    Leaves:     leaves,
}

// VerifyMultiProof
verified, err := smt.VerifyMultiProof(root, multiProof)
```

### `smt.Load`

```go
value := "{\"format\":\"standard-v1\",\"tree\":[\"0xd4dee0beab2d53f2cc83e567171bd2820e49898130a22622b10ead383e90bd77\",\"0xeb02c421cfa48976e66dfb29120745909ea3a0f843456c263cf8f1253483e283\",\"0xb92c48e9d7abe27fd8dfd6b5dfdbfb1c9a463f80c712b66f3a5180a090cccafc\"],\"values\":[{\"value\":[\"0x1111111111111111111111111111111111111111\",\"5000000000000000000\"],\"treeIndex\":1},{\"value\":[\"0x2222222222222222222222222222222222222222\",\"2500000000000000000\"],\"treeIndex\":2}],\"leafEncoding\":[\"address\",\"uint256\"]}"
tree , err := smt.Load([]byte(value))
```

### `tree.AddLeaf`

```go
value := []interface{}{
    smt.SolAddress("0x1111111111111111111111111111111111111111"),
    smt.SolNumber("5000000000000000000"),
}
leafHash, err := tree.AddLeaf(value)
```

### `tree.MakeTree`
```go
rootHash, err := tree.MakeTree()
```

### `tree.GetRoot`

```go
rootHash := tree.GetRoot()
rootHashValue , err := hexutil.Encode(rootHash)
```

### `tree.Dump` or `tree.TreeMarshal`
```go
treeData := tree.Dump()
jsonValue, err := json.Marshal(treeData)
```
or
```go
jsonValue, err := tree.TreeMarshal()
```

### `tree.GetProof`
```go
value := []interface{}{
    smt.SolAddress("0x1111111111111111111111111111111111111111"),
    smt.SolNumber("5000000000000000000"),
}
proof, err := tree.GetProof(value)
```
### `tree.GetProofWithIndex`
```go
proof, err := tree.GetProofWithIndex(0)
```

### `tree.GetMultiProof`

```go
value := []interface{}{
    smt.SolAddress("0x1111111111111111111111111111111111111111"),
    smt.SolNumber("5000000000000000000"),
}

leaves := [][]interface{}{
    value,
}

multiProof, err := tree.GetMultiProof(leaves)
```

### `tree.GetMultiProofWithIndices`
```go
indices := []int{
    0,
}
multiProof, err := tree.GetMultiProofWithIndices(indices)
```

### `tree.Verify`

```go
value := []interface{}{
    smt.SolAddress("0x1111111111111111111111111111111111111111"),
    smt.SolNumber("5000000000000000000"),
}
verified, err := tree.Verify(proof, value)
```

### `tree.VerifyWithIndex`
```go
verified, err := tree.VerifyWithIndex(proof, 0)
```

### `tree.VerifyMultiProof`
```go
verified, err := tree.VerifyMultiProof(multiProof)
```

### `tree.Entries`
```go
entries := tree.Entries()
for k, v := range entries {
    fmt.Println("ValueIndex:", v.ValueIndex, " TreeIndex:", v.TreeIndex, " Hash:", hexutil.Encode(v.Hash), " Value:", v.Value)
    proof, err := tree.GetProofWithIndex(k)
    for pk, pv := range proof {
        fmt.Println("Proof k:", pk, " v:", hexutil.Encode(pv))
    }
}
```

### `tree.LeafHash`

```go
value := []interface{}{
    smt.SolAddress("0x1111111111111111111111111111111111111111"),
    smt.SolNumber("5000000000000000000"),
}
leafHash, err := tree.LeafHash(value)
```

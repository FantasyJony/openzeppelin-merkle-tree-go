# Openzeppelin Standard Tree Go

A Go library to generate merkle trees and merkle proofs.

[![Version](https://img.shields.io/badge/version-v1.0.6-blue)](https://github.com/FantasyJony/openzeppelin-merkle-tree-go/releases/tag/v1.0.4)
[![PkgGoDev](https://pkg.go.dev/badge/ggithub.com/FantasyJony/openzeppelin-merkle-tree-go/logrus.svg)](https://pkg.go.dev/github.com/FantasyJony/openzeppelin-merkle-tree-go)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## @openzeppelin/merkle-tree

fork from [@openzeppelin/merkle-tree](https://github.com/OpenZeppelin/merkle-tree)

contracts doc [MerkleProof](https://docs.openzeppelin.com/contracts/4.x/api/utils#MerkleProof)

## Installation
```
go get github.com/FantasyJony/openzeppelin-merkle-tree-go@v1.0.6
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

```

## API & Examples

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


## Dump

### `TreeMarshal`

```json
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

### `TreeProofMarshal`
```json
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

## Solidity Leaf Hash

```solidity
bytes32 leaf = keccak256(bytes.concat(keccak256(abi.encode(addr, amount))));
```

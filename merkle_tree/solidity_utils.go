package merkle_tree

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/sha3"
)

func Keccak256(value []byte) ([]byte, error) {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(value)
	return hash.Sum(nil), nil
}

func AbiPack(types []string, values ...interface{}) ([]byte, error) {
	var args abi.Arguments
	for _, v := range types {
		typ, err := abi.NewType(v, "string", nil)
		if err != nil {
			return nil, err
		}
		args = append(args, abi.Argument{
			Type: typ,
		})
	}
	packed, err := args.Pack(values...)
	return packed, err
}

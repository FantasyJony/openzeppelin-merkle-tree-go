package standard_tree

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func SolAddress(value string) common.Address {
	return common.HexToAddress(value)
}

func SolNumber(value string) *big.Int {
	bigNumber := big.NewInt(0)
	bigNumber.SetString(value, 10)
	return bigNumber
}

func SolBytes(value string) []byte {
	r, _ := hex.DecodeString(value)
	return r
}

func SolString(value string) []byte {
	return SolBytes(value)
}

func SolBool(value bool) bool {
	return value
}

func SolValueUnmarshal(value interface{}, leafEncoding string) interface{} {
	switch leafEncoding {
	case SOL_ADDRESS:
		{
			instance, _ := value.(string)
			return SolAddress(instance)
		}
	case SOL_UINT8, SOL_UINT16, SOL_UINT32, SOL_UINT64, SOL_UINT128, SOL_UINT256, SOL_INT8, SOL_INT16, SOL_INT32, SOL_INT64, SOL_INT128, SOL_INT256:
		{
			instance, _ := value.(string)
			return SolNumber(instance)
		}
	case SOL_STRING, SOL_BYTES, SOL_BYTES1, SOL_BYTES2, SOL_BYTES3, SOL_BYTES4, SOL_BYTES5, SOL_BYTES6, SOL_BYTES7, SOL_BYTES8, SOL_BYTES9, SOL_BYTES10, SOL_BYTES11, SOL_BYTES12, SOL_BYTES13, SOL_BYTES14, SOL_BYTES15, SOL_BYTES16, SOL_BYTES17, SOL_BYTES18, SOL_BYTES19, SOL_BYTES20, SOL_BYTES21, SOL_BYTES22, SOL_BYTES23, SOL_BYTES24, SOL_BYTES25, SOL_BYTES26, SOL_BYTES27, SOL_BYTES28, SOL_BYTES29, SOL_BYTES30, SOL_BYTES31, SOL_BYTES32:
		{
			instance, _ := value.(string)
			return SolBytes(instance)
		}
	case SOL_BOOL:
		{
			instance, _ := value.(bool)
			return SolBool(instance)
		}
	}
	return ""
}

func SolValueMarshal(value interface{}, leafEncoding string) interface{} {
	switch leafEncoding {
	case SOL_ADDRESS:
		{
			instance, _ := value.(common.Address)
			return instance.Hex()
		}
	case SOL_UINT8, SOL_UINT16, SOL_UINT32, SOL_UINT64, SOL_UINT128, SOL_UINT256, SOL_INT8, SOL_INT16, SOL_INT32, SOL_INT64, SOL_INT128, SOL_INT256:
		{
			instance, _ := value.(*big.Int)
			return instance.String()
		}
	case SOL_STRING, SOL_BYTES, SOL_BYTES1, SOL_BYTES2, SOL_BYTES3, SOL_BYTES4, SOL_BYTES5, SOL_BYTES6, SOL_BYTES7, SOL_BYTES8, SOL_BYTES9, SOL_BYTES10, SOL_BYTES11, SOL_BYTES12, SOL_BYTES13, SOL_BYTES14, SOL_BYTES15, SOL_BYTES16, SOL_BYTES17, SOL_BYTES18, SOL_BYTES19, SOL_BYTES20, SOL_BYTES21, SOL_BYTES22, SOL_BYTES23, SOL_BYTES24, SOL_BYTES25, SOL_BYTES26, SOL_BYTES27, SOL_BYTES28, SOL_BYTES29, SOL_BYTES30, SOL_BYTES31, SOL_BYTES32:
		{
			instance, _ := value.([]byte)
			return hex.EncodeToString(instance)
		}
	case SOL_BOOL:
		{
			instance, _ := value.(bool)
			return instance
		}
	}
	return ""
}

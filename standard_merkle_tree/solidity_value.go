package standard_merkle_tree

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func SolAddress(value string) common.Address {
	return common.HexToAddress(value)
}

func SolAddressArray(value []interface{}) []common.Address {
	list := make([]common.Address, len(value))
	for k, v := range value {
		list[k] = SolAddress(v.(string))
	}
	return list
}

func SolNumber(value string) *big.Int {
	bigNumber := big.NewInt(0)
	bigNumber.SetString(value, 10)
	return bigNumber
}

func SolNumberArray(value []interface{}) []*big.Int {
	list := make([]*big.Int, len(value))
	for k, v := range value {
		list[k] = SolNumber(v.(string))
	}
	return list
}

func SolBytes(value string) []byte {
	r, _ := hex.DecodeString(value)
	return r
}

func SolBytesArray(value []interface{}) [][]byte {
	list := make([][]byte, len(value))
	for k, v := range value {
		list[k] = SolBytes(v.(string))
	}
	return list
}

func SolString(value string) []byte {
	return SolBytes(value)
}

func SolStringArray(value []interface{}) [][]byte {
	return SolBytesArray(value)
}

func SolBool(value bool) bool {
	return value
}

func SolBoolArray(value []interface{}) []bool {
	list := make([]bool, len(value))
	for k, v := range value {
		list[k] = v.(bool)
	}
	return list
}

func SolValueUnmarshal(value interface{}, leafEncoding string) interface{} {
	switch leafEncoding {
	case SOL_ADDRESS:
		{
			instance, _ := value.(string)
			return SolAddress(instance)
		}
	case SOL_ADDRESS_ARRAY:
		{
			instance, _ := value.([]interface{})
			return SolAddressArray(instance)
		}
	case SOL_UINT8, SOL_UINT16, SOL_UINT32, SOL_UINT64, SOL_UINT128, SOL_UINT256, SOL_INT8, SOL_INT16, SOL_INT32, SOL_INT64, SOL_INT128, SOL_INT256:
		{
			instance, _ := value.(string)
			return SolNumber(instance)
		}
	case SOL_UINT8_ARRAY, SOL_UINT16_ARRAY, SOL_UINT32_ARRAY, SOL_UINT64_ARRAY, SOL_UINT128_ARRAY, SOL_UINT256_ARRAY, SOL_INT8_ARRAY, SOL_INT16_ARRAY, SOL_INT32_ARRAY, SOL_INT64_ARRAY, SOL_INT128_ARRAY, SOL_INT256_ARRAY:
		{
			instance, _ := value.([]interface{})
			return SolNumberArray(instance)
		}
	case SOL_STRING, SOL_BYTES, SOL_BYTES1, SOL_BYTES2, SOL_BYTES3, SOL_BYTES4, SOL_BYTES5, SOL_BYTES6, SOL_BYTES7, SOL_BYTES8, SOL_BYTES9, SOL_BYTES10, SOL_BYTES11, SOL_BYTES12, SOL_BYTES13, SOL_BYTES14, SOL_BYTES15, SOL_BYTES16, SOL_BYTES17, SOL_BYTES18, SOL_BYTES19, SOL_BYTES20, SOL_BYTES21, SOL_BYTES22, SOL_BYTES23, SOL_BYTES24, SOL_BYTES25, SOL_BYTES26, SOL_BYTES27, SOL_BYTES28, SOL_BYTES29, SOL_BYTES30, SOL_BYTES31, SOL_BYTES32:
		{
			instance, _ := value.(string)
			return SolBytes(instance)
		}
	case SOL_BYTES32_ARRAY:
		{
			instance, _ := value.([]interface{})
			return SolBytesArray(instance)
		}
	case SOL_BOOL:
		{
			instance, _ := value.(bool)
			return SolBool(instance)
		}
	case SOL_BOOL_ARRAY:
		{
			instance, _ := value.([]interface{})
			return SolBoolArray(instance)
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
	case SOL_ADDRESS_ARRAY:
		{
			instance, _ := value.([]common.Address)
			list := make([]string, len(instance))
			for k, v := range instance {
				list[k] = v.Hex()
			}
			return list
		}
	case SOL_UINT8, SOL_UINT16, SOL_UINT32, SOL_UINT64, SOL_UINT128, SOL_UINT256, SOL_INT8, SOL_INT16, SOL_INT32, SOL_INT64, SOL_INT128, SOL_INT256:
		{
			instance, _ := value.(*big.Int)
			return instance.String()
		}
	case SOL_UINT8_ARRAY, SOL_UINT16_ARRAY, SOL_UINT32_ARRAY, SOL_UINT64_ARRAY, SOL_UINT128_ARRAY, SOL_UINT256_ARRAY, SOL_INT8_ARRAY, SOL_INT16_ARRAY, SOL_INT32_ARRAY, SOL_INT64_ARRAY, SOL_INT128_ARRAY, SOL_INT256_ARRAY:
		{
			instance, _ := value.([]*big.Int)
			list := make([]string, len(instance))
			for k, v := range instance {
				list[k] = v.String()
			}
			return list
		}
	case SOL_STRING, SOL_BYTES, SOL_BYTES1, SOL_BYTES2, SOL_BYTES3, SOL_BYTES4, SOL_BYTES5, SOL_BYTES6, SOL_BYTES7, SOL_BYTES8, SOL_BYTES9, SOL_BYTES10, SOL_BYTES11, SOL_BYTES12, SOL_BYTES13, SOL_BYTES14, SOL_BYTES15, SOL_BYTES16, SOL_BYTES17, SOL_BYTES18, SOL_BYTES19, SOL_BYTES20, SOL_BYTES21, SOL_BYTES22, SOL_BYTES23, SOL_BYTES24, SOL_BYTES25, SOL_BYTES26, SOL_BYTES27, SOL_BYTES28, SOL_BYTES29, SOL_BYTES30, SOL_BYTES31, SOL_BYTES32:
		{
			instance, _ := value.([]byte)
			return hex.EncodeToString(instance)
		}
	case SOL_BYTES32_ARRAY:
		{
			instance, _ := value.([]string)
			list := make([]string, len(instance))
			for k, v := range instance {
				list[k] = hex.EncodeToString([]byte(v))
			}
			return list
		}
	case SOL_BOOL:
		{
			instance, _ := value.(bool)
			return instance
		}
	case SOL_BOOL_ARRAY:
		{
			instance, _ := value.([]bool)
			return instance
		}
	}
	return ""
}

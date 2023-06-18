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

func ToSolValue(value interface{}, leafEncoding string) interface{} {

	switch leafEncoding {
	case SOL_ADDRESS:
		{
			instance, ok := value.(string)
			if ok {
				return SolAddress(instance)
			}
		}
	case SOL_ADDRESS_ARRAY:
		{
			instance, ok := value.([]interface{})
			if ok {
				return SolAddressArray(instance)
			}
		}

	case SOL_INT8, SOL_INT16, SOL_INT24, SOL_INT32, SOL_INT40, SOL_INT48, SOL_INT56, SOL_INT64, SOL_INT72, SOL_INT80, SOL_INT88, SOL_INT96, SOL_INT104, SOL_INT112, SOL_INT120, SOL_INT128, SOL_INT136, SOL_INT144, SOL_INT152, SOL_INT160, SOL_INT168, SOL_INT176, SOL_INT184, SOL_INT192, SOL_INT200, SOL_INT208, SOL_INT216, SOL_INT224, SOL_INT232, SOL_INT240, SOL_INT248, SOL_INT256:
		{
			instance, ok := value.(string)
			if ok {
				return SolNumber(instance)
			}
		}
	case SOL_INT8_ARRAY, SOL_INT16_ARRAY, SOL_INT24_ARRAY, SOL_INT32_ARRAY, SOL_INT40_ARRAY, SOL_INT48_ARRAY, SOL_INT56_ARRAY, SOL_INT64_ARRAY, SOL_INT72_ARRAY, SOL_INT80_ARRAY, SOL_INT88_ARRAY, SOL_INT96_ARRAY, SOL_INT104_ARRAY, SOL_INT112_ARRAY, SOL_INT120_ARRAY, SOL_INT128_ARRAY, SOL_INT136_ARRAY, SOL_INT144_ARRAY, SOL_INT152_ARRAY, SOL_INT160_ARRAY, SOL_INT168_ARRAY, SOL_INT176_ARRAY, SOL_INT184_ARRAY, SOL_INT192_ARRAY, SOL_INT200_ARRAY, SOL_INT208_ARRAY, SOL_INT216_ARRAY, SOL_INT224_ARRAY, SOL_INT232_ARRAY, SOL_INT240_ARRAY, SOL_INT248_ARRAY, SOL_INT256_ARRAY:
		{
			instance, ok := value.([]interface{})
			if ok {
				return SolNumberArray(instance)
			}
		}
	case SOL_UINT8, SOL_UINT16, SOL_UINT24, SOL_UINT32, SOL_UINT40, SOL_UINT48, SOL_UINT56, SOL_UINT64, SOL_UINT72, SOL_UINT80, SOL_UINT88, SOL_UINT96, SOL_UINT104, SOL_UINT112, SOL_UINT120, SOL_UINT128, SOL_UINT136, SOL_UINT144, SOL_UINT152, SOL_UINT160, SOL_UINT168, SOL_UINT176, SOL_UINT184, SOL_UINT192, SOL_UINT200, SOL_UINT208, SOL_UINT216, SOL_UINT224, SOL_UINT232, SOL_UINT240, SOL_UINT248, SOL_UINT256:
		{
			instance, ok := value.(string)
			if ok {
				return SolNumber(instance)
			}
		}
	case SOL_UINT8_ARRAY, SOL_UINT16_ARRAY, SOL_UINT24_ARRAY, SOL_UINT32_ARRAY, SOL_UINT40_ARRAY, SOL_UINT48_ARRAY, SOL_UINT56_ARRAY, SOL_UINT64_ARRAY, SOL_UINT72_ARRAY, SOL_UINT80_ARRAY, SOL_UINT88_ARRAY, SOL_UINT96_ARRAY, SOL_UINT104_ARRAY, SOL_UINT112_ARRAY, SOL_UINT120_ARRAY, SOL_UINT128_ARRAY, SOL_UINT136_ARRAY, SOL_UINT144_ARRAY, SOL_UINT152_ARRAY, SOL_UINT160_ARRAY, SOL_UINT168_ARRAY, SOL_UINT176_ARRAY, SOL_UINT184_ARRAY, SOL_UINT192_ARRAY, SOL_UINT200_ARRAY, SOL_UINT208_ARRAY, SOL_UINT216_ARRAY, SOL_UINT224_ARRAY, SOL_UINT232_ARRAY, SOL_UINT240_ARRAY, SOL_UINT248_ARRAY, SOL_UINT256_ARRAY:
		{
			instance, ok := value.([]interface{})
			if ok {
				return SolNumberArray(instance)
			}
		}
	case SOL_BYTES1, SOL_BYTES2, SOL_BYTES3, SOL_BYTES4, SOL_BYTES5, SOL_BYTES6, SOL_BYTES7, SOL_BYTES8, SOL_BYTES9, SOL_BYTES10, SOL_BYTES11, SOL_BYTES12, SOL_BYTES13, SOL_BYTES14, SOL_BYTES15, SOL_BYTES16, SOL_BYTES17, SOL_BYTES18, SOL_BYTES19, SOL_BYTES20, SOL_BYTES21, SOL_BYTES22, SOL_BYTES23, SOL_BYTES24, SOL_BYTES25, SOL_BYTES26, SOL_BYTES27, SOL_BYTES28, SOL_BYTES29, SOL_BYTES30, SOL_BYTES31, SOL_BYTES32:
		{
			instance, _ := value.(string)
			return SolBytes(instance)
		}
	case SOL_BYTES1_ARRAY, SOL_BYTES2_ARRAY, SOL_BYTES3_ARRAY, SOL_BYTES4_ARRAY, SOL_BYTES5_ARRAY, SOL_BYTES6_ARRAY, SOL_BYTES7_ARRAY, SOL_BYTES8_ARRAY, SOL_BYTES9_ARRAY, SOL_BYTES10_ARRAY, SOL_BYTES11_ARRAY, SOL_BYTES12_ARRAY, SOL_BYTES13_ARRAY, SOL_BYTES14_ARRAY, SOL_BYTES15_ARRAY, SOL_BYTES16_ARRAY, SOL_BYTES17_ARRAY, SOL_BYTES18_ARRAY, SOL_BYTES19_ARRAY, SOL_BYTES20_ARRAY, SOL_BYTES21_ARRAY, SOL_BYTES22_ARRAY, SOL_BYTES23_ARRAY, SOL_BYTES24_ARRAY, SOL_BYTES25_ARRAY, SOL_BYTES26_ARRAY, SOL_BYTES27_ARRAY, SOL_BYTES28_ARRAY, SOL_BYTES29_ARRAY, SOL_BYTES30_ARRAY, SOL_BYTES31_ARRAY, SOL_BYTES32_ARRAY:
		{
			instance, ok := value.([]interface{})
			if ok {
				return SolBytesArray(instance)
			}
		}
	case SOL_STRING, SOL_BYTES:
		{
			instance, ok := value.(string)
			if ok {
				return SolBytes(instance)
			}
		}
	case SOL_STRING_ARRAY, SOL_BYTES_ARRAY:
		{
			instance, ok := value.([]interface{})
			if ok {
				return SolBytesArray(instance)
			}
		}
	case SOL_BOOL:
		{
			instance, ok := value.(bool)
			if ok {
				return SolBool(instance)
			}
		}
	case SOL_BOOL_ARRAY:
		{
			instance, ok := value.([]interface{})
			if ok {
				return SolBoolArray(instance)
			}
		}
	}

	return nil
}

func ToJsonValue(value interface{}, leafEncoding string) interface{} {
	switch leafEncoding {
	case SOL_ADDRESS:
		{
			instance, ok := value.(common.Address)
			if ok {
				return instance.Hex()
			}
		}
	case SOL_ADDRESS_ARRAY:
		{
			instance, ok := value.([]common.Address)
			if ok {
				list := make([]string, len(instance))
				for k, v := range instance {
					list[k] = v.Hex()
				}
				return list
			}
		}
	case SOL_INT8, SOL_INT16, SOL_INT24, SOL_INT32, SOL_INT40, SOL_INT48, SOL_INT56, SOL_INT64, SOL_INT72, SOL_INT80, SOL_INT88, SOL_INT96, SOL_INT104, SOL_INT112, SOL_INT120, SOL_INT128, SOL_INT136, SOL_INT144, SOL_INT152, SOL_INT160, SOL_INT168, SOL_INT176, SOL_INT184, SOL_INT192, SOL_INT200, SOL_INT208, SOL_INT216, SOL_INT224, SOL_INT232, SOL_INT240, SOL_INT248, SOL_INT256:
		{
			instance, ok := value.(*big.Int)
			if ok {
				return instance.String()
			}
		}
	case SOL_INT8_ARRAY, SOL_INT16_ARRAY, SOL_INT24_ARRAY, SOL_INT32_ARRAY, SOL_INT40_ARRAY, SOL_INT48_ARRAY, SOL_INT56_ARRAY, SOL_INT64_ARRAY, SOL_INT72_ARRAY, SOL_INT80_ARRAY, SOL_INT88_ARRAY, SOL_INT96_ARRAY, SOL_INT104_ARRAY, SOL_INT112_ARRAY, SOL_INT120_ARRAY, SOL_INT128_ARRAY, SOL_INT136_ARRAY, SOL_INT144_ARRAY, SOL_INT152_ARRAY, SOL_INT160_ARRAY, SOL_INT168_ARRAY, SOL_INT176_ARRAY, SOL_INT184_ARRAY, SOL_INT192_ARRAY, SOL_INT200_ARRAY, SOL_INT208_ARRAY, SOL_INT216_ARRAY, SOL_INT224_ARRAY, SOL_INT232_ARRAY, SOL_INT240_ARRAY, SOL_INT248_ARRAY, SOL_INT256_ARRAY:
		{
			instance, ok := value.([]*big.Int)
			if ok {
				list := make([]string, len(instance))
				for k, v := range instance {
					list[k] = v.String()
				}
				return list
			}

		}
	case SOL_UINT8, SOL_UINT16, SOL_UINT24, SOL_UINT32, SOL_UINT40, SOL_UINT48, SOL_UINT56, SOL_UINT64, SOL_UINT72, SOL_UINT80, SOL_UINT88, SOL_UINT96, SOL_UINT104, SOL_UINT112, SOL_UINT120, SOL_UINT128, SOL_UINT136, SOL_UINT144, SOL_UINT152, SOL_UINT160, SOL_UINT168, SOL_UINT176, SOL_UINT184, SOL_UINT192, SOL_UINT200, SOL_UINT208, SOL_UINT216, SOL_UINT224, SOL_UINT232, SOL_UINT240, SOL_UINT248, SOL_UINT256:
		{
			instance, ok := value.(*big.Int)
			if ok {
				return instance.String()
			}
		}
	case SOL_UINT8_ARRAY, SOL_UINT16_ARRAY, SOL_UINT24_ARRAY, SOL_UINT32_ARRAY, SOL_UINT40_ARRAY, SOL_UINT48_ARRAY, SOL_UINT56_ARRAY, SOL_UINT64_ARRAY, SOL_UINT72_ARRAY, SOL_UINT80_ARRAY, SOL_UINT88_ARRAY, SOL_UINT96_ARRAY, SOL_UINT104_ARRAY, SOL_UINT112_ARRAY, SOL_UINT120_ARRAY, SOL_UINT128_ARRAY, SOL_UINT136_ARRAY, SOL_UINT144_ARRAY, SOL_UINT152_ARRAY, SOL_UINT160_ARRAY, SOL_UINT168_ARRAY, SOL_UINT176_ARRAY, SOL_UINT184_ARRAY, SOL_UINT192_ARRAY, SOL_UINT200_ARRAY, SOL_UINT208_ARRAY, SOL_UINT216_ARRAY, SOL_UINT224_ARRAY, SOL_UINT232_ARRAY, SOL_UINT240_ARRAY, SOL_UINT248_ARRAY, SOL_UINT256_ARRAY:
		{
			instance, ok := value.([]*big.Int)
			if ok {
				list := make([]string, len(instance))
				for k, v := range instance {
					list[k] = v.String()
				}
				return list
			}
		}
	case SOL_BYTES1, SOL_BYTES2, SOL_BYTES3, SOL_BYTES4, SOL_BYTES5, SOL_BYTES6, SOL_BYTES7, SOL_BYTES8, SOL_BYTES9, SOL_BYTES10, SOL_BYTES11, SOL_BYTES12, SOL_BYTES13, SOL_BYTES14, SOL_BYTES15, SOL_BYTES16, SOL_BYTES17, SOL_BYTES18, SOL_BYTES19, SOL_BYTES20, SOL_BYTES21, SOL_BYTES22, SOL_BYTES23, SOL_BYTES24, SOL_BYTES25, SOL_BYTES26, SOL_BYTES27, SOL_BYTES28, SOL_BYTES29, SOL_BYTES30, SOL_BYTES31, SOL_BYTES32:
		{
			instance, ok := value.(string)
			if ok {
				return SolBytes(instance)
			}
		}
	case SOL_BYTES1_ARRAY, SOL_BYTES2_ARRAY, SOL_BYTES3_ARRAY, SOL_BYTES4_ARRAY, SOL_BYTES5_ARRAY, SOL_BYTES6_ARRAY, SOL_BYTES7_ARRAY, SOL_BYTES8_ARRAY, SOL_BYTES9_ARRAY, SOL_BYTES10_ARRAY, SOL_BYTES11_ARRAY, SOL_BYTES12_ARRAY, SOL_BYTES13_ARRAY, SOL_BYTES14_ARRAY, SOL_BYTES15_ARRAY, SOL_BYTES16_ARRAY, SOL_BYTES17_ARRAY, SOL_BYTES18_ARRAY, SOL_BYTES19_ARRAY, SOL_BYTES20_ARRAY, SOL_BYTES21_ARRAY, SOL_BYTES22_ARRAY, SOL_BYTES23_ARRAY, SOL_BYTES24_ARRAY, SOL_BYTES25_ARRAY, SOL_BYTES26_ARRAY, SOL_BYTES27_ARRAY, SOL_BYTES28_ARRAY, SOL_BYTES29_ARRAY, SOL_BYTES30_ARRAY, SOL_BYTES31_ARRAY, SOL_BYTES32_ARRAY:
		{
			instance, ok := value.([]interface{})
			if ok {
				return SolBytesArray(instance)
			}
		}
	case SOL_STRING, SOL_BYTES:
		{
			instance, ok := value.(string)
			if ok {
				return SolBytes(instance)
			}
		}
	case SOL_STRING_ARRAY, SOL_BYTES_ARRAY:
		{
			instance, ok := value.([]interface{})
			if ok {
				return SolBytesArray(instance)
			}
		}
	case SOL_BOOL:
		{
			instance, ok := value.(bool)
			if ok {
				return instance
			}
		}
	case SOL_BOOL_ARRAY:
		{
			instance, ok := value.([]bool)
			if ok {
				return instance
			}
		}
	}
	return ""
}

// https://github.com/ethereum/go-ethereum/blob/master/accounts/abi/type_test.go
func abiArgConvert(types []string, values ...interface{}) []interface{} {
	//convertValue := make([]interface{}, len(values))
	for k, v := range types {
		switch v {
		case SOL_UINT8:
			{
				instance, ok := values[k].(*big.Int)
				if ok {
					values[k] = uint8(instance.Uint64())
				}
			}
			break
		case SOL_UINT16:
			{
				instance, ok := values[k].(*big.Int)
				if ok {
					values[k] = uint16(instance.Uint64())
				}
			}
			break
		case SOL_UINT32:
			{
				instance, ok := values[k].(*big.Int)
				if ok {
					values[k] = uint32(instance.Uint64())
				}
			}
			break
		case SOL_UINT64:
			{
				instance, ok := values[k].(*big.Int)
				if ok {
					values[k] = instance.Uint64()
				}
			}
			break
		case SOL_INT8:
			{
				instance, ok := values[k].(*big.Int)
				if ok {
					values[k] = int8(instance.Int64())
				}
			}
			break
		case SOL_INT16:
			{
				instance, ok := values[k].(*big.Int)
				if ok {
					values[k] = int16(instance.Int64())
				}
			}
			break
		case SOL_INT32:
			{
				instance, ok := values[k].(*big.Int)
				if ok {
					values[k] = int32(instance.Int64())
				}
			}
			break
		case SOL_INT64:
			{
				instance, ok := values[k].(*big.Int)
				if ok {
					values[k] = instance.Int64()
				}
			}
			break
		case SOL_UINT8_ARRAY:
			{
				instance, ok := values[k].([]*big.Int)
				if ok {
					list := make([]uint8, len(instance))
					for ik, iv := range instance {
						list[ik] = uint8(iv.Uint64())
					}
					values[k] = list
				}
			}
			break
		case SOL_UINT16_ARRAY:
			{
				instance, ok := values[k].([]*big.Int)
				if ok {
					list := make([]uint16, len(instance))
					for ik, iv := range instance {
						list[ik] = uint16(iv.Uint64())
					}
					values[k] = list
				}
			}
			break
		case SOL_UINT32_ARRAY:
			{
				instance, ok := values[k].([]*big.Int)
				if ok {
					list := make([]uint32, len(instance))
					for ik, iv := range instance {
						list[ik] = uint32(iv.Uint64())
					}
					values[k] = list
				}
			}
			break
		case SOL_UINT64_ARRAY:
			{
				instance, ok := values[k].([]*big.Int)
				if ok {
					list := make([]uint64, len(instance))
					for ik, iv := range instance {
						list[ik] = uint64(iv.Uint64())
					}
					values[k] = list
				}
			}
			break
		case SOL_INT8_ARRAY:
			{
				instance, ok := values[k].([]*big.Int)
				if ok {
					list := make([]int8, len(instance))
					for ik, iv := range instance {
						list[ik] = int8(iv.Int64())
					}
					values[k] = list
				}
			}
			break
		case SOL_INT16_ARRAY:
			{
				instance, ok := values[k].([]*big.Int)
				if ok {
					list := make([]int16, len(instance))
					for ik, iv := range instance {
						list[ik] = int16(iv.Int64())
					}
					values[k] = list
				}
			}
			break
		case SOL_INT32_ARRAY:
			{
				instance, ok := values[k].([]*big.Int)
				if ok {
					list := make([]int32, len(instance))
					for ik, iv := range instance {
						list[ik] = int32(iv.Int64())
					}
					values[k] = list
				}
			}
			break
		case SOL_INT64_ARRAY:
			{
				instance, ok := values[k].([]*big.Int)
				if ok {
					list := make([]int64, len(instance))
					for ik, iv := range instance {
						list[ik] = int64(iv.Int64())
					}
					values[k] = list
				}
			}
			break
		}
	}
	return values
}

package util

import (
	"math"
	"reflect"
)

type Int interface {
	int | int8 | int16 | int32 | int64
}

type Uint interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

type Slice[T string | Int | Uint | Float] []T

func (s Slice[T]) Contains(search T) bool {
	for _, value := range s {
		if value == search {
			return true
		}
	}
	return false
}

func TransformInterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

// SpiltList 切割固定大小的 切片
func SpiltList[T any](list []T, size int) (dataList [][]T) {
	lens := len(list)
	mod := math.Ceil(float64(lens) / float64(size))
	//fmt.Println(mod) // 3
	//dataList := make([][]T, 0)
	for i := 0; i < int(mod); i++ {
		tmpList := make([]T, 0, size)
		//fmt.Println("i=", i)
		if i == int(mod)-1 {
			tmpList = list[i*size:]
		} else {
			tmpList = list[i*size : i*size+size]
		}
		dataList = append(dataList, tmpList)
	}
	//for i, sp := range dataList {
	//	fmt.Println(i, " ==> ", sp)
	//}
	return
}

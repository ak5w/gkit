package debug

import (
	"fmt"
	"reflect"
)

func isSlice(v interface{}) bool {
	value := reflect.ValueOf(v)
	kind := value.Kind()
	return kind == reflect.Slice
}

func Slice(s interface{}) {

	// 获取切片值
	value := reflect.ValueOf(s)
	// 检查是否为切片类型
	if value.Kind() == reflect.Slice {
		// 获取切片长度
		length := value.Len()
		cap := value.Cap()
		// 获取切片元素类型
		elementType := value.Type().Elem()
		fmt.Printf("type=%v,len=%d, cap=%d, slice=%v\n", elementType, length, cap, s) //
	} else {
		fmt.Println("Type is not slice.")
	}
}

package libs

import "reflect"

//把任意切片类型转化为[]interface{}切片类型
func CreateAnyTypeSlice(slice interface{}) []interface{} {

	val, ok := isSlice(slice)

	if !ok {
		return nil
	}

	sliceLen := val.Len()

	out := make([]interface{}, sliceLen)

	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}

	return out
}

//判断是否为切片类型
func isSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)

	if val.Kind() == reflect.Slice {
		ok = true
	}

	return
}

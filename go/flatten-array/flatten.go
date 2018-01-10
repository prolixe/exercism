package flatten

/*
import "fmt"
import "reflect"
import "unsafe"

func Flatten(i interface{}) []interface{} {

	output := make([]interface{}, 0)

	v := reflect.ValueOf(i)
	fmt.Printf("Value i is %#v of type %T and kind %s\n", i, i, v.Kind().String())
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		output = append(output, interface{}(v.Int()))
	case reflect.Slice:
		for j := 0; j < v.Len(); j++ {
			output = append(output, Flatten(interface{}(v.Index(j)))...)
		}
		//output = append(output, interface{}(v.Int()))
	case reflect.Struct:
		for j := 0; j < v.NumField(); j++ {
			fmt.Printf("Value v.Field(j) is %#v of type %T and kind %s\n", v.Field(j), v.Field(j), v.Field(j).String())
			switch v.Field(j).Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				output = append(output, v.Field(j).Int())
			case reflect.UnsafePointer:
				fmt.Printf("Unsafe pointer!!!\n")
				p := v.Field(j).Pointer()
				//valueInt64 := *(*int64)(p)
				valueInt64 := *(*int)(unsafe.Pointer(p))
				fmt.Printf("%d\n", valueInt64)
			}

		}

	}

	return output

}

// */

func Flatten(f interface{}) []interface{} {

	output := []interface{}{}

	switch e := f.(type) {
	case []interface{}:
		for _, v := range e {
			output = append(output, Flatten(v)...)
		}
	case interface{}:
		output = append(output, e)
	}
	return output
}

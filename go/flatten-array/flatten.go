package flatten

import "reflect"

func Flatten(i interface{}) []interface{} {

	output := make([]interface{}, 0)

	switch reflect.TypeOf(i).Kind() {
	case reflect.Slice:
		slice := reflect.ValueOf(i)
		for j := 0; j < slice.Len(); j++ {
			val := slice.Index(j).Interface()
			if val == nil {
				continue
			}
			switch reflect.TypeOf(val).Kind() {
			case reflect.Int:
				output = append(output, val)
			case reflect.Slice:
				output = append(output, Flatten(val)...)
			}
		}
	}

	return output

}

// */

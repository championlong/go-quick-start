package utils

import "reflect"

func StructToMap(data interface{}, tag string) map[string]interface{} {
	result := make(map[string]interface{})
	values := reflect.ValueOf(data)
	types := values.Type()
	for i := 0; i < types.NumField(); i++ {
		column := types.Field(i).Tag.Get(tag)
		if column != "" && column != "-" {
			result[column] = values.Field(i).Interface()
			if len(column) > 4 && (column[len(column)-4:] == "time" || column[:3] == "gmt") {
				if values.Field(i).String() == "" {
					result[column] = nil
				} else if values.Field(i).String() == "-" {
					result[column] = ""
				}

			}
		}
	}
	return result
}

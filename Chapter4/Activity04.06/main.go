package main

import "fmt"

func getTypeAsString(v any) string {
	switch v.(type) {
	case int, int32, int64:
		return "int"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	}
	return "unknown"
}

func main() {
	slice := []any{
		int(1), float64(3.14), "hello", true, struct{}{},
	}
	for _, v := range slice {
		fmt.Printf("%v is %s\n", v, getTypeAsString(v))
	}
}

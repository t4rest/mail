package main

func showMeTheType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case uint:
		return "uint"
	case int8:
		return "int8"
	case float64:
		return "float64"
	case string:
		return "string"
	case rune:
		return "int32"
	case []int:
		return "[]int"
	case map[string]bool:
		return "map[string]bool"
	default:
		return "undefined"
	}

}

func main() {

}

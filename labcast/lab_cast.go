package labcast

import (
	"fmt"
	"strconv"
	"strings"
)

// CastStringToInt
func CastIntToString(value int) string {
	return strconv.Itoa(value)
}

// Convert string to int
func CastStringToInt(value string) int {
	res, _ := strconv.Atoi(value)

	return res
}

// Convert string to int64
func CastStringToInt64(value string) int64 {
	return int64(CastStringToInt(value))
}

// Convert int64 to string
func CastInt64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}

// Convert an array of byte in an hexadecimal value
func ConvertBytes2Hexa(n []byte) string {
	return fmt.Sprintf("%X", n)
}

// Convert a string to an array of string based on the given separator
func ConvertArrayStringToListOfParams(list []string) string {
	return "'" + strings.Join(list, "','") + "'"
}

// Convert a string to an array of string based on the given separator
func ConvertStringToArray(str, sep string) []string {
	return strings.Split(str, sep)
}

// Convert a coma separated list of element to an array and clean the last element if it's empty
func ConvertListToArray(list string) []string {
	array := strings.Split(list, ",")

	if len(array) > 0 && array[len(array)-1] == "" {
		array = array[:len(array)-1]
	}

	return array
}

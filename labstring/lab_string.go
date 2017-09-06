package labstring

import (
	"strconv"
	"strings"
)

func ToLowerCase(str string) string {
	return strings.ToLower(str)
}

func ToUpperCase(str string) string {
	return strings.ToUpper(str)
}

func StringBuild(s string, args ...string) (sRes string) {
	sRes = s

	// Reverse Range
	for iCnt := len(args) - 1; iCnt >= 0; iCnt-- {
		sRes = Replace(sRes, "%"+strconv.Itoa(iCnt+1), args[iCnt])
	}

	return
}

func EscapeQuote(s string) string {
	return strings.Replace(s, "'", "''", -1)
}

func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

func Replace(s, sOld, sNew string) string {
	// -1 is for all instances of the string
	return strings.Replace(s, sOld, sNew, -1)
}

// Get the value but, if empty, return the default value
func GetValue(defaultValue, value string) string {
	if value != "" {
		return value
	}

	return defaultValue
}

package task1

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

func getThreeInts() []int {
	ints := make([]int, 0, 3)
	justInt := rand.Intn(1000)
	octal := rand.Intn(0o666)
	hex := rand.Intn(0xFFFF)
	ints = append(ints, justInt, octal, hex)
	return ints
}

func strconvForThreeInts(ints []int) []string {
	result := make([]string, 0, 3)
	common := fmt.Sprintf("%d", ints[0])
	octal := fmt.Sprintf("0%o", ints[1])
	hex := fmt.Sprintf("0x%x", ints[2])
	result = append(result, common, octal, hex)
	return result
}

func getFloat() float64 {
	return rand.Float64() * 1000.0
}

func getBoolean() bool {
	return rand.Intn(2) == 0
}

func getComplex() complex64 {
	return complex(rand.Float32(), rand.Float32())
}

func getString() string {
	return "realstring"
}

func assertValue(val any) {
	switch val.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case bool:
		fmt.Println("bool")
	case complex64:
		fmt.Println("complex64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown type")
	}
}

func makeSlice(value ...any) (result []any) {
	for _, v := range value {
		result = append(result, v)
	}
	return
}

func Task1() {
	ints := getThreeInts()
	floatVal := getFloat()
	boolVal := getBoolean()
	complexVal := getComplex()
	stringVal := getString()

	convertedInts := strconvForThreeInts(ints)

	var stringWithValues strings.Builder

	values := makeSlice(ints[0], ints[1], ints[2], floatVal, boolVal, complexVal, stringVal)
	for i, v := range values {
		assertValue(v)
		if i < 3 {
			stringWithValues.WriteString(convertedInts[i])
		} else {
			stringWithValues.WriteString(fmt.Sprintf("%v", v))
		}
	}

	valuesBeforeHashing := []rune(stringWithValues.String())

	salt := []rune("go-2024")

	valuesWithSaltBeforeHashing := slices.Concat(valuesBeforeHashing[:len(valuesBeforeHashing)/2], salt, valuesBeforeHashing[len(valuesBeforeHashing)/2+1:])

	resultSHA256 := sha256.Sum256([]byte(string(valuesWithSaltBeforeHashing)))
	fmt.Printf("SHA256 hash: %x\n", resultSHA256)
}

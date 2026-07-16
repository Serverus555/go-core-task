package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

func main() {
	hash := do(42, 052, 0x2A, 3.14, "Golang", true, complex64(1+2i))
	println(fmt.Sprintf("%x", hash))
}

func do(a, b, c int, fp float64, str string, boolean bool, comp complex64) [32]byte {
	var allStr string
	all := [7]any{a, b, c, fp, str, boolean, comp}
	for _, val := range all {
		println(reflect.TypeOf(val).Name())
		allStr += fmt.Sprintf("%v", val)
	}
	runes := []rune(allStr)
	half := len(runes) / 2
	salt := []rune("go-2024")
	salted := make([]rune, 0, len(runes)+len(salt))
	salted = append(salted, runes[:half]...)
	salted = append(salted, salt...)
	salted = append(salted, runes[half:]...)
	return sha256.Sum256([]byte(string(salted)))
}

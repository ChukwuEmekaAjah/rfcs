package main

import (
	"fmt"

	"github.com/ChukwuEmekaAjah/rfcs/8259/parser"
)

const json string = "[23,329,19,90,78]"
const json2 string = "false true"
const json3 string = "[null, null, 23,\"hello world\"]"
const json4 string = "{\"b\":{}, \"vals\": [null, null, 23,\"hello world\"], \"hello\": \"dave chapelle\", \"name\": \"chuks\"}"
const json5 string = ""

func main() {
	scanner := parser.NewScanner(json4)
	tokens := scanner.ScanTokens()
	p := parser.NewParser(tokens)
	v := p.Decode()
	for _, token := range tokens {
		fmt.Printf("token is %v\n", token)
	}

	fmt.Printf("value is %v ", v)

}

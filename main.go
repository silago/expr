package main

import (
	. "expr/lib"
	"log"
	"os"
)

func main() {

	expression := os.Args[1]

	var y, _ = ExtractVariables(expression)
	log.Println(y)

	//var x, _ = ParseExpresstion(expression, nil)
	//log.Println(x)

}

package main

import (
	. "expr/lib"
	"log"
	"os"
)

func main() {
	expression := os.Args[1]
	var x,_=ParseExpresstion(expression, nil)
	log.Println(x)
}

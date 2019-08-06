package main

import (
	. "expr/lib"
	"log"

)

func main() {
	var x,_=ParseExpresstion("2^2", nil)
	log.Println(x)
}

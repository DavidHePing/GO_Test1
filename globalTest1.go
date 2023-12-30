package main

import "fmt"

var global = "global"

func globalTest1() {
	global := "local"
	fmt.Println("local first than global:", global)
}

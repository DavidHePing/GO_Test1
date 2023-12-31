package main

import (
	"Go_Test1/printTest"
	"Go_Test1/threadSafe"
)

func main() {
	//print
	printTest.TestPrint()
	// printTest.PrintTestLowercase()

	//slice(list)
	// sliceTest1()
	// sliceTest2()

	//array
	// ArrayTestPrint()

	//for
	// ForTest1()
	// ForTestArray()

	//if
	// ifTest1()
	// ifTest1()

	//switch
	// switchTest1()

	//func
	// funcTest1()
	// funcTest2()

	//defer
	// deferTest1()

	//delegate
	// delegateTest1()
	// lambdaTest1()

	//panic
	// panicTest1()

	//error
	// errorTest1()

	//string
	// stringTest1()

	//stringBuilder
	// stringBuilderTest1()

	//strconv
	// strconvTest1()

	//global
	// globalTest1()

	//map
	// mapTest1()

	//threadSafe
	// threadSafe.SyncMapTest1()
	threadSafe.ConncurrentMapTest1()
}

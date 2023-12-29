package main

import (
	"errors"
	"fmt"
)

func errorTest1() {
	errorTest1_1("123")
	errorTest1_1("456")
	errorTest1_1("")
}

func errorTest1_1(st string) {
	result, err := errorTest1_1_1(st)

	if err != nil {
		fmt.Println("Error!, ", err)
	} else {
		fmt.Println("Success! result: ", result)
	}
}

func errorTest1_1_1(st string) (result string, err error) {
	if st == "" {
		err = errors.New("st is empty !!!")
		result = fmt.Sprint("result: ", st, ", state: failed")
		return result, err
	} else {
		result = fmt.Sprint("result: ", st, ", state: Success")
		return result, nil
	}
}

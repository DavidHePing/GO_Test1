package threadSafe

import (
	"fmt"
	"sync"
)

func MapNotThreadSafe() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 1; i <= 100; i++ {
			m["a"] = 100
		}
	}()

	go func() {
		defer wg.Done()
		defer fmt.Println("")

		for i := 1; i <= 100; i++ {
			fmt.Print(m["a"])
		}
	}()

	wg.Wait()
	fmt.Println(m)
}

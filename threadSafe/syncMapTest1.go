package threadSafe

import (
	"fmt"
	"sync"
)

func SyncMapTest1() {
	m := sync.Map{}
	m.Store("a", 1)
	m.Store(1, "b")

	fmt.Println(m.Load("a"))
	fmt.Println(m.Load(1))
	fmt.Println(m.Load("1"))

	m.Delete("a")
	fmt.Println(m.Load("a"))

}

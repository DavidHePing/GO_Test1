package threadSafe

import (
	"fmt"

	cmap "github.com/orcaman/concurrent-map/v2"
)

func ConncurrentMapTest1() {
	m := cmap.New[int]()
	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 3)

	fmt.Println(m.Get("a"))
	fmt.Println(m.Get("z"))
}

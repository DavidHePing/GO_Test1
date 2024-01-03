package threadSafe

import (
	"strconv"
	"sync"

	cmap "github.com/orcaman/concurrent-map/v2"
)

func SyncMapReadWriteDiffKey(times int) {
	var wg sync.WaitGroup
	wg.Add(times * 2)
	m := sync.Map{}

	for i := 0; i < times; i++ {
		num := i
		go func() {
			defer wg.Done()
			m.Store(strconv.Itoa(num), num)
		}()
	}

	for i := 0; i < times; i++ {
		num := i
		go func() {
			defer wg.Done()
			m.Load(strconv.Itoa(num))
		}()
	}

	wg.Wait()
}

func ConncurrentMapReadWriteDiffKey(times int) {
	var wg sync.WaitGroup
	wg.Add(times * 2)
	m := cmap.New[int]()

	for i := 0; i < times; i++ {
		num := i
		go func() {
			defer wg.Done()
			m.Set(strconv.Itoa(num), num)
		}()
	}

	for i := 0; i < times; i++ {
		num := i
		go func() {
			defer wg.Done()
			m.Get(strconv.Itoa(num))
		}()
	}

	wg.Wait()
}

func SyncMapReadWriteSameKey(times int) {
	var wg sync.WaitGroup
	wg.Add(times * 2)
	m := sync.Map{}

	for i := 0; i < times; i++ {
		num := i
		go func() {
			defer wg.Done()
			m.Store(strconv.Itoa(num%100), num)
		}()
	}

	for i := 0; i < times; i++ {
		num := i
		go func() {
			defer wg.Done()
			m.Load(strconv.Itoa(num % 100))
		}()
	}

	wg.Wait()
}

func ConncurrentMapReadWriteSameKey(times int) {
	var wg sync.WaitGroup
	wg.Add(times * 2)
	m := cmap.New[int]()

	for i := 0; i < times; i++ {
		num := i
		go func() {
			defer wg.Done()
			m.Set(strconv.Itoa(num%100), num)
		}()
	}

	for i := 0; i < times; i++ {
		num := i
		go func() {
			defer wg.Done()
			m.Get(strconv.Itoa(num % 100))
		}()
	}

	wg.Wait()
}

func SyncMapRead(times int) {
	times *= 2
	var wg sync.WaitGroup
	wg.Add(times)
	m := sync.Map{}
	m.Store("1", 1)

	for i := 0; i < times; i++ {
		num := i
		go func() {
			defer wg.Done()
			m.Load(strconv.Itoa(num % 2))
		}()
	}

	wg.Wait()
}

func ConncurrentMapRead(times int) {
	times *= 2
	var wg sync.WaitGroup
	wg.Add(times)
	m := cmap.New[int]()
	m.Set("1", 1)

	for i := 0; i < times; i++ {
		num := i
		go func() {
			defer wg.Done()
			m.Get(strconv.Itoa(num % 2))
		}()
	}

	wg.Wait()
}

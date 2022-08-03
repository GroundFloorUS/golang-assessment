package main

import (
	"fmt"
	"sync"
	"time"
)

type Value struct {
	Mu    *sync.Mutex
	Name  string
	Value int
}

func New(v int, name string) *Value {
	value := Value{
		Mu:    &sync.Mutex{},
		Name:  name,
		Value: v,
	}
	return &value
}

func IncrementAndSwapValues(wg *sync.WaitGroup, v1, v2 *Value) {
	defer wg.Done()
	v1.Mu.Lock()
	defer v1.Mu.Unlock()
	time.Sleep(time.Second)
	v2.Mu.Lock()
	defer v2.Mu.Unlock()
	value1 := v1.Value + 1
	value2 := v2.Value + 1
	v1.Value = value2
	v2.Value = value1
}

func main() {
	var wg sync.WaitGroup
	a := New(1, "a")
	b := New(2, "b")
	wg.Add(2)
	go IncrementAndSwapValues(&wg, a, b)
	go IncrementAndSwapValues(&wg, b, a)
	wg.Wait()
	fmt.Println("New Value for a:\t", a.Value)
	fmt.Println("New Value for b:\t", b.Value)
}

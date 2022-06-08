package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
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

func PrintSum(wg *sync.WaitGroup, v1, v2 *Value) {
	defer wg.Done()
	v1.Mu.Lock()
	defer v1.Mu.Unlock()
	_, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}
	v2.Mu.Lock()
	defer v2.Mu.Unlock()

	fmt.Printf("sum=%v\n", v1.Value+v2.Value)
}

func main() {
	var wg sync.WaitGroup
	a := New(0, "a")
	b := New(0, "b")
	wg.Add(2)
	go PrintSum(&wg, a, b)
	go PrintSum(&wg, b, a)
	wg.Wait()
}

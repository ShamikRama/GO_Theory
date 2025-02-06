package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Base struct {
	name string
}

type Child struct {
	lastName string
	Base
}

func (b *Base) Say() {
	fmt.Printf("Hello %s!\n", b.name)
}

func (c *Child) Say() {
	fmt.Printf("Hello %s %s!\n", c.lastName, c.name)
}

func newObject(objType string, name string, lastName string) interface{} {
	switch objType {
	case "Base":
		return &Base{name: name}
	case "Child":
		return &Child{Base: Base{name: name}, lastName: lastName}
	default:
		return nil
	}
}

func proccessObject(obj []interface{}) []interface{} {
	wg := sync.WaitGroup{}

	res := make([]interface{}, 0, len(obj))

	for _, val := range obj {
		wg.Add(1)
		go func() {
			defer wg.Done()
			switch v := val.(type) {
			case *Base:
				v.Say()
			case *Child:
				v.Say()
			}
		}()
	}

	wg.Wait()
	return res
}

func generateObj(ctx context.Context) []interface{} {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	res := make([]interface{}, 0)

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-time.After(time.Second * 1):
			mu.Lock()
			res = append(res, &Base{name: "Parent"})
			mu.Unlock()
		case <-ctx.Done():
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-time.After(time.Second * 2):
			mu.Lock()
			res = append(res, &Child{lastName: "Inherited", Base: Base{name: "Parent"}})
			mu.Unlock()
		case <-ctx.Done():
			return
		}
	}()

	wg.Wait()

	return res
}

func main() {
	b1 := &Base{
		name: "Parent",
	}

	c1 := &Child{
		Base:     Base{name: "Child"},
		lastName: "Inherited",
	}

	b1.Say()
	c1.Say()

	arr := []interface{}{b1, c1}

	for _, val := range arr {
		switch v := val.(type) {
		case *Base:
			v.Say()
		case *Child:
			v.Say()
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 11*time.Second)
	defer cancel()
	fmt.Println(newObject("Base", "Shamil", ""))
	fmt.Println(newObject("Child", "Shamil", "Ramazanov"))
	generateObj(ctx)
	proccessObject(arr)

}

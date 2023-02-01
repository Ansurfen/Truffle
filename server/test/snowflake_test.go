package test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"truffle/utils"
)

func TestSameNode(t *testing.T) {
	node1 := utils.NewSnowFlake(1)
	node2 := utils.NewSnowFlake(1)
	var dict sync.Map
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			id := node1.Generate()
			key := strconv.FormatInt(id.Int64(), 10)
			if _, ok := dict.Load(key); ok {
				dict.Store(key, 2)
			} else {
				dict.Store(key, 1)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			id := node2.Generate()
			key := strconv.FormatInt(id.Int64(), 10)
			if _, ok := dict.Load(key); ok {
				dict.Store(key, 2)
			} else {
				dict.Store(key, 1)
			}
		}
		wg.Done()
	}()
	wg.Wait()
	dict.Range(func(key, value any) bool {
		fmt.Print(value)
		return true
	})
}

func TestDifferentNode(t *testing.T) {
	node1 := utils.NewSnowFlake(1)
	node2 := utils.NewSnowFlake(2)
	var dict sync.Map
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			id := node1.Generate()
			key := strconv.FormatInt(id.Int64(), 10)
			if _, ok := dict.Load(key); ok {
				dict.Store(key, 2)
			} else {
				dict.Store(key, 1)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			id := node2.Generate()
			key := strconv.FormatInt(id.Int64(), 10)
			if _, ok := dict.Load(key); ok {
				dict.Store(key, 2)
			} else {
				dict.Store(key, 1)
			}
		}
		wg.Done()
	}()
	wg.Wait()
	dict.Range(func(key, value any) bool {
		fmt.Print(value)
		return true
	})
}

func TestDifferentNodeResult(t *testing.T) {
	node1 := utils.NewSnowFlake(1)
	node2 := utils.NewSnowFlake(2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			id := node1.Generate()
			fmt.Printf("node1 %v\n", id)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			id := node2.Generate()
			fmt.Printf("node2 %v\n", id)
		}
		wg.Done()
	}()
	wg.Wait()
}

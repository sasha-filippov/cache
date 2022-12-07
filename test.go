package main

import (
	"fmt"
	"github.com/sasha-filippov/cache/cache"
)

func main() {
	cache := cache.New()
	fmt.Println(cache)
	fmt.Println(cache.Size())
	cache.Set("test", "something")
	cache.Set("test2", "something2")
	cache.Set("test3", "something3")
	fmt.Println(cache.Get("test"))
	cache.Set("test", "nothing")
	fmt.Println(cache.Get("test"))
	fmt.Println(cache.Size())
	cache.Delete("test")
	fmt.Println(cache.Size())
	cache.Delete("test2")
	cache.Delete("test3")
	fmt.Println(cache.Size())
	fmt.Println(cache.Get("test"))
}

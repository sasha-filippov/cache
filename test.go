package main

import (
	"fmt"
	"github.com/sasha-filippov/cache/cache"
)

func main() {
	c := cache.New()
	fmt.Println(c)
	fmt.Println(c.Size())
	c.Set("test", "something")
	c.Set("test2", "something2")
	c.Set("test3", "something3")
	fmt.Println(c.Get("test"))
	c.Set("test", "nothing")
	fmt.Println(c.Get("test"))
	fmt.Println(c.Size())
	c.Delete("test")
	fmt.Println(c.Size())
	c.Delete("test2")
	c.Delete("test3")
	fmt.Println(c.Size())
	fmt.Println(c.Get("test"))
}

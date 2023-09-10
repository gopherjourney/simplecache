package main

import (
	"fmt"
	"time"

	"github.com/gopherjourney/simplecache"
)

func logThing() {
	fmt.Println("Defer")
}

func test() {
	defer logThing()

	fmt.Println("Starting Func")
}

func main() {
	c := simplecache.New()

	c.SetWithTTL("test", 10, 2*time.Second)

	time.Sleep(5 * time.Second)

	v, ok := c.Get("test")
	if ok {
		fmt.Println("VALUE:", v)
	} else {
		fmt.Printf("%s not available\n", "test")
	}
}

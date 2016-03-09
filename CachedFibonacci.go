package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"strconv"
	"time"
)

func fibonacci(s int, c *cache.Cache) int {
	cached, found := c.Get(strconv.Itoa(s))
	if found {
		return cached.(int)
	} else if s < 2 {
		c.Set(strconv.Itoa(s), 1, cache.DefaultExpiration)
		return 1
	} else {
		result := fibonacci(s-1, c) + fibonacci(s-2, c)
		c.Set(strconv.Itoa(s), result, cache.DefaultExpiration)
		return result
	}
}

func main() {
	c := cache.New(5*time.Minute, 30*time.Second)

	var input int
	fmt.Print("Number : ")
	fmt.Scanln(&input)
	fmt.Println(fibonacci(input, c))
}

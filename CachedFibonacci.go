package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"strconv"
	"time"
)

func fibonacci(s uint64, c *cache.Cache) uint64 {
	cached, found := c.Get(strconv.FormatUint(s, 32))
	if found {
		return cached.(uint64)
	} else if s < 2 {
		c.Set(strconv.FormatUint(s, 32), uint64(1), cache.DefaultExpiration)
		return 1
	} else {
		result := fibonacci(s-1, c) + fibonacci(s-2, c)
		c.Set(strconv.FormatUint(s, 32), result, cache.DefaultExpiration)
		return result
	}
}

func main() {
	c := cache.New(5*time.Minute, 30*time.Second)

	var input uint64
	fmt.Print("Number : ")
	fmt.Scanln(&input)
	fmt.Println(fibonacci(input, c))
}

package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"os"
	"strconv"
	"time"
)

func fibonacci(s uint64, c *cache.Cache) uint64 {
	cached, found := c.Get(strconv.FormatUint(s, 32))
	if found {
		return cached.(uint64)
	} else if s == 0 {
		c.Set(strconv.FormatUint(s, 32), uint64(0), cache.DefaultExpiration)
		return 0
	} else if s < 3 {
		c.Set(strconv.FormatUint(s, 32), uint64(1), cache.DefaultExpiration)
		return 1
	} else {
		result := fibonacci(s-1, c) + fibonacci(s-2, c)
		c.Set(strconv.FormatUint(s, 32), result, cache.DefaultExpiration)
		return result
	}
}

func factorial(s uint64, c *cache.Cache) uint64 {
	cached, found := c.Get(strconv.FormatUint(s, 32))
	if found {
		return cached.(uint64)
	} else if s < 1 {
		c.Set(strconv.FormatUint(s, 32), uint64(1), cache.DefaultExpiration)
		return 1
	} else {
		result := s * factorial(s-1, c)
		c.Set(strconv.FormatUint(s, 32), result, cache.DefaultExpiration)
		return result
	}
}

func intInSlice(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func main() {
	c := cache.New(5*time.Minute, 30*time.Second)

	var funcType int
	var input uint64
	choiceSet := []int{0, 1}

	fmt.Println("Please select your function call from the following")
	fmt.Println("(0) Fibonacci Numbers")
	fmt.Println("(1) Factorial of Number\n")
	for true {
		fmt.Print("Your Choice: ")
		fmt.Scanln(&funcType)
		if intInSlice(funcType, choiceSet) {
			break
		}
	}

	for true {
		fmt.Print("Number : ")
		fmt.Scanln(&input)
		if input == 0 {
			os.Exit(1)
		} else if funcType == 0 {
			fmt.Println(fibonacci(input, c))
		} else if funcType == 1 {
			fmt.Println(factorial(input, c))
		}
	}
}

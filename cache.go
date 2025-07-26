package main

import (
	"fmt"
	"time"
)

type serviceCache struct {
	cacher *cacheNow
}

func (sc *serviceCache) heavyFunction(key string) string {
	time.Sleep(5 * time.Second) // Simulate a heavy computation
	return fmt.Sprintf("Result data %s %s", key, "successfully cached")
}


type cacheNow struct {
	storage map[string]string
}


func (c *cacheNow) set(key, value string){
	c.storage[key] = value
}

func (c *cacheNow) get(key string) string {
	v, ok := c.storage[key]
	if !ok {
		return ""
	}
	return v
}

func (sc *serviceCache) getCachedData(key string) string {
	if sc.cacher != nil {
		if cacheData := sc.cacher.get(key); cacheData != "" {
			fmt.Println("Data found in cache:", cacheData)
			return cacheData
		}
	}
	result := sc.heavyFunction(key)
	if sc.cacher != nil {
		sc.cacher.set(key, result)
	}
	return result
}

func main() {
	cacher := &cacheNow{
		storage: map[string]string{},
	}
	service := &serviceCache{
		cacher: cacher,
	}

	data := "Data 1"

	start := time.Now() 

	fmt.Println("First call to heavyFunction:")

	result := service.getCachedData(data)

	fmt.Println("How long did it take to execute the first call:", time.Since(start))
	fmt.Println("Result:", result)


	start = time.Now() 

	fmt.Println("First call to heavyFunction:")

	result = service.getCachedData(data)

	fmt.Println("How long did it take to execute the first call after cached :", time.Since(start))
	fmt.Println("Result:", result)
}

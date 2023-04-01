package proxy_mod

import (
	"fmt"
	"testing"
)

func TestCache_Read(t *testing.T) {
	cache := NewCache()
	for i := 1; i <= 3; i++ {
		fmt.Println("-----------------")
		fmt.Printf("第 %d 次读:\n", i)
		key := "foo"
		fmt.Printf("key: %s, value: %s\n", key, cache.Read(key))
		fmt.Println("-----------------")
	}
}

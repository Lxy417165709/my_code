package proxy_mod

import "fmt"

/*
问 ChatGPT: 以下代码符合代理模式吗？

是的，以上代码可以认为是代理模式的一种应用。
在这个例子中，Cache 充当代理角色，DB 充当被代理角色。Cache 中维护了一个内部状态 data，用于存储从 DB 中读取的数据。
在 Cache 的 Read 方法中，如果请求的数据已经被缓存了，则直接返回缓存中的数据；
否则，Cache 会委托 DB 对象去读取数据，并在读取到数据之后将其缓存。
在这个过程中，Cache 控制了对 DB 对象的访问，并在必要时增强了其功能，例如添加缓存等特性。
因此，可以将以上代码视为代理模式的一种实现。

*/

// Cache 缓存。(充当代理角色)
type Cache struct {
	db   *DB
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{
		db:   &DB{}, // 这里可以用到创建型模式。
		data: make(map[string]string),
	}
}

// Read 读取数据。
func (c *Cache) Read(key string) string {
	// 1. 从缓存读取。
	v, ok := c.data[key]
	if ok {
		fmt.Printf("从缓存读取成功，key=[%+v]。\n", key)
		return v
	}

	// 2. 从数据库读取、写入缓存。
	c.data[key] = c.db.Read(key)
	fmt.Printf("从数据库读取成功，key=[%+v]。\n", key)

	// 3. 返回。
	return c.data[key]
}

// DB 数据库。(充当被代理角色)
type DB struct{}

// Read 读取数据。
func (c *DB) Read(key string) string {
	// ...
	// 经过一系列查询后，返回。
	return "bar"
}

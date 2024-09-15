#### ETCD 工具类

##### 初始化一个ETCD 客户端

**参数**：`addr` etcd服务器地址 

**返回**：`*clientv3.Client` etcd 客户端

```go
func initEtcd(add string) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{add},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return cli
}

```

##### 上送服务地址

**参数**：`etcdAddr` etcd服务器地址，`serviceName` 服务名称, `addr` 服务地址

**返回**：`error`

```go
// DeliveryAddress 上送服务地址
func DeliveryAddress(etcdAddr string, serviceName string, addr string) error {
	split := strings.Split(addr, ":")
	if len(split) != 2 {
		return errors.New("上送地址错误")
	}
	if split[0] == "0.0.0.0" {
		ip := netx.InternalIp()
		addr = strings.ReplaceAll(addr, "0.0.0.0", ip)
	}
	client := initEtcd(etcdAddr)
	_, err := client.Put(context.Background(), serviceName, addr)
	if err != nil {
		return errors.New("上送地址失败")
	}
	logx.Infof("上送地址成功 %s  %s", serviceName, addr)
	return nil
}
```

##### 获取服务地址

**参数**：`etcdAddr` etcd服务器地址，`serviceName` 服务名称

**返回**：服务地址

```go
func GetServiceAddr(etcdAddr string, serviceName string) string {
	client := initEtcd(etcdAddr)
	res, err := client.Get(context.Background(), serviceName)
	if err == nil && len(res.Kvs) > 0 {
		return string(res.Kvs[0].Value) // 返回服务地址
	}
	return ""
}

```

##### 完整代码

**etcd_util.go**

```go
package etcd

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/netx"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

func initEtcd(add string) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{add},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return cli
}

// DeliveryAddress 上送服务地址
func DeliveryAddress(etcdAddr string, serviceName string, addr string) error {
	split := strings.Split(addr, ":")
	if len(split) != 2 {
		return errors.New("上送地址错误")
	}
	if split[0] == "0.0.0.0" {
		ip := netx.InternalIp()
		addr = strings.ReplaceAll(addr, "0.0.0.0", ip)
	}
	client := initEtcd(etcdAddr)
	_, err := client.Put(context.Background(), serviceName, addr)
	if err != nil {
		return errors.New("上送地址失败")
	}
	logx.Infof("上送地址成功 %s  %s", serviceName, addr)
	return nil
}

func GetServiceAddr(etcdAddr string, serviceName string) string {
	client := initEtcd(etcdAddr)
	res, err := client.Get(context.Background(), serviceName)
	if err == nil && len(res.Kvs) > 0 {
		return string(res.Kvs[0].Value) // 返回服务地址
	}
	return ""
}

```

#### CacheBox 缓存工具类

##### 加密URI

```go
func MD5URI(uri string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(uri)))
}
```

##### 从 Redis 缓存中获取数据 

**参数：** `uri` redis中的 key

**返回：** `*Cache` 缓存类 [🛸](/各种类型/?id=cache-类)

```go
func (c *CacheBox) Get(uri string) *Cache {
	client := c.pool.Get()
	defer c.pool.Put(client)
	cacheStr := client.Get(context.Background(), MD5URI(uri)).Val() // 从redis中获取缓存
	// 判断缓存是否命中

	cacheByte := []byte(cacheStr)
	cache := new(Cache)
	json.Unmarshal(cacheByte, &cache)
	return cache
}
```

##### 从 Redis 中删除缓存

**参数：** `uri` redis中的 key

```go
func (c *CacheBox) Delete(uri string) {
	client := c.pool.Get()
	defer c.pool.Put(client)
	n := client.Del(context.Background(), MD5URI(uri))
	if n.Val() == 1 {
		logger.Info("删除缓存URI: %s 成功", uri)
	} else {
		logger.Info("删除缓存URI: %s 失败", uri)
	}
}
```

##### 判断响应类型是否能缓存

> 有些响应是不让缓存的，有些响应是能缓存的
>
> 要判断响应是否能被缓存，需要看响应中的字段是否符合条件

**参数：** `*http.Response` http 响应类型

**返回：** `bool` 是否能缓存该响应

```go
func IsCache(resp *http.Response) bool {
	cacheControl := resp.Header.Get("Cache-Control")
	contentType := resp.Header.Get("Content-Type")
	if strings.Index(cacheControl, "private") != -1 || // 私有，不缓存
		strings.Index(cacheControl, "no-store") != -1 || // 私有，不缓存
		strings.Index(contentType, "application") != -1 ||
		strings.Index(contentType, "video") != -1 || 
		strings.Index(contentType, "audio") != -1 || // 视频，音频,应用，不缓存
		(strings.Index(cacheControl, "max-age") == -1 && //	
			strings.Index(cacheControl, "s-maxage") == -1 &&
			resp.Header.Get("Etag") == "" &&	
			resp.Header.Get("Last-Modified") == "" &&
			(resp.Header.Get("Expires") == "" || resp.Header.Get("Expires") == "0")) {
		return false
	}
	return true
}
```

##### 判断响应类型并缓存

```go
func (c *CacheBox) CheckAndStore(uri string, resp *http.Response, host string) {
	// 不满足缓存条件，直接返回
	// 这里的 uri 是指完整的url , 比如 http://127.0.0.1:9700/api/auth/authentication
	if !IsCache(resp) {
		logger.Infof("URI: %s 响应不满足缓存条件", uri)
		return
	}
	ctx := context.Background()
	cache := NewCacheForResp(resp, host, uri)
	cacheByte, _ := json.Marshal(cache)
	logger.Infof("缓存URI: %s 响应数据", uri)
	client := c.pool.Get()
	defer c.pool.Put(client)
	// 开启事务
	pipe := client.TxPipeline()
	pipe.Set(ctx, MD5URI(uri), cacheByte, time.Duration(cache.maxAge)*time.Second)
	_, err := pipe.Exec(ctx)
	if err != nil {
		logger.Error("缓存URI: %s 失败", uri)
		return
	}
	logger.Info("缓存URI: %s 成功", uri)
}
```

##### 完整代码

```go
package cache

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"server/common/logger"
	"strings"
	"time"
)

func MD5URI(uri string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(uri)))
}

type CacheBox struct {
	pool *RedisPool
}

func NewCacheBox(address string, password string) *CacheBox {
	return &CacheBox{
		pool: GlobalRedisPool,
	}
}

func (c *CacheBox) Get(uri string) *Cache {
	client := c.pool.Get()
	defer c.pool.Put(client)
	cacheStr := client.Get(context.Background(), MD5URI(uri)).Val() // 从redis中获取缓存
	// 判断缓存是否命中
	if cacheStr == "" {
		logger.Info("缓存URI: %s 未命中", uri)
		return nil
	}
	cacheByte := []byte(cacheStr)
	cache := new(Cache)
	json.Unmarshal(cacheByte, &cache)
	return cache
}

func (c *CacheBox) Delete(uri string) {
	client := c.pool.Get()
	defer c.pool.Put(client)
	n := client.Del(context.Background(), MD5URI(uri))
	if n.Val() == 1 {
		logger.Info("删除缓存URI: %s 成功", uri)
	} else {
		logger.Info("删除缓存URI: %s 失败", uri)
	}
}

func (c *CacheBox) CheckAndStore(uri string, resp *http.Response, host string) {
	// 不满足缓存条件，直接返回
	// 这里的 uri 是指完整的url , 比如 http://127.0.0.1:9700/api/auth/authentication
	if !IsCache(resp) {
		logger.Infof("URI: %s 响应不满足缓存条件", uri)
		return
	}
	ctx := context.Background()
	cache := NewCacheForResp(resp, host, uri)
	cacheByte, _ := json.Marshal(cache)
	logger.Infof("缓存URI: %s 响应数据", uri)
	client := c.pool.Get()
	defer c.pool.Put(client)
	// 开启事务
	pipe := client.TxPipeline()
	pipe.Set(ctx, MD5URI(uri), cacheByte, time.Duration(cache.maxAge)*time.Second)
	_, err := pipe.Exec(ctx)
	if err != nil {
		logger.Error("缓存URI: %s 失败", uri)
		return
	}
	logger.Info("缓存URI: %s 成功", uri)
}

func IsCache(resp *http.Response) bool {
	cacheControl := resp.Header.Get("Cache-Control")
	contentType := resp.Header.Get("Content-Type")
	if strings.Index(cacheControl, "private") != -1 || // 私有，不缓存
		strings.Index(cacheControl, "no-store") != -1 || // 私有，不缓存
		strings.Index(contentType, "application") != -1 ||
		strings.Index(contentType, "video") != -1 || // 视频，音频,应用，不缓存
		strings.Index(contentType, "audio") != -1 ||
		(strings.Index(cacheControl, "max-age") == -1 && strings.Index(cacheControl, "s-maxage") == -1 &&
			resp.Header.Get("Etag") == "" &&
			resp.Header.Get("Last-Modified") == "" &&
			(resp.Header.Get("Expires") == "" || resp.Header.Get("Expires") == "0")) {
		return false
	}
	return true
}

func CanCache(w http.ResponseWriter, lastModified string, etag string) {
	cacheControl := "public, max-age=86400"
	w.Header().Set("Cache-Control", cacheControl)
	w.Header().Set("Expires", time.Now().Add(24*time.Hour).Format(http.TimeFormat))
	w.Header().Set("Last-Modified", lastModified)
	w.Header().Set("ETag", etag)
}

```


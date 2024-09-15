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

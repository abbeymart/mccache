// @Author: abbeymart | Abi Akindele | @Created: 2020-03-09 | @Updated: 2020-03-09
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect cache - map[string]interface{}

package mccache

import (
	"time"
)

// Initialise cache object/dictionary (map)
var mcCache = SimpleCache{
	items:    make(map[string]CacheValue),
	capacity: 10_000,
}

// SecretCode is secret code for added security | default value
const SecretCode = "mcconnect_20200320"

func SetCache(key string, value ValueType, expire int64) CacheResponse {
	// validate required params
	if key == "" || value == nil {
		return CacheResponse{
			Ok:      false,
			Message: "cache key and ItemValue are required",
		}
	}
	// expire default ItemValue (in seconds)
	if expire == 0 {
		expire = 300
	}
	cacheKey := key + SecretCode

	// set cache ItemValue
	mcCache.mu.Lock()
	defer mcCache.mu.Unlock()
	mcCache.items[cacheKey] = CacheValue{
		value:     value,
		expire:    time.Now().Unix() + expire,
		createdAt: time.Now().Unix(),
	}
	// read cache-value
	var setCacheValue ValueType = nil
	if cValue, ok := mcCache.items[cacheKey]; ok {
		setCacheValue = cValue.value
	}
	// return successful response
	if setCacheValue != nil {
		return CacheResponse{
			Ok:      true,
			Message: "task completed successfully",
			Value:   setCacheValue,
		}
	}
	// check/track error
	return CacheResponse{
		Ok:      false,
		Message: "unable to set cache value",
		Value:   nil,
	}
}

func GetCache(key string) CacheResponse {
	// validate required params
	if key == "" {
		return CacheResponse{
			Ok:      false,
			Message: "cache-key is required",
		}
	}
	cacheKey := key + SecretCode
	mcCache.mu.Lock()
	defer mcCache.mu.Unlock()
	cValue, ok := mcCache.items[cacheKey]
	if (ok && cValue.value != nil) && cValue.expire > time.Now().Unix() {
		return CacheResponse{
			Ok:      true,
			Message: "task completed successfully",
			Value:   cValue.value,
		}
	} else if (ok && cValue.value != nil) && cValue.expire < time.Now().Unix() {
		// delete expired cache
		delete(mcCache.items, cacheKey)
		return CacheResponse{
			Ok:      false,
			Value:   nil,
			Message: "cache expired and deleted",
		}
	} else {
		return CacheResponse{
			Ok:      false,
			Value:   nil,
			Message: "cache info does not exist",
		}
	}
}

func DeleteCache(key string) CacheResponse {
	// validate required params
	if key == "" {
		return CacheResponse{
			Ok:      false,
			Message: "cache key is required",
		}
	}
	cacheKey := key + SecretCode
	mcCache.mu.Lock()
	defer mcCache.mu.Unlock()
	if _, ok := mcCache.items[cacheKey]; ok {
		delete(mcCache.items, cacheKey)
		return CacheResponse{
			Ok:      true,
			Message: "task completed successfully",
		}
	}
	return CacheResponse{
		Ok:      false,
		Message: "task not completed, cache-key-value not found",
	}
}

func ClearCache() CacheResponse {
	// clear mcCache map content
	mcCache.mu.Lock()
	defer mcCache.mu.Unlock()
	clear(mcCache.items)
	return CacheResponse{
		Ok:      true,
		Message: "task completed successfully",
	}
}

// Instance approach

// NewCache create a cache repository instance
func NewCache(capacity int, secretCode string) *SimpleCache {
	if secretCode == "" {
		secretCode = SecretCode
	}
	if capacity <= 0 {
		capacity = 10_000
	}
	return &SimpleCache{
		items:      make(map[string]CacheValue),
		capacity:   capacity,
		secretCode: secretCode,
	}
}

// removeDatedCache remove dated or expired cache item
func (c *SimpleCache) removeDatedCache() bool {
	// remove the oldest cache item
	cKey := ""
	expire := time.Now().Unix()
	createdAt := time.Now().Unix()
	// capture a cache item, as baseline/reference
	for key, val := range c.items {
		cKey = key
		expire = val.expire
		createdAt = val.createdAt
		break
	}
	// determine the oldest or expired cache item
	for k, v := range c.items {
		if v.createdAt < createdAt {
			cKey = k
			expire = v.expire
			createdAt = v.createdAt
		} else if v.expire < expire {
			cKey = k
			expire = v.expire
			createdAt = v.createdAt
			break
		}
	}
	// delete the oldest cache item
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, cKey)
	return true
}

// SetCache set the cache-value, by cache-key - including creation timestamp
func (c *SimpleCache) SetCache(key string, value ValueType, expire int64) CacheResponse {
	// validate required params
	if key == "" || value == nil {
		return CacheResponse{
			Ok:      false,
			Message: "valid cache key and value are required",
		}
	}
	// expire default ItemValue (in seconds)
	if expire == 0 {
		expire = 300
	}
	cacheKey := key + c.secretCode

	c.mu.Lock()
	defer c.mu.Unlock()
	// validate capacity
	if len(c.items) > c.capacity {
		// remove the most-dated or expired cache item
		c.removeDatedCache()
	}
	// set cache ItemValue
	c.items[cacheKey] = CacheValue{
		value:     value,
		expire:    time.Now().Unix() + expire,
		createdAt: time.Now().Unix(),
	}

	// read cache-value
	var setCacheValue ValueType = nil
	if cValue, ok := c.items[cacheKey]; ok {
		setCacheValue = cValue.value
	}

	// return successful response
	if setCacheValue != nil {
		return CacheResponse{
			Ok:      true,
			Message: "task completed successfully",
			Value:   setCacheValue,
		}
	}
	// check/track error
	return CacheResponse{
		Ok:      false,
		Message: "unable to set cache value",
		Value:   nil,
	}
}

// GetCache fetches non-expired cache-value, by cache-key
func (c *SimpleCache) GetCache(key string) CacheResponse {
	// validate required params
	if key == "" {
		return CacheResponse{
			Ok:      false,
			Message: "cache-key is required",
		}
	}
	cacheKey := key + c.secretCode
	c.mu.Lock()
	defer c.mu.Unlock()
	cValue, ok := c.items[cacheKey]
	if (ok && cValue.value != nil) && cValue.expire > time.Now().Unix() {
		return CacheResponse{
			Ok:      true,
			Message: "task completed successfully",
			Value:   cValue.value,
		}
	} else if (ok && cValue.value != nil) && cValue.expire < time.Now().Unix() {
		// delete expired cache
		delete(c.items, cacheKey)
		return CacheResponse{
			Ok:      false,
			Value:   nil,
			Message: "cache expired and deleted",
		}
	} else {
		return CacheResponse{
			Ok:      false,
			Value:   nil,
			Message: "cache info does not exist",
		}
	}
}

// DeleteCache deletes cache by cache-key
func (c *SimpleCache) DeleteCache(key string) CacheResponse {
	// validate required params
	if key == "" {
		return CacheResponse{
			Ok:      false,
			Message: "cache key is required",
		}
	}
	cacheKey := key + c.secretCode
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.items[cacheKey]; ok {
		delete(c.items, cacheKey)
		return CacheResponse{
			Ok:      true,
			Message: "task completed successfully",
		}
	}
	return CacheResponse{
		Ok:      false,
		Message: "task not completed, cache-key-value not found",
	}
}

func (c *SimpleCache) ClearCache() CacheResponse {
	// clear mcCache map content
	c.mu.Lock()
	defer c.mu.Unlock()
	clear(c.items)
	return CacheResponse{
		Ok:      true,
		Message: "task completed successfully",
	}
}

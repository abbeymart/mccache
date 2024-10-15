// @Author: abbeymart | Abi Akindele | @Created: 2020-03-09 | @Updated: 2020-03-09
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect cache hash option - key: map[string]interface{}

package mccache

import (
	"time"
)

// Initialise cache object/dictionary (map)
var mcHashCache = HashCache{
	items:    make(map[string]HashCacheValueType),
	capacity: 10_000,
}

func SetHashCache(key string, hash string, value ValueType, expire int64) CacheResponse {
	// validate required params
	if key == "" || hash == "" || value == nil {
		return CacheResponse{
			Ok:      false,
			Message: "hash, cache-key and value are required",
		}
	}
	// expire default Value (in seconds)
	if expire == 0 {
		expire = 300
	}
	cacheKey := key + SecretCode
	hashKey := hash + SecretCode

	// initialise a hashCacheValue
	hashCacheValue := HashCacheValueType{}

	hashCacheValue[cacheKey] = CacheValue{
		value:     value,
		expire:    time.Now().Unix() + expire,
		createdAt: time.Now().Unix(),
	}
	// set cache Value: mcHashCache.set(cacheKey, {Value: Value, expire: Date.now() + expire * 1000});
	var setCacheValue ValueType = nil

	mcHashCache.mu.Lock()
	defer mcHashCache.mu.Unlock()
	mcHashCache.items[hashKey] = hashCacheValue

	// read cache-value
	if _, ok := mcHashCache.items[hashKey]; ok {
		if cValue, cok := mcHashCache.items[hashKey][cacheKey]; cok {
			setCacheValue = cValue.value
		}
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

func GetHashCache(key string, hash string) CacheResponse {
	// validate required params
	if key == "" || hash == "" {
		return CacheResponse{
			Ok:      false,
			Message: "hash and cache-key are required",
		}
	}
	cacheKey := key + SecretCode
	hashKey := hash + SecretCode
	cValue, ok := mcHashCache.items[hashKey][cacheKey]
	if (ok && cValue.value != nil) && cValue.expire > time.Now().Unix() {
		return CacheResponse{
			Ok:      true,
			Message: "task completed successfully",
			Value:   cValue.value,
		}
	} else if (ok && cValue.value != nil) && cValue.expire < time.Now().Unix() {
		// delete expired cache
		mcHashCache.mu.Lock()
		defer mcHashCache.mu.Unlock()
		delete(mcHashCache.items[hashKey], cacheKey)
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

func DeleteHashCache(key string, hash string, by string) CacheResponse {
	// by default Value
	if by == "" {
		by = ByKey
	}
	// validate required params
	if key == "" || hash == "" && by == ByKey {
		return CacheResponse{
			Ok:      false,
			Message: "hash and cache keys are required",
		}
	}
	if hash == "" && by == ByHash {
		return CacheResponse{
			Ok:      false,
			Message: "hash key is required",
		}
	}
	cacheKey := key + SecretCode
	hashKey := hash + SecretCode

	if by == ByKey {
		// perform find and delete action
		if _, ok := mcHashCache.items[hashKey][cacheKey]; ok {
			mcHashCache.mu.Lock()
			defer mcHashCache.mu.Unlock()
			delete(mcHashCache.items[hashKey], cacheKey)
			return CacheResponse{
				Ok:      true,
				Message: "task completed successfully",
			}
		}
		return CacheResponse{
			Ok:      false,
			Message: "task not completed, hash-cache-key-value not found",
		}
	}
	if by == ByHash {
		// perform find and delete action
		if _, ok := mcHashCache.items[hashKey]; ok {
			mcHashCache.mu.Lock()
			defer mcHashCache.mu.Unlock()
			delete(mcHashCache.items, hashKey)
			return CacheResponse{
				Ok:      true,
				Message: "task completed successfully",
			}
		}
		return CacheResponse{
			Ok:      false,
			Message: "task not completed, hash-value not found",
		}
	}
	return CacheResponse{
		Ok:      false,
		Message: "task could not be completed due to incomplete inputs",
	}
}

func ClearHashCache() CacheResponse {
	// clear mcHashCache map content
	mcHashCache.mu.Lock()
	defer mcHashCache.mu.Unlock()
	clear(mcHashCache.items)
	return CacheResponse{
		Ok:      true,
		Message: "task completed successfully",
	}
}

// Instance approach

// NewHashCache creates a hash-cache repository instance
func NewHashCache(capacity int, secretCode string) *HashCache {
	if secretCode == "" {
		secretCode = SecretCode
	}
	if capacity <= 0 {
		capacity = 10_000
	}
	return &HashCache{
		items:    make(map[string]HashCacheValueType),
		capacity: capacity,
	}
}

// removeDatedCache remove dated or expired cache item
func (c *HashCache) removeDatedCache() bool {
	// remove the oldest cache item
	cKey := ""
	hashKey := ""
	expire := time.Now().Unix()
	createdAt := time.Now().Unix()
	// capture a hash-cache item, as baseline/reference
	for _, val := range c.items {
		for k, v := range val {
			cKey = k
			expire = v.expire
			createdAt = v.createdAt
		}
		break
	}
	// determine the oldest or expired hash-cache item
	for hKey, val := range c.items {
		hashKey = hKey
		isExpired := false
		for k, v := range val {
			if v.createdAt < createdAt {
				cKey = k
				expire = v.expire
				createdAt = v.createdAt
			} else if v.expire < expire {
				cKey = k
				expire = v.expire
				createdAt = v.createdAt
				isExpired = true
			}
			if isExpired {
				break
			}
		}
	}
	// delete the oldest hash-cache item
	delete(c.items[hashKey], cKey)
	return true
}

// SetCache set the cache-value, by hash and cache-key - including creation timestamp
func (c *HashCache) SetCache(key string, hash string, value ValueType, expire int64) CacheResponse {
	// validate required params
	if key == "" || hash == "" || value == nil {
		return CacheResponse{
			Ok:      false,
			Message: "hash, cache-key and value are required",
		}
	}
	// expire default Value (in seconds)
	if expire == 0 {
		expire = 300
	}
	cacheKey := key + c.secretCode
	hashKey := hash + c.secretCode

	// initialise a hashCacheValue
	hashCacheValue := HashCacheValueType{}

	hashCacheValue[cacheKey] = CacheValue{
		value:     value,
		expire:    time.Now().Unix() + expire,
		createdAt: time.Now().Unix(),
	}
	// set cache Value: mcHashCache.set(cacheKey, {Value: Value, expire: Date.now() + expire * 1000});
	var setCacheValue ValueType = nil

	c.mu.Lock()
	defer c.mu.Unlock()
	// validate capacity
	cacheLen := 0
	for _, val := range c.items {
		cacheLen += len(val)
	}
	if cacheLen > c.capacity {
		// remove the most-dated or expired cache item
		c.removeDatedCache()
	}
	c.items[hashKey] = hashCacheValue

	// read cache-value
	if _, ok := c.items[hashKey]; ok {
		if cValue, cok := c.items[hashKey][cacheKey]; cok {
			setCacheValue = cValue.value
		}
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
func (c *HashCache) GetCache(key string, hash string) CacheResponse {
	// validate required params
	if key == "" || hash == "" {
		return CacheResponse{
			Ok:      false,
			Message: "hash and cache-key are required",
		}
	}
	cacheKey := key + c.secretCode
	hashKey := hash + c.secretCode
	c.mu.Lock()
	defer c.mu.Unlock()
	cValue, ok := c.items[hashKey][cacheKey]
	if (ok && cValue.value != nil) && cValue.expire > time.Now().Unix() {
		return CacheResponse{
			Ok:      true,
			Message: "task completed successfully",
			Value:   cValue.value,
		}
	} else if (ok && cValue.value != nil) && cValue.expire < time.Now().Unix() {
		// delete expired cache
		delete(c.items[hashKey], cacheKey)
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

// DeleteCache deletes cache by hash or cache-key
func (c *HashCache) DeleteCache(key string, hash string, by string) CacheResponse {
	// by default Value
	if by == "" {
		by = ByKey
	}
	// validate required params
	if key == "" || hash == "" && by == ByKey {
		return CacheResponse{
			Ok:      false,
			Message: "hash and cache keys are required",
		}
	}
	if hash == "" && by == ByHash {
		return CacheResponse{
			Ok:      false,
			Message: "hash key is required",
		}
	}
	cacheKey := key + c.secretCode
	hashKey := hash + c.secretCode
	c.mu.Lock()
	defer c.mu.Unlock()
	if by == ByKey {
		// perform find and delete action
		if _, ok := c.items[hashKey][cacheKey]; ok {
			delete(c.items[hashKey], cacheKey)
			return CacheResponse{
				Ok:      true,
				Message: "task completed successfully",
			}
		}
		return CacheResponse{
			Ok:      false,
			Message: "task not completed, hash-cache-key-value not found",
		}
	}
	if by == ByHash {
		// perform find and delete action
		if _, ok := c.items[hashKey]; ok {
			delete(c.items, hashKey)
			return CacheResponse{
				Ok:      true,
				Message: "task completed successfully",
			}
		}
		return CacheResponse{
			Ok:      false,
			Message: "task not completed, hash-value not found",
		}
	}
	return CacheResponse{
		Ok:      false,
		Message: "task could not be completed due to incomplete inputs",
	}
}

// ClearCache clears hash-cache
func (c *HashCache) ClearCache() CacheResponse {
	// clear mcHashCache map content
	c.mu.Lock()
	defer c.mu.Unlock()
	clear(c.items)
	return CacheResponse{
		Ok:      true,
		Message: "task completed successfully",
	}
}

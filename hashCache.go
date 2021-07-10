// @Author: abbeymart | Abi Akindele | @Created: 2020-03-09 | @Updated: 2020-03-09
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect cache hash option - key: map[string]interface{}

package mccache

import "time"

// Initialise cache object/dictionary (map)
var mcHashCache = make(map[string]HashCacheValueType)

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
	cacheKey := key + keyCode
	hashKey := hash + keyCode

	// initialise a hashCacheValue
	hashCacheValue := HashCacheValueType{}

	hashCacheValue[cacheKey] = CacheValue{
		value:  value,
		expire: time.Now().Unix() + expire,
	}
	// set cache Value: mcHashCache.set(cacheKey, {Value: Value, expire: Date.now() + expire * 1000});
	mcHashCache[hashKey] = hashCacheValue
	// return successful response
	if _, ok := mcHashCache[hashKey]; ok {
		if cValue, cok := mcHashCache[hashKey][cacheKey]; cok {
			return CacheResponse{
				Ok:      true,
				Message: "task completed successfully",
				Value:   cValue.value,
			}
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
	cacheKey := key + keyCode
	hashKey := hash + keyCode
	cValue, ok := mcHashCache[hashKey][cacheKey]
	if (ok && cValue.value != nil) && cValue.expire > time.Now().Unix() {
		return CacheResponse{
			Ok:      true,
			Message: "task completed successfully",
			Value:   cValue.value,
		}
	} else if (ok && cValue.value != nil) && cValue.expire < time.Now().Unix() {
		// delete expired cache
		delete(mcHashCache[hashKey], cacheKey)
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
		by = "key"
	}
	// validate required params
	if key == "" || hash == "" && by == "key" {
		return CacheResponse{
			Ok:      false,
			Message: "hash and cache keys are required",
		}
	}
	if hash == "" && by == "hash" {
		return CacheResponse{
			Ok:      false,
			Message: "hash key is required",
		}
	}
	cacheKey := key + keyCode
	hashKey := hash + keyCode

	if by == "key" {
		// perform find and delete action
		if _, ok := mcHashCache[hashKey][cacheKey]; ok {
			delete(mcHashCache[hashKey], cacheKey)
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
	if by == "hash" {
		// perform find and delete action
		if _, ok := mcHashCache[hashKey]; ok {
			delete(mcHashCache, hashKey)
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
	for key := range mcHashCache {
		delete(mcHashCache, key)
	}
	// mcHashCache = map[string]HashCacheValueType{}
	return CacheResponse{
		Ok:      true,
		Message: "task completed successfully",
	}
}

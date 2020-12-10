// @Author: abbeymart | Abi Akindele | @Created: 2020-03-09 | @Updated: 2020-03-09
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect cache hash option - key: map[string]interface{}

package mccache

import "time"

// Initialise cache object/dictionary (map)
//var mcCache map[string]CacheValue
var mcHashCache = make(map[string]HashCacheValueType)

// secret keyCode for added security
//const keyCode = "mcconnect_20200320"

func SetHashCache(key string, hash string, value ValueType, expire uint) CacheResponse {
	// validate required params
	if key == "" || hash == "" || value == nil {
		return CacheResponse{
			Ok:      false,
			Message: "cache key, hash and Value are required",
		}
	}
	// expire default Value
	if expire == 0 {
		expire = 10
	}
	cacheKey := key + keyCode
	hashKey := hash + keyCode

	// initialise a hashCacheValue
	hashCacheValue := HashCacheValueType{}
	//var hashCacheValue HashCacheValueType

	hashCacheValue[hashKey] = CacheValue{
		value:  value,
		expire: uint(time.Now().Unix()) + expire,
	}
	// set cache Value: mcHashCache.set(cacheKey, {Value: Value, expire: Date.now() + expire * 1000});
	mcHashCache[cacheKey] = hashCacheValue
	// return successful response
	if _, ok := mcHashCache[cacheKey]; ok {
		if hValue, hok := mcHashCache[cacheKey][hashKey]; hok {
			return CacheResponse{
				Ok:      true,
				Message: "task completed successfully",
				Value:   hValue.value,
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
			Message: "key and hash-key are required",
		}
	}
	cacheKey := key + keyCode
	hashKey := hash + keyCode
	cValue, ok := mcHashCache[cacheKey][hashKey]
	if (ok && cValue.value != nil) && cValue.expire > uint(time.Now().Unix()) {
		return CacheResponse{
			Ok:      true,
			Message: "task completed successfully",
			Value:   cValue.value,
		}
	} else if (ok && cValue.value != nil) && cValue.expire < uint(time.Now().Unix()) {
		// delete expired cache
		delete(mcHashCache, cacheKey)
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
	// validate required params
	if key == "" || hash == "" && by == "hash" {
		return CacheResponse{
			Ok:      false,
			Message: "cache key is required",
		}
	}
	// by default Value
	if by == "" {
		by = "hash"
	}
	cacheKey := key + keyCode
	hashKey := hash + keyCode

	if key != "" && by == "key" {
		// perform find and delete action
		if _, ok := mcHashCache[cacheKey]; ok {
			delete(mcHashCache, cacheKey)
			return CacheResponse{
				Ok:      true,
				Message: "task completed successfully",
			}
		}
		return CacheResponse{
			Ok:      false,
			Message: "task not completed, cache-key-Value not found",
		}
	}
	if key != "" && hash != "" && by == "hash" {
		// perform find and delete action
		if _, ok := mcHashCache[cacheKey][hashKey]; ok {
			delete(mcHashCache[cacheKey], hashKey)
			return CacheResponse{
				Ok:      true,
				Message: "task completed successfully",
			}
		}
		return CacheResponse{
			Ok:      false,
			Message: "task not completed, cache-key-hash-Value not found",
		}
	}
	return CacheResponse{
		Ok:      false,
		Message: "task could not be completed due to incomplete inputs",
	}
}

func ClearHashCache() CacheResponse {
	// clear mcHashCache map content
	mcHashCache = map[string]HashCacheValueType{}
	//for key := range mcHashCache {
	//	delete(mcHashCache, key)
	//}
	return CacheResponse{
		Ok:      true,
		Message: "task completed successfully",
	}
}

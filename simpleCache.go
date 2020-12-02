// @Author: abbeymart | Abi Akindele | @Created: 2020-03-09 | @Updated: 2020-03-09
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect cache - map[string]interface{}

package mccache

import "time"

// Initialise cache object/dictionary (map)
//var mcCache map[string]CacheValue
var mcCache = make(map[string]CacheValue)

// secret keyCode for added security
const keyCode = "mcconnect_20200320"

func SetCache(key string, value ValueType, expire uint) CacheResponse {
	// validate required params
	if key == "" || value == nil {
		return CacheResponse{
			Ok:      false,
			Message: "cache key and Value are required",
		}
	}
	// expire default Value
	if expire == 0 {
		expire = 10
	}
	cacheKey := key + keyCode
	// set cache Value: mcCache.set(cacheKey, {Value: Value, expire: Date.now() + expire * 1000});
	// TODO: check cache set error
	mcCache[cacheKey] = CacheValue{
		value:  value,
		expire: uint(time.Now().Unix()) + expire,
	}
	// return successful response | TODO: or error response
	return CacheResponse{
		Ok:      true,
		Message: "task completed successfully",
		Value:   mcCache[cacheKey].value,
	}
}

func GetCache(key string) CacheResponse {
	// validate required params
	if key == "" {
		return CacheResponse{
			Ok:      false,
			Message: "cache key is required",
		}
	}
	cacheKey := key + keyCode
	cValue := mcCache[cacheKey]
	if cValue.value != nil && cValue.expire > uint(time.Now().Unix()) {
		return CacheResponse{
			Ok:      true,
			Message: "task completed successfully",
			Value:   cValue.value,
		}
	} else if cValue.value != nil && cValue.expire < uint(time.Now().Unix()) {
		// delete expired cache
		delete(mcCache, cacheKey)
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
	cacheKey := key + keyCode
	cValue := mcCache[cacheKey]
	if cValue.value != nil {
		delete(mcCache, cacheKey)
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

func ClearCache() CacheResponse {
	// clear mcCache map content
	mcCache = map[string]CacheValue{}
	//for key := range mcCache {
	//	delete(mcCache, key)
	//}
	return CacheResponse{
		Ok:      true,
		Message: "task completed successfully",
	}
}

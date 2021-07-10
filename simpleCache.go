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

func SetCache(key string, value ValueType, expire int64) CacheResponse {
	// validate required params
	if key == "" || value == nil {
		return CacheResponse{
			Ok:      false,
			Message: "cache key and Value are required",
		}
	}
	// expire default Value (in seconds)
	if expire == 0 {
		expire = 300
	}
	cacheKey := key + keyCode
	// set cache Value
	mcCache[cacheKey] = CacheValue{
		value:  value,
		expire: time.Now().Unix() + expire,
	}
	// return successful response
	if cValue, ok := mcCache[cacheKey]; ok {
		return CacheResponse{
			Ok:      true,
			Message: "task completed successfully",
			Value:   cValue.value,
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
	cacheKey := key + keyCode
	cValue, ok := mcCache[cacheKey]
	if (ok && cValue.value != nil) && cValue.expire > time.Now().Unix() {
		return CacheResponse{
			Ok:      true,
			Message: "task completed successfully",
			Value:   cValue.value,
		}
	} else if (ok && cValue.value != nil) && cValue.expire < time.Now().Unix() {
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
	if _, ok := mcCache[cacheKey]; ok {
		delete(mcCache, cacheKey)
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
	for key := range mcCache {
		delete(mcCache, key)
	}
	//mcCache = map[string]CacheValue{}
	return CacheResponse{
		Ok:      true,
		Message: "task completed successfully",
	}
}

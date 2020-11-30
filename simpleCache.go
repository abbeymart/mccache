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
			ok:      false,
			message: "cache key and value are required",
		}
	}
	// expire default value
	if expire == 0 {
		expire = 300
	}
	cacheKey := key + keyCode
	// set cache value: mcCache.set(cacheKey, {value: value, expire: Date.now() + expire * 1000});
	// TODO: check cache set error
	mcCache[cacheKey] = CacheValue{
		value:  value,
		expire: uint(time.Now().Unix()) + expire*1000,
		//expire: time.Unix() + expire.Seconds(),
	}
	// return successful response | TODO: or error response
	return CacheResponse{
		ok:      true,
		message: "task completed successfully",
		value:   mcCache[cacheKey].value,
	}
}

func getCache(key string) CacheResponse {
	// validate required params
	if key == "" {
		return CacheResponse{
			ok:      false,
			message: "cache key is required",
		}
	}
	cacheKey := key + keyCode
	cValue := mcCache[cacheKey]
	if cValue.value != nil && cValue.expire > uint(time.Now().Unix()) {
		return CacheResponse{
			ok:      true,
			message: "task completed successfully",
			value:   cValue.value,
		}
	} else if cValue.value != nil && cValue.expire < uint(time.Now().Unix()) {
		// delete expired cache
		delete(mcCache, cacheKey)
		return CacheResponse{
			ok:      false,
			message: "cache expired and deleted",
		}
	} else {
		return CacheResponse{
			ok:      false,
			message: "cache info does not exist",
		}
	}
}

func deleteCache(key string) CacheResponse {
	// validate required params
	if key == "" {
		return CacheResponse{
			ok:      false,
			message: "cache key is required",
		}
	}
	cacheKey := key + keyCode
	cValue := mcCache[cacheKey]
	if cValue.value != nil {
		delete(mcCache, cacheKey)
		return CacheResponse{
			ok:      true,
			message: "task completed successfully",
		}
	}
	return CacheResponse{
		ok:      false,
		message: "task not completed, cache-key-value not found",
	}
}

func clearCache() {
	// clear mcCache map content
	mcCache = map[string]CacheValue{}
	//for key := range mcCache {
	//	delete(mcCache, key)
	//}
}

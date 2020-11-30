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
			ok:      false,
			message: "cache key, hash and value are required",
		}
	}
	// expire default value
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
	// set cache value: mcHashCache.set(cacheKey, {value: value, expire: Date.now() + expire * 1000});
	// TODO: check cache set error
	mcHashCache[cacheKey] = hashCacheValue
	// return successful response | TODO: or error response
	return CacheResponse{
		ok:      true,
		message: "task completed successfully",
		value:   mcHashCache[cacheKey][hashKey].value,
	}
}

func GetHashCache(key string, hash string) CacheResponse {
	// validate required params
	if key == "" || hash == "" {
		return CacheResponse{
			ok:      false,
			message: "key and hash-key are required",
		}
	}
	cacheKey := key + keyCode
	hashKey := hash + keyCode
	cValue := mcHashCache[cacheKey][hashKey]
	if cValue.value != nil && cValue.expire > uint(time.Now().Unix()) {
		return CacheResponse{
			ok:      true,
			message: "task completed successfully",
			value:   cValue.value,
		}
	} else if cValue.value != nil && cValue.expire < uint(time.Now().Unix()) {
		// delete expired cache
		delete(mcHashCache, cacheKey)
		return CacheResponse{
			ok:      false,
			value:   nil,
			message: "cache expired and deleted",
		}
	} else {
		return CacheResponse{
			ok:      false,
			value:   nil,
			message: "cache info does not exist",
		}
	}
}

func DeleteHashCache(key string, hash string, by string) CacheResponse {
	// validate required params
	if key == "" || hash == "" && by == "hash" {
		return CacheResponse{
			ok:      false,
			message: "cache key is required",
		}
	}
	// by default value
	if by == "" {
		by = "hash"
	}
	cacheKey := key + keyCode
	hashKey := hash + keyCode

	if key != "" && by == "key" {
		cValue := mcHashCache[cacheKey]
		if cValue != nil {
			delete(mcHashCache, cacheKey)
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
	if key != "" && hash != "" && by == "hash" {
		cValue := mcHashCache[cacheKey][hashKey]
		if cValue.value != nil {
			delete(mcHashCache[cacheKey], hashKey)
			return CacheResponse{
				ok:      true,
				message: "task completed successfully",
			}
		}
		return CacheResponse{
			ok:      false,
			message: "task not completed, cache-key-hash-value not found",
		}
	}
	return CacheResponse{
		ok:      false,
		message: "task could not be completed due to incomplete inputs",
	}
}

func ClearHashCache() CacheResponse {
	// clear mcHashCache map content
	mcHashCache = map[string]HashCacheValueType{}
	//for key := range mcHashCache {
	//	delete(mcHashCache, key)
	//}
	return CacheResponse{
		ok:      true,
		message: "task completed successfully",
	}
}

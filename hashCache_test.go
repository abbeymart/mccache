// @Author: abbeymart | Abi Akindele | @Created: 2020-03-09 | @Updated: 2020-03-09
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect cache hash option - testing

package mccache

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestSetHashCache(t *testing.T) {
	jsonStr, _ := json.Marshal(cKeyValue)
	jsonStr2, _ := json.Marshal(cHashValue)
	cacheKey := string(jsonStr)
	hashKey := string(jsonStr2)
	fmt.Println("HASH-CACHE-TESTING:")
	fmt.Println("**********************")

	fmt.Println("should set and return valid cacheHashValue:")
	{
		setCacheRes := SetHashCache(cacheKey, hashKey, cacheValue, expiryTime)
		if !setCacheRes.ok {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.ok, true)
		}
		if setCacheRes.value != cacheValue {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.value, cacheValue)
		}
		if setCacheRes.message != okMsg {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.message, okMsg)
		}
		getCacheRes := GetHashCache(cacheKey, hashKey)
		if !getCacheRes.ok {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.ok, true)
		}
		if getCacheRes.value != cacheValue {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.value, cacheValue)
		}
		if getCacheRes.message != okMsg {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.message, okMsg)
		}
	}

	fmt.Println("should clear the cache and return nil/empty value:")
	{
		clearCacheRes := ClearHashCache()
		if !clearCacheRes.ok {
			t.Errorf("Got: %v, Expected: %v", clearCacheRes.ok, true)
		}
		if clearCacheRes.message != okMsg {
			t.Errorf("Got: %v, Expected: %v", clearCacheRes.message, okMsg)
		}
		getCacheRes2 := GetHashCache(cacheKey, hashKey)
		if getCacheRes2.ok {
			t.Errorf("Got: %v, Expected: %v", getCacheRes2.ok, false)
		}
		if getCacheRes2.value != nil {
			t.Errorf("Got: %v, Expected: %v", getCacheRes2.value, nil)
		}
		if getCacheRes2.message != notExistMsg {
			t.Errorf("Got: %v, Expected: %v", getCacheRes2.message, notExistMsg)
		}
	}

	fmt.Println("should set and return valid cacheValue -> before timeout/expiration:")
	{
		// change the expiry time to 2 seconds
		setCacheRes3 := SetHashCache(cacheKey, hashKey, cacheValue, 2)
		if !setCacheRes3.ok {
			t.Errorf("Got: %v, Expected: %v", setCacheRes3.ok, true)
		}
		if setCacheRes3.value != cacheValue {
			t.Errorf("Got: %v, Expected: %v", setCacheRes3.value, cacheValue)
		}
		if setCacheRes3.message != okMsg {
			t.Errorf("Got: %v, Expected: %v", setCacheRes3.message, okMsg)
		}
		getCacheRes3 := GetHashCache(cacheKey, hashKey)
		if !getCacheRes3.ok {
			t.Errorf("Got: %v, Expected: %v", getCacheRes3.ok, true)
		}
		if getCacheRes3.value != cacheValue {
			t.Errorf("Got: %v, Expected: %v", getCacheRes3.value, cacheValue)
		}
		if getCacheRes3.message != okMsg {
			t.Errorf("Got: %v, Expected: %v", getCacheRes3.message, okMsg)
		}
	}

	fmt.Println("should return nil value after timeout/expiration:")
	{
		time.Sleep(4 * time.Second)
		getCacheRes := GetHashCache(cacheKey, hashKey)
		if getCacheRes.ok {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.ok, false)
		}
		if getCacheRes.value != nil {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.value, nil)
		}
		if getCacheRes.message != expiredMsg {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.message, expiredMsg)
		}
	}

	fmt.Println("should set and return valid cacheValue, repeat prior to deleteCache testing:")
	{
		// change the expiry time to 10 seconds
		setCacheRes := SetHashCache(cacheKey, hashKey, cacheValue, 10)
		if !setCacheRes.ok {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.ok, true)
		}
		if setCacheRes.value != cacheValue {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.value, cacheValue)
		}
		if setCacheRes.message != okMsg {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.message, okMsg)
		}
		getCacheRes := GetHashCache(cacheKey, hashKey)
		if !getCacheRes.ok {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.ok, true)
		}
		if getCacheRes.value != cacheValue {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.value, cacheValue)
		}
		if getCacheRes.message != okMsg {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.message, okMsg)
		}
	}

	fmt.Println("should delete the cache and return nil/empty value:")
	{
		deleteCacheRes := DeleteHashCache(cacheKey, hashKey, "hash")
		if !deleteCacheRes.ok {
			t.Errorf("Got: %v, Expected: %v", deleteCacheRes.ok, true)
		}
		if deleteCacheRes.message != okMsg {
			t.Errorf("Got: %v, Expected: %v", deleteCacheRes.message, okMsg)
		}
		getCacheRes := GetHashCache(cacheKey, cacheKey)
		if getCacheRes.ok {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.ok, false)
		}
		if getCacheRes.value != nil {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.value, nil)
		}
		if getCacheRes.message != notExistMsg {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.message, notExistMsg)
		}
	}
}

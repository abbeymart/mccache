// @Author: abbeymart | Abi Akindele | @Created: 2020-03-09 | @Updated: 2020-03-09
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect cache - testing

package mccache

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestSetCache(t *testing.T) {
	jsonStr, _ := json.Marshal(cKeyValue)
	cacheKey := string(jsonStr)
	fmt.Println("SIMPLE-CACHE-TESTING:")
	fmt.Println("**********************")

	fmt.Println("should set and return valid cacheValue:")
	{
		setCacheRes := SetCache(cacheKey, cacheValue, expiryTime)
		if !setCacheRes.Ok {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.Ok, true)
		}
		if setCacheRes.Value != cacheValue {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.Value, cacheValue)
		}
		if setCacheRes.Message != okMsg {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.Message, okMsg)
		}
		getCacheRes := GetCache(cacheKey)
		if !getCacheRes.Ok {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.Ok, true)
		}
		if getCacheRes.Value != cacheValue {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.Value, cacheValue)
		}
		if getCacheRes.Message != okMsg {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.Message, okMsg)
		}
	}

	fmt.Println("should clear the cache and return nil/empty Value:")
	{
		clearCacheRes := ClearCache()
		if !clearCacheRes.Ok {
			t.Errorf("Got: %v, Expected: %v", clearCacheRes.Ok, true)
		}
		if clearCacheRes.Message != okMsg {
			t.Errorf("Got: %v, Expected: %v", clearCacheRes.Message, okMsg)
		}
		getCacheRes2 := GetCache(cacheKey)
		if getCacheRes2.Ok {
			t.Errorf("Got: %v, Expected: %v", getCacheRes2.Ok, false)
		}
		if getCacheRes2.Value != nil {
			t.Errorf("Got: %v, Expected: %v", getCacheRes2.Value, nil)
		}
		if getCacheRes2.Message != notExistMsg {
			t.Errorf("Got: %v, Expected: %v", getCacheRes2.Message, notExistMsg)
		}
	}

	fmt.Println("should set and return valid cacheValue -> before timeout/expiration:")
	{
		// change the expiry time to 2 seconds
		setCacheRes := SetCache(cacheKey, cacheValue, 2)
		if !setCacheRes.Ok {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.Ok, true)
		}
		if setCacheRes.Value != cacheValue {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.Value, cacheValue)
		}
		if setCacheRes.Message != okMsg {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.Message, okMsg)
		}
		getCacheRes3 := GetCache(cacheKey)
		if !getCacheRes3.Ok {
			t.Errorf("Got: %v, Expected: %v", getCacheRes3.Ok, true)
		}
		if getCacheRes3.Value != cacheValue {
			t.Errorf("Got: %v, Expected: %v", getCacheRes3.Value, cacheValue)
		}
		if getCacheRes3.Message != okMsg {
			t.Errorf("Got: %v, Expected: %v", getCacheRes3.Message, okMsg)
		}
	}

	fmt.Println("should return nil Value after timeout/expiration:")
	{
		time.Sleep(3 * time.Second)
		getCacheRes := GetCache(cacheKey)
		if getCacheRes.Ok {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.Ok, false)
		}
		if getCacheRes.Value != nil {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.Value, nil)
		}
		if getCacheRes.Message != expiredMsg {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.Message, expiredMsg)
		}
	}

	fmt.Println("should set and return valid cacheValue, repeat prior to deleteCache testing:")
	{
		// change the expiry time to 10 seconds
		setCacheRes := SetCache(cacheKey, cacheValue, 10)
		if !setCacheRes.Ok {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.Ok, true)
		}
		if setCacheRes.Value != cacheValue {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.Value, cacheValue)
		}
		if setCacheRes.Message != okMsg {
			t.Errorf("Got: %v, Expected: %v", setCacheRes.Message, okMsg)
		}
		getCacheRes := GetCache(cacheKey)
		if !getCacheRes.Ok {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.Ok, true)
		}
		if getCacheRes.Value != cacheValue {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.Value, cacheValue)
		}
		if getCacheRes.Message != okMsg {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.Message, okMsg)
		}
	}

	fmt.Println("should delete the cache and return nil/empty Value:")
	{
		deleteCacheRes := DeleteCache(cacheKey)
		if !deleteCacheRes.Ok {
			t.Errorf("Got: %v, Expected: %v", deleteCacheRes.Ok, true)
		}
		if deleteCacheRes.Message != okMsg {
			t.Errorf("Got: %v, Expected: %v", deleteCacheRes.Message, okMsg)
		}
		getCacheRes := GetCache(cacheKey)
		if getCacheRes.Ok {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.Ok, false)
		}
		if getCacheRes.Value != nil {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.Value, nil)
		}
		if getCacheRes.Message != notExistMsg {
			t.Errorf("Got: %v, Expected: %v", getCacheRes.Message, notExistMsg)
		}
	}
}

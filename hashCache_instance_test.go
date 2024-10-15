package mccache

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)
import "github.com/abbeymart/mctest"

func TestHashCacheInstance(t *testing.T) {
	jsonStr, _ := json.Marshal(cKeyValue)
	jsonStr2, _ := json.Marshal(cHashValue)
	jsonVal, _ := json.Marshal(cacheValue)
	cacheKey := string(jsonStr)
	hashKey := string(jsonStr2)

	cache := NewHashCache(1000, "test-hash-cache")

	fmt.Println("HASH-CACHE-INSTANCE-TESTING:")
	fmt.Println("****************************")

	mctest.McTest(mctest.OptionValue{
		Name: "should set and return valid cacheValue:",
		TestFunc: func() {
			setCacheRes := cache.SetCache(cacheKey, hashKey, cacheValue, expiryTime)
			mctest.AssertEquals(t, setCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, setCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
			mctest.AssertEquals(t, setCacheRes.Message, okMsg, "response should be: "+okMsg)
			getCacheRes := cache.GetCache(cacheKey, hashKey)
			mctest.AssertEquals(t, getCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, getCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
			mctest.AssertEquals(t, getCacheRes.Message, okMsg, "response should be: "+okMsg)
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should clear the cache and return nil/empty Value:",
		TestFunc: func() {
			clearCacheRes := cache.ClearCache()
			mctest.AssertEquals(t, clearCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, clearCacheRes.Message, okMsg, "response should be: "+okMsg)
			getCacheRes := cache.GetCache(cacheKey, hashKey)
			mctest.AssertEquals(t, getCacheRes.Ok, false, "response should be: false")
			mctest.AssertEquals(t, getCacheRes.Value, nil, "response should be: nil:")
			mctest.AssertEquals(t, getCacheRes.Message, notExistMsg, "response should be: "+notExistMsg)
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should set and return valid cacheValue -> before timeout/expiration:",
		TestFunc: func() {
			// change the expiry time to 2 seconds
			setCacheRes := cache.SetCache(cacheKey, hashKey, cacheValue, 2)
			mctest.AssertEquals(t, setCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, setCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
			mctest.AssertEquals(t, setCacheRes.Message, okMsg, "response should be: "+okMsg)
			getCacheRes := cache.GetCache(cacheKey, hashKey)
			mctest.AssertEquals(t, getCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, getCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
			mctest.AssertEquals(t, getCacheRes.Message, okMsg, "response should be: "+okMsg)
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should return nil Value after timeout/expiration:",
		TestFunc: func() {
			time.Sleep(4 * time.Second)
			getCacheRes := cache.GetCache(cacheKey, hashKey)
			mctest.AssertEquals(t, getCacheRes.Ok, false, "response should be: false")
			mctest.AssertEquals(t, getCacheRes.Value, nil, "response should be: nil")
			mctest.AssertEquals(t, getCacheRes.Message, expiredMsg, "response should be: "+expiredMsg)
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should set and return valid cacheValue, repeat prior to deleteCache testing:",
		TestFunc: func() {
			// change the expiry time to 10 seconds
			setCacheRes := cache.SetCache(cacheKey, hashKey, cacheValue, 10)
			mctest.AssertEquals(t, setCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, setCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
			mctest.AssertEquals(t, setCacheRes.Message, okMsg, "response should be: "+okMsg)
			getCacheRes := cache.GetCache(cacheKey, hashKey)
			mctest.AssertEquals(t, getCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, getCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
			mctest.AssertEquals(t, getCacheRes.Message, okMsg, "response should be: "+okMsg)
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should delete the cache and return nil/empty Value:",
		TestFunc: func() {
			deleteCacheRes := cache.DeleteCache(cacheKey, hashKey, "hash")
			mctest.AssertEquals(t, deleteCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, deleteCacheRes.Message, okMsg, "response should be: "+okMsg)
			getCacheRes := cache.GetCache(cacheKey, hashKey)
			mctest.AssertEquals(t, getCacheRes.Ok, false, "response should be: false")
			mctest.AssertEquals(t, getCacheRes.Value, nil, "response should be: nil")
			mctest.AssertEquals(t, getCacheRes.Message, notExistMsg, "response should be: "+notExistMsg)
		},
	})

	mctest.PostTestResult()
}

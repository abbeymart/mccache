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

	var results []mctest.UnitTestResult

	fmt.Println("HASH-CACHE-TESTING:")
	fmt.Println("****************************")

	test1 := mctest.NewTest(mctest.ParamsType{
		Name: "should set and return valid cacheValue:",
	})
	test1.SetTestFunction(func() {
		setCacheRes := cache.SetCache(cacheKey, hashKey, cacheValue, expiryTime)
		test1.AssertEquals(setCacheRes.Ok, true, "response should be: true")
		test1.AssertEquals(setCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
		test1.AssertEquals(setCacheRes.Message, okMsg, "response should be: "+okMsg)
		getCacheRes := cache.GetCache(cacheKey, hashKey)
		test1.AssertEquals(getCacheRes.Ok, true, "response should be: true")
		test1.AssertEquals(getCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
		test1.AssertEquals(getCacheRes.Message, okMsg, "response should be: "+okMsg)
	})
	test1Result := test1.RunTest()
	results = append(results, test1Result)

	test2 := mctest.NewTest(mctest.ParamsType{
		Name: "should clear the cache and return nil/empty ItemValue:",
	})
	test2.SetTestFunction(func() {
		clearCacheRes := cache.ClearCache()
		test2.AssertEquals(clearCacheRes.Ok, true, "response should be: true")
		test2.AssertEquals(clearCacheRes.Message, okMsg, "response should be: "+okMsg)
		getCacheRes := cache.GetCache(cacheKey, hashKey)
		test2.AssertEquals(getCacheRes.Ok, false, "response should be: false")
		test2.AssertEquals(getCacheRes.Value, nil, "response should be: nil:")
		test2.AssertEquals(getCacheRes.Message, notExistMsg, "response should be: "+notExistMsg)
	})
	test2Result := test2.RunTest()
	results = append(results, test2Result)

	test3 := mctest.NewTest(mctest.ParamsType{
		Name: "should set and return valid cacheValue -> before timeout/expiration:",
	})
	test3.SetTestFunction(func() {
		// change the expiry time to 2 seconds
		setCacheRes := cache.SetCache(cacheKey, hashKey, cacheValue, 2)
		test3.AssertEquals(setCacheRes.Ok, true, "response should be: true")
		test3.AssertEquals(setCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
		test3.AssertEquals(setCacheRes.Message, okMsg, "response should be: "+okMsg)
		getCacheRes := cache.GetCache(cacheKey, hashKey)
		test3.AssertEquals(getCacheRes.Ok, true, "response should be: true")
		test3.AssertEquals(getCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
		test3.AssertEquals(getCacheRes.Message, okMsg, "response should be: "+okMsg)
	})
	test3Result := test3.RunTest()
	results = append(results, test3Result)

	test4 := mctest.NewTest(mctest.ParamsType{
		Name: "should return nil ItemValue after timeout/expiration:",
	})
	time.Sleep(4 * time.Second)
	test4.SetTestFunction(func() {
		getCacheRes := cache.GetCache(cacheKey, hashKey)
		test4.AssertEquals(getCacheRes.Ok, false, "response should be: false")
		test4.AssertEquals(getCacheRes.Value, nil, "response should be: nil")
		test4.AssertEquals(getCacheRes.Message, expiredMsg, "response should be: "+expiredMsg)
	})
	test4Result := test4.RunTest()
	results = append(results, test4Result)

	test5 := mctest.NewTest(mctest.ParamsType{
		Name: "should set and return valid cacheValue, repeat prior to deleteCache testing:",
	})
	test5.SetTestFunction(func() {
		// change the expiry time to 10 seconds
		setCacheRes := cache.SetCache(cacheKey, hashKey, cacheValue, 10)
		test5.AssertEquals(setCacheRes.Ok, true, "response should be: true")
		test5.AssertEquals(setCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
		test5.AssertEquals(setCacheRes.Message, okMsg, "response should be: "+okMsg)
		getCacheRes := cache.GetCache(cacheKey, hashKey)
		test5.AssertEquals(getCacheRes.Ok, true, "response should be: true")
		test5.AssertEquals(getCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
		test5.AssertEquals(getCacheRes.Message, okMsg, "response should be: "+okMsg)
	})
	test5Result := test5.RunTest()
	results = append(results, test5Result)

	test6 := mctest.NewTest(mctest.ParamsType{
		Name: "should delete the cache and return nil/empty ItemValue:",
	})
	test6.SetTestFunction(func() {
		deleteCacheRes := cache.DeleteCache(cacheKey, hashKey, "hash")
		test6.AssertEquals(deleteCacheRes.Ok, true, "response should be: true")
		test6.AssertEquals(deleteCacheRes.Message, okMsg, "response should be: "+okMsg)
		getCacheRes := cache.GetCache(cacheKey, hashKey)
		test6.AssertEquals(getCacheRes.Ok, false, "response should be: false")
		test6.AssertEquals(getCacheRes.Value, nil, "response should be: nil")
		test6.AssertEquals(getCacheRes.Message, notExistMsg, "response should be: "+notExistMsg)
	})

	test6Result := test6.RunTest()
	results = append(results, test6Result)

	mctest.TestResult(results)
}

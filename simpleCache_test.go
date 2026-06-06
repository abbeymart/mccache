// @Author: abbeymart | Abi Akindele | @Created: 2020-03-09 | @Updated: 2026-06-05
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect cache - testing

package mccache

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)
import "github.com/abbeymart/mctest"

func TestCache(t *testing.T) {
	jsonStr, _ := json.Marshal(cKeyValue)
	cacheKey := string(jsonStr)
	jsonVal, _ := json.Marshal(cacheValue)

	var results []mctest.UnitTestResult

	fmt.Println("SIMPLE-CACHE-TESTING:")
	fmt.Println("**********************")

	test1 := mctest.NewTest(mctest.ParamsType{
		Name: "should set and return valid cacheValue:",
	})
	test1.SetTestFunction(func() {
		setCacheRes := SetCache(cacheKey, cacheValue, expiryTime)
		test1.AssertEquals(setCacheRes.Ok, true, "response should be: true")
		test1.AssertEquals(setCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
		test1.AssertEquals(setCacheRes.Message, okMsg, "response should be: "+okMsg)
		getCacheRes := GetCache(cacheKey)
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
		clearCacheRes := ClearCache()
		test2.AssertEquals(clearCacheRes.Ok, true, "response should be: true")
		test2.AssertEquals(clearCacheRes.Message, okMsg, "response should be: "+okMsg)
		getCacheRes := GetCache(cacheKey)
		test2.AssertEquals(getCacheRes.Ok, false, "response should be: false")
		test2.AssertEquals(getCacheRes.Value, nil, "response should be: nil")
		test2.AssertEquals(getCacheRes.Message, notExistMsg, "response should be: "+notExistMsg)
	})
	test2Result := test2.RunTest()
	results = append(results, test2Result)

	test3 := mctest.NewTest(mctest.ParamsType{
		Name: "should set and return valid cacheValue -> before timeout/expiration:",
	})
	test3.SetTestFunction(func() {
		// change the expiry time to 2 seconds
		setCacheRes := SetCache(cacheKey, cacheValue, 2)
		test3.AssertEquals(setCacheRes.Ok, true, "response should be: true")
		test3.AssertEquals(setCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
		test3.AssertEquals(setCacheRes.Message, okMsg, "response should be: "+okMsg)
		getCacheRes := GetCache(cacheKey)
		test3.AssertEquals(getCacheRes.Ok, true, "response should be: true")
		test3.AssertEquals(getCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
		test3.AssertEquals(getCacheRes.Message, okMsg, "response should be: "+okMsg)
	})
	test3Result := test3.RunTest()
	results = append(results, test3Result)

	test4 := mctest.NewTest(mctest.ParamsType{
		Name: "should return nil ItemValue after timeout/expiration:",
	})
	time.Sleep(3 * time.Second)
	test4.SetTestFunction(func() {
		getCacheRes := GetCache(cacheKey)
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
		setCacheRes := SetCache(cacheKey, cacheValue, 10)
		test5.AssertEquals(setCacheRes.Ok, true, "response should be: true")
		test5.AssertEquals(setCacheRes.Value, cacheValue, "response should be: "+string(jsonVal))
		test5.AssertEquals(setCacheRes.Message, okMsg, "response should be: "+okMsg)
		getCacheRes := GetCache(cacheKey)
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
		deleteCacheRes := DeleteCache(cacheKey)
		test6.AssertEquals(deleteCacheRes.Ok, true, "response should be: true")
		test6.AssertEquals(deleteCacheRes.Message, okMsg, "response should be: "+okMsg)
		getCacheRes := GetCache(cacheKey)
		test6.AssertEquals(getCacheRes.Ok, false, "response should be: false")
		test6.AssertEquals(getCacheRes.Value, nil, "response should be: nil:")
		test6.AssertEquals(getCacheRes.Message, notExistMsg, "response should be: "+notExistMsg)
	})
	test6Result := test6.RunTest()
	results = append(results, test6Result)

	mctest.TestResult(results)
}

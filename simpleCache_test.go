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
import "github.com/abbeymart/mctestgo"

func TestSetCache(t *testing.T) {
	jsonStr, _ := json.Marshal(cKeyValue)
	cacheKey := string(jsonStr)
	jsonVal, _ := json.Marshal(cacheValue)

	fmt.Println("SIMPLE-CACHE-TESTING:")
	fmt.Println("**********************")

	mctest.McTest(mctest.OptionValue{
		Name: "should set and return valid cacheValue:",
		TestFunc: func() {
			setCacheRes := SetCache(cacheKey, cacheValue, expiryTime)
			mctest.AssertEquals(t, setCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, setCacheRes.Value, cacheValue, "response should be: " + string(jsonVal))
			mctest.AssertEquals(t, setCacheRes.Message, okMsg, "response should be: " + okMsg)
			getCacheRes := GetCache(cacheKey)
			mctest.AssertEquals(t, getCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, getCacheRes.Value, cacheValue, "response should be: " + string(jsonVal))
			mctest.AssertEquals(t, getCacheRes.Message, okMsg, "response should be: " + okMsg)
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should clear the cache and return nil/empty Value:",
		TestFunc: func() {
			clearCacheRes := ClearCache()
			mctest.AssertEquals(t, clearCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, clearCacheRes.Message, okMsg, "response should be: " + okMsg)
			getCacheRes := GetCache(cacheKey)
			mctest.AssertEquals(t, getCacheRes.Ok, false, "response should be: false")
			mctest.AssertEquals(t, getCacheRes.Value, nil, "response should be: nil")
			mctest.AssertEquals(t, getCacheRes.Message, notExistMsg, "response should be: " + notExistMsg)
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should set and return valid cacheValue -> before timeout/expiration:",
		TestFunc: func() {
			// change the expiry time to 2 seconds
			setCacheRes := SetCache(cacheKey, cacheValue, 2)
			mctest.AssertEquals(t, setCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, setCacheRes.Value, cacheValue, "response should be: " + string(jsonVal))
			mctest.AssertEquals(t, setCacheRes.Message, okMsg, "response should be: " + okMsg)
			getCacheRes := GetCache(cacheKey)
			mctest.AssertEquals(t, getCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, getCacheRes.Value, cacheValue, "response should be: " + string(jsonVal))
			mctest.AssertEquals(t, getCacheRes.Message, okMsg, "response should be: " + okMsg)
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should return nil Value after timeout/expiration:",
		TestFunc: func() {
			time.Sleep(3 * time.Second)
			getCacheRes := GetCache(cacheKey)
			mctest.AssertEquals(t, getCacheRes.Ok, false, "response should be: false")
			mctest.AssertEquals(t, getCacheRes.Value, nil, "response should be: nil")
			mctest.AssertEquals(t, getCacheRes.Message, expiredMsg, "response should be: " + expiredMsg)
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should set and return valid cacheValue, repeat prior to deleteCache testing:",
		TestFunc: func() {
			// change the expiry time to 10 seconds
			setCacheRes := SetCache(cacheKey, cacheValue, 10)
			mctest.AssertEquals(t, setCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, setCacheRes.Value, cacheValue, "response should be: " + string(jsonVal))
			mctest.AssertEquals(t, setCacheRes.Message, okMsg, "response should be: " + okMsg)
			getCacheRes := GetCache(cacheKey)
			mctest.AssertEquals(t, getCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, getCacheRes.Value, cacheValue, "response should be: " + string(jsonVal))
			mctest.AssertEquals(t, getCacheRes.Message, okMsg, "response should be: " + okMsg)
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should delete the cache and return nil/empty Value:",
		TestFunc: func() {
			deleteCacheRes := DeleteCache(cacheKey)
			mctest.AssertEquals(t, deleteCacheRes.Ok, true, "response should be: true")
			mctest.AssertEquals(t, deleteCacheRes.Message, okMsg, "response should be: " + okMsg)
			getCacheRes := GetCache(cacheKey)
			mctest.AssertEquals(t, getCacheRes.Ok, false, "response should be: false")
			mctest.AssertEquals(t, getCacheRes.Value, nil, "response should be: nil:")
			mctest.AssertEquals(t, getCacheRes.Message, notExistMsg, "response should be: " + notExistMsg)
		},
	})

	mctest.PostTestResult()
}

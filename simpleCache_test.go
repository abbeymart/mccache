// @Author: abbeymart | Abi Akindele | @Created: 2020-03-09 | @Updated: 2020-03-09
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect cache - testing

package mccache

import (
	"encoding/json"
	"fmt"
	"testing"
)

// test data
var cacheValue = struct {
	firstName string
	lastName  string
	location  string
}{
	firstName: "Abi",
	lastName:  "Akindele",
	location:  "Toronto-Canada",
}

var cKeyValue = map[string]interface{}{
	"name":     "Abi",
	"location": "Toronto-Canada",
}

//var cacheKeyValue = struct {
//	name     string
//	location string
//}{
//	name:     "Abi",
//	location: "Toronto-Canada",
//}

const expiryTime = 5 // 5 seconds

const okMsg = "task completed successfully"

func TestSetCache(t *testing.T) {
	jsonStr, _ := json.Marshal(cKeyValue)
	cacheKey := string(jsonStr)
	fmt.Println("should set and return valid cacheValue:")
	setCacheRes := SetCache(cacheKey, cacheValue, expiryTime)
	if !setCacheRes.ok {
		t.Errorf("Got: %v, Expected: %v", setCacheRes.ok, true)
	}
	if setCacheRes.value != cacheValue {
		t.Errorf("Got: %v, Expected: %v", setCacheRes.value, cacheValue)
	}
	if setCacheRes.message != okMsg {
		t.Errorf("Got: %v, Expected: %v", setCacheRes.message, okMsg)
	}
	getCacheRes := getCache(cacheKey)
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

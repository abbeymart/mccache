// @Author: abbeymart | Abi Akindele | @Created: 2020-11-29 | @Updated: 2020-11-29
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mccache

type ValueType interface{}

type CacheValue struct {
	value  ValueType
	expire uint
}

type CacheResponse struct {
	ok      bool
	message string
	value   ValueType
}

type HashCacheValueType map[string]CacheValue

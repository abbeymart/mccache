// @Author: abbeymart | Abi Akindele | @Created: 2020-11-29 | @Updated: 2020-11-29
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mccache

type ValueType interface{}

type CacheValue struct {
	value  ValueType
	expire int64
}

type CacheResponse struct {
	Ok      bool
	Message string
	Value   ValueType
}

type HashCacheValueType map[string]CacheValue

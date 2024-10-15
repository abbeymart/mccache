// @Author: abbeymart | Abi Akindele | @Created: 2020-11-29 | @Updated: 2020-11-29
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mccache

import (
	"container/list"
	"sync"
)

type ValueType interface{}

type CacheValue struct {
	value     ValueType
	expire    int64
	createdAt int64
}

type CacheResponse struct {
	Ok      bool
	Message string
	Value   ValueType
}

type HashCacheValueType map[string]CacheValue

type SimpleCache struct {
	mu         sync.RWMutex
	items      map[string]CacheValue
	capacity   int // placeholder - for setting the default cache capacity
	iItems     map[string]*list.Element
	iItemsList *list.List
	secretCode string
}

type HashCache struct {
	mu         sync.RWMutex
	items      map[string]HashCacheValueType
	capacity   int // placeholder - for setting the default cache capacity
	iItems     map[string]*list.Element
	iItemsList *list.List
	secretCode string
}

const (
	ByKey  = "key"
	ByHash = "hash"
)

// @Author: abbeymart | Abi Akindele | @Created: 2020-11-30 | @Updated: 2020-11-30
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mccache

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

const (
	okMsg                  = "task completed successfully"
	notExistMsg            = "cache info does not exist"
	expiredMsg             = "cache expired and deleted"
	//notCompleteMsg         = "task not completed, cache-key-value not found"
	//cacheKeyMsg            = "cache key is required"
	//cacheExpiredDeletedMsg = "cache expired and deleted"
)

// test data, in addition to simpleCache test-data
var cHashValue = map[string]interface{}{
	"hash1": "Hash1",
	"hash2": "Hash2",
}

# kl2-cache
With kl2-cache, You can cache any struct that can be serialized by json.Unmarshal

Usage: 
1. go get github.com/KL-Engineering/kl2-cache
2. kl2cache.Init
3. kl2cache.DefaultProvider.Get
4. kl2cache.DefaultProvider.BatchGet

# examples
refer to [redis_test](./redis_test.go)
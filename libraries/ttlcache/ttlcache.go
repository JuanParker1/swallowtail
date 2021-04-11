package ttlcache

import (
	"sync"
	"time"
)

type cacheResult struct {
	expirationDate time.Time
	value          interface{}
}

func (ttlr *cacheResult) HasExpired() bool {
	return ttlr.expirationDate.Before(time.Now())
}

// Move to own library
type TTLCache struct {
	ttl   time.Duration
	cache map[string]*cacheResult
	mu    sync.RWMutex
}

func New(ttl time.Duration) *TTLCache {
	return &TTLCache{
		cache: map[string]*cacheResult{},
		ttl:   ttl,
	}
}

// Get retreives the value for the given key and whether it is expired or not
func (tc *TTLCache) Get(key string) (interface{}, bool) {
	tc.mu.RLock()
	defer tc.mu.RUnlock()
	r, ok := tc.cache[key]
	if !ok {
		return nil, false
	}
	return r.value, r.HasExpired()
}

// Set sets the value to the given key. It is thread-safe
func (tc *TTLCache) Set(key string, value interface{}) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	r := &cacheResult{
		expirationDate: time.Now().Add(tc.ttl),
		value:          value,
	}
	tc.cache[key] = r
}

//GetAndRefreshExpiry will get the key, if expired it will update the expiry data to now plus the default ttl
// the first return value is the value stored for that key, and the second is whether it is expired or not.
func (tc *TTLCache) GetAndRefreshExpiry(key string) (interface{}, bool) {
	v, expired := tc.Get(key)
	if !expired {
		return v, false
	}
	tc.Set(key, v)
	return v, true
}

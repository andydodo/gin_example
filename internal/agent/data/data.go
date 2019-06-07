package data

import (
	"sync"
)

type DetectedItem struct {
}

type CheckResult struct {
}

type DetectedItemSafeMap struct {
	sync.RWMutex
	M map[string][]*DetectedItem
}

var (
	DetectedItemMap = &DetectedItemSafeMap{M: make(map[string][]*DetectedItem)}
)

func (this *DetectedItemSafeMap) Get(key string) ([]*DetectedItem, bool) {
	this.RLock()
	defer this.RUnlock()
	ipItem, exists := this.M[key]
	return ipItem, exists
}

func (this *DetectedItemSafeMap) Set(detectedItemMap map[string][]*DetectedItem) {
	this.Lock()
	defer this.Unlock()
	this.M = detectedItemMap
}

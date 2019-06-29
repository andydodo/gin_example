package data

import (
	"sync"
)

type CheckResult struct {
	Name   string
	Url    string
	Status int
	Code   int
}

type DetectedItem struct {
	Url string `json:"url"`
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
	Item, exists := this.M[key]
	return Item, exists
}

func (this *DetectedItemSafeMap) Set(detectedItemMap map[string][]*DetectedItem) {
	this.Lock()
	defer this.Unlock()
	this.M = detectedItemMap
}

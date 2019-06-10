package cron

import (
	"fmt"
	"github.com/LIYINGZHEN/ginexample/internal/agent/data"
	"github.com/LIYINGZHEN/ginexample/internal/app/postgres"
	"github.com/LIYINGZHEN/ginexample/internal/app/types"
	"log"
	"time"
)

type Agent struct {
	postgres.Repository
}

func New(repository *postgres.Repository) *Agent {
	return &Agent{
		*repository,
	}
}

func (a *Agent) GetItem() {
	t1 := time.NewTicker(time.Duration(30) * time.Second)
	for {
		err := a.getItem()
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}
		<-t1.C
	}
}

func (a *Agent) getItem() error {
	detectedItemMap := make(map[string][]*data.DetectedItem)
	items, err := a.LinkRepository.FindAll()
	if err != nil {
		return fmt.Errorf("get items error:", err)
	}
	for _, item := range items {
		detectedItem := newDetectedItem(item)
		key := item.Name

		if _, exists := detectedItemMap[key]; exists {
			detectedItemMap[key] = append(detectedItemMap[key], &detectedItem)
		} else {
			detectedItemMap[key] = []*data.DetectedItem{&detectedItem}
		}
	}

	for k, v := range detectedItemMap {
		log.Println(k)
		for _, i := range v {
			log.Println(i)
		}
	}

	data.DetectedItemMap.Set(detectedItemMap)
	return nil
}

func newDetectedItem(item types.Link) data.DetectedItem {
	detectedItem := data.DetectedItem{
		Url: item.Url,
	}

	return detectedItem
}

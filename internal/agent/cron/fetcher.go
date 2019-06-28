package cron

import (
	"github.com/LIYINGZHEN/ginexample/internal/agent/data"
	"github.com/LIYINGZHEN/ginexample/internal/app/postgres"
	"github.com/LIYINGZHEN/ginexample/internal/app/types"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"

	list "github.com/LIYINGZHEN/ginexample/pkg/utils"
)

var (
	CheckResultQueue *list.SafeLinkedList
	WorkerChan       chan int
)

func init() {
	WorkerChan = make(chan int, 500)
	CheckResultQueue = list.NewSafeLinkedList()
}

type Agent struct {
	postgres.Repository
	*log.Logger
}

func New(repository *postgres.Repository, logger *log.Logger) *Agent {
	return &Agent{
		*repository,
		logger,
	}
}

func (a *Agent) StartCheck() {
	t1 := time.NewTicker(time.Duration(30) * time.Second)
	for {
		items, err := a.getItem()
		if err != nil || items == nil {
			continue
		}

		for _, item := range items {
			WorkerChan <- 1
			go CheckTargetStatus(item)
		}
		<-t1.C
	}
}

func (a *Agent) getItem() ([]types.Link, error) {
	items, err := a.LinkRepository.FindAll()
	if err != nil {
		a.Printf("get items from db failed: %v", err)
		return nil, errors.Wrap(err, "get items from db failed")
	}

	return items, nil
}

func CheckTargetStatus(item types.Link) {
	defer func() {
		<-WorkerChan
	}()

	checkResult := checkTargetStatus(item)
	CheckResultQueue.PushFront(checkResult)
}

//todo: check health
func checkTargetStatus(item types.Link) (itemCheckResult *data.CheckResult) {
	itemCheckResult = &data.CheckResult{
		Name: item.Name,
		Url:  item.Url,
		Code: 0,
	}

	if item.Enable {
		resp, err := http.Get("http://" + item.Url)

		if err != nil {
			itemCheckResult.Status = 0
			return
		} else {
			itemCheckResult.Status = 1
			itemCheckResult.Code = resp.StatusCode
		}

		if resp.Body != nil {
			defer resp.Body.Close()
		}
	}

	return
}

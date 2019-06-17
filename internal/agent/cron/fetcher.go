package cron

import (
	"github.com/LIYINGZHEN/ginexample/internal/agent/data"
	"github.com/LIYINGZHEN/ginexample/internal/app/postgres"
	"github.com/LIYINGZHEN/ginexample/internal/app/types"
	"github.com/pkg/errors"
	"log"
	"time"

	list "github.com/LIYINGZHEN/ginexample/pkg/utils"
)

var (
	CheckResultQueue *list.SafeLinkedList
	WorkerChan       chan int
)

func init() {
	WorkerChan = make(chan int, 50)
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
		items, _ := a.getItem()

		for _, item := range items {
			WorkerChan <- 1
			go CheckTargetStatus(item)
		}
		<-t1.C
	}
}

func (a *Agent) getItem() (*[]types.Link, error) {
	items, err := a.LinkRepository.FindAll()
	if err != nil {
		a.Printf("get items from db failed: %v", err)
		return nil, errors.Wrap(err, "get items from db failed")
	}

	return &items, nil
}

func CheckTargetStatus(item *types.Link) {
	defer func() {
		<-WorkerChan
	}()

	checkResult := checkTargetStatus(item)
	CheckResultQueue.PushFront(checkResult)
}

//todo: check health
func checkTargetStatus(item *types.Link) (itemCheckResult *types.Link) {
	itemCheckResult = &CheckResult{
		Sid:      item.Sid,
		Domain:   item.Domain,
		Creator:  item.Creator,
		Tag:      item.Tag,
		Target:   item.Target,
		Ip:       item.Ip,
		RespTime: item.Timeout,
		RespCode: "0",
  }
   
  return
  }
package cron

import (
	"fmt"
	"time"
)

func GetItem() {
	t1 := time.NewTicker(time.Duration(30) * time.Second)
	for {
		err := getItem()
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}
		<-t1.C
	}
}

func getItem() error {
	items, err := db.FindAll()
	if err != nil {
		return fmt.Errorf("get items error:", err)
	}
	for _, item := range items {
		_, domain, _, _ := utils.ParseUrl(item.Url)
		var ipIdcArr []g.IpIdc
		if s.IP != "" {
			ips := strings.Split(s.IP, ",")
			for _, ip := range ips {
				var tmp g.IpIdc
				tmp.Ip = ip
				tmp.Idc = "default"
				ipIdcArr = append(ipIdcArr, tmp)
			}
		} else {
			ipIdcArr = getIpAndIdc(domain)
		}

		for _, tmp := range ipIdcArr {
			detectedItem := newDetectedItem(s, tmp.Ip, tmp.Idc)
			key := utils.Getkey(tmp.Idc, int(detectedItem.Sid))

			if _, exists := detectedItemMap[key]; exists {
				detectedItemMap[key] = append(detectedItemMap[key], &detectedItem)
			} else {
				detectedItemMap[key] = []*g.DetectedItem{&detectedItem}
			}
		}
	}

	for k, v := range detectedItemMap {
		log.Println(k)
		for _, i := range v {
			log.Println(i)
		}
	}

	g.DetectedItemMap.Set(detectedItemMap)
	return nil
}

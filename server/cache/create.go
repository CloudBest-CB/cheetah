package cache

import (
	"../db"
	"../model"
	"fmt"
	"time"
)

var cacheQueue *LoopQueue

func CreateCacheQueue() error {
	fmt.Println("---------开始监控cache数据")
	cacheQueue = CreateLoopQueue(1000)
	go SentToDb()
	return nil
}

func SaveInCache(e []*model.MetaData) {
	cacheQueue.EnQueue(e)
	//fmt.Println(cacheQueue)
}
func SentToDb() {
	for {
		if !cacheQueue.IsEmpty() {
			data := cacheQueue.GetFront()
			for _, value := range data.([]*model.MetaData) {
				err := db.SaveInformation(value)
				if err != nil {
					fmt.Println("influxdb cha ru chu cuo", err)
				}
			}
			_ = cacheQueue.DeQueue()
			//db.SaveInformation(data.(*model.MetaData))
			//if err==nil {
			//	_=cacheQueue.DeQueue()
			//}
			//db.SaveInformation(cacheQueue.DeQueue().(*model.MetaData))
		} else {
			time.Sleep(10 * time.Second)
		}
		if cacheQueue.GetSize() > 200 {
			fmt.Println("---------cache中数据量过高", cacheQueue.GetSize())
		}
		fmt.Println("---------cache中数据量", cacheQueue.GetSize())
	}
}
func sentToDb() {

}

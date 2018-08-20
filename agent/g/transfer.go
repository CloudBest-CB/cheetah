// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package g

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"../model"
)

var (
	TransferClientsLock = new(sync.RWMutex)
	TransferClients     = map[string]*SingleConnRpcClient{}
)

func SendMetrics(metrics []*model.MetricValue, resp *model.TransferResponse) {
	rand.Seed(time.Now().UnixNano())
	addr := Config().Server.Addrs
	c := getTransferClient(addr)
	if c == nil {
		c = initTransferClient(addr)
	}

	updateMetrics(c, metrics, resp)

}

func initTransferClient(addr string) *SingleConnRpcClient {
	var c = &SingleConnRpcClient{
		RpcServer: addr,
		Timeout:   time.Duration(Config().Server.Timeout) * time.Millisecond,
	}
	TransferClientsLock.Lock()
	defer TransferClientsLock.Unlock()
	TransferClients[addr] = c

	return c
}

func updateMetrics(c *SingleConnRpcClient, metrics []*model.MetricValue, resp *model.TransferResponse) bool {
	err := c.Call("Server.Update", metrics, resp)
	if err != nil {
		log.Println("call Server.Update fail:", c, err)
		return false
	}
	return true
}

func getTransferClient(addr string) *SingleConnRpcClient {
	TransferClientsLock.RLock()
	defer TransferClientsLock.RUnlock()

	if c, ok := TransferClients[addr]; ok {
		return c
	}
	return nil
}

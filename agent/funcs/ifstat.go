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

package funcs

import (
	"../g"
	"../model"
	"github.com/toolkits/nux"
	"log"
)

func NetMetrics() []*model.MetricValue {
	return CoreNetMetrics(g.Config().Collector.IfacePrefix)
}

func CoreNetMetrics(ifacePrefix []string) []*model.MetricValue {

	netIfs, err := nux.NetIfs(ifacePrefix)
	if err != nil {
		log.Println(err)
		return []*model.MetricValue{}
	}

	cnt := len(netIfs)
	ret := make([]*model.MetricValue, cnt*23)

	for idx, netIf := range netIfs {
		iface := "iface=" + netIf.Iface
		ret[idx*23+0] = CounterValue("net.in.bytes", netIf.InBytes, iface)
		ret[idx*23+1] = CounterValue("net.in.packets", netIf.InPackages, iface)
		ret[idx*23+2] = CounterValue("net.in.errors", netIf.InErrors, iface)
		ret[idx*23+3] = CounterValue("net.in.dropped", netIf.InDropped, iface)
		ret[idx*23+4] = CounterValue("net.in.fifo.errs", netIf.InFifoErrs, iface)
		ret[idx*23+5] = CounterValue("net.in.frame.errs", netIf.InFrameErrs, iface)
		ret[idx*23+6] = CounterValue("net.in.compressed", netIf.InCompressed, iface)
		ret[idx*23+7] = CounterValue("net.in.multicast", netIf.InMulticast, iface)
		ret[idx*23+8] = CounterValue("net.out.bytes", netIf.OutBytes, iface)
		ret[idx*23+9] = CounterValue("net.out.packets", netIf.OutPackages, iface)
		ret[idx*23+10] = CounterValue("net.out.errors", netIf.OutErrors, iface)
		ret[idx*23+11] = CounterValue("net.out.dropped", netIf.OutDropped, iface)
		ret[idx*23+12] = CounterValue("net.out.fifo.errs", netIf.OutFifoErrs, iface)
		ret[idx*23+13] = CounterValue("net.out.collisions", netIf.OutCollisions, iface)
		ret[idx*23+14] = CounterValue("net.out.carrier.errs", netIf.OutCarrierErrs, iface)
		ret[idx*23+15] = CounterValue("net.out.compressed", netIf.OutCompressed, iface)
		ret[idx*23+16] = CounterValue("net.total.bytes", netIf.TotalBytes, iface)
		ret[idx*23+17] = CounterValue("net.total.packets", netIf.TotalPackages, iface)
		ret[idx*23+18] = CounterValue("net.total.errors", netIf.TotalErrors, iface)
		ret[idx*23+19] = CounterValue("net.total.dropped", netIf.TotalDropped, iface)
		ret[idx*23+20] = GaugeValue("net.speed.bits", netIf.SpeedBits, iface)
		ret[idx*23+21] = CounterValue("net.in.percent", netIf.InPercent, iface)
		ret[idx*23+22] = CounterValue("net.out.percent", netIf.OutPercent, iface)
	}
	return ret
}

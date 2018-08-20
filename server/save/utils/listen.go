package utils

import (
	"../g"
	"../influx"
	"../model"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
	"time"
)

type Server int
type Agent int

func Start() {
	addr := g.Config().Listen

	server := rpc.NewServer()
	// server.Register(new(filter.Filter))
	server.Register(new(Agent))
	server.Register(new(Server))

	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatalln("listen error:", e)
	} else {
		log.Println("listening", addr)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("listener accept fail:", err)
			time.Sleep(time.Duration(100) * time.Millisecond)
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
func (t *Agent) ReportStatus(args *model.AgentReportRequest, reply *model.SimpleRpcResponse) error {
	if args.Hostname == "" {
		reply.Code = 1
		return nil
	}
	//fmt.Println("success receive message",args.String())

	//cache.Agents.Put(args)
	err := SaveAgentReport(args)
	if err != nil {
		return err
	}
	return nil
}

func SaveAgentReport(args *model.AgentReportRequest) error {
	fmt.Println(args.String())
	return nil
}

func (t *Server) Update(metric []*model.MetricValue, reply *model.TransferResponse) error {
	fmt.Println("```````````````wei chaoshi ```````````````````````````````")
	start := time.Now()
	reply.Invalid = 0

	items := []*model.MetaData{}
	for _, v := range metric {
		if v == nil {
			reply.Invalid += 1
			continue
		}

		// 历史遗留问题.
		// 老版本agent上报的metric=kernel.hostname的数据,其取值为string类型,现在已经不支持了;所以,这里硬编码过滤掉
		if v.Metric == "kernel.hostname" {
			reply.Invalid += 1
			continue
		}

		if v.Metric == "" || v.Endpoint == "" {
			reply.Invalid += 1
			continue
		}

		//if v.Type != g.COUNTER && v.Type != g.GAUGE && v.Type != g.DERIVE {
		//	reply.Invalid += 1
		//	continue
		//}

		if v.Value == "" {
			reply.Invalid += 1
			continue
		}

		if v.Step <= 0 {
			reply.Invalid += 1
			continue
		}

		if len(v.Metric)+len(v.Tags) > 510 {
			reply.Invalid += 1
			continue
		}

		// TODO 呵呵,这里需要再优雅一点
		now := start.Unix()
		if v.Timestamp <= 0 || v.Timestamp > now*2 {
			v.Timestamp = now
		}

		fv := &model.MetaData{
			Metric:      v.Metric,
			Endpoint:    v.Endpoint,
			Timestamp:   v.Timestamp,
			Step:        v.Step,
			CounterType: v.Type,
			Tags:        DictedTagstring(v.Tags), //TODO tags键值对的个数,要做一下限制
		}

		valid := true
		var vv float64
		var err error

		switch cv := v.Value.(type) {
		case string:
			vv, err = strconv.ParseFloat(cv, 64)
			if err != nil {
				valid = false
			}
		case float64:
			vv = cv
		case int64:
			vv = float64(cv)
		default:
			valid = false
		}
		if !valid {
			reply.Invalid += 1
			continue
		}
		fv.Value = vv
		items = append(items, fv)

	}

	for _, v := range items {
		influx.SaveInformation(v)
	}

	reply.Message = "ok"
	reply.Total = len(metric)
	reply.Latency = (time.Now().UnixNano() - start.UnixNano()) / 1000000

	return nil
}

package db

import (
	"log"
	"time"

	"../g"
	"../model"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
)

var (
	myDB          string
	username      string
	password      string
	myMeasurement string
	addr          string
)
var clientObject client.Client

func create() client.Client {

	cfg := g.Config()
	myDB = cfg.Database
	username = cfg.Username
	password = cfg.Password
	myMeasurement = cfg.MyMeasurement
	addr = cfg.InfluxAddr

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
		fmt.Println("connect error")
	}
	defer c.Close()

	return c
}

//SaveMetric
func SaveInformation(md *model.MetaData) error {
	if clientObject == nil {
		clientObject = create()
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  myDB,
		Precision: "us",
	})
	if err != nil {
		log.Fatal(err)
		fmt.Println("batch point error")
	}
	tags := map[string]string{
		"Endpoint":    md.Endpoint,
		"CounterType": md.CounterType,
		//"Metric":      md.Metric,
		//"Tags":        md.Tags,
		//"Timestamp":   md.Timestamp,
	}

	fields := map[string]interface{}{
		"Step":  md.Step,
		"Value": md.Value,
	}

	pt, err := client.NewPoint(
		md.Metric,
		tags,
		fields,
		time.Unix(md.Timestamp, 0),
	)

	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	err = clientObject.Write(bp)
	return err
}

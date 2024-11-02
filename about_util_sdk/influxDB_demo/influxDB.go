package main

import (
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"golang-first-step/about_util_sdk/gopsutil_demo"
	"log"
	"time"
)

var cli client.Client

func initConnInflux() (err error) {
	cli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://127.0.0.1:8086",
		Username: "admin",
		Password: "",
	})

	return
}

// query
func queryDB(cli client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: "test",
	}
	if response, err := cli.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

// insert
func writesPoints(percent float64) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "monitor",
		Precision: "s", //精度，默认ns
	})
	if err != nil {
		log.Fatal(err)
	}
	tags := map[string]string{"cpu": "cpu0"}
	fields := map[string]interface{}{
		"cpu percent": percent,
	}

	pt, err := client.NewPoint("cpu_percent", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)
	err = cli.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert success")
}

func main() {
	err := initConnInflux()
	if err != nil {
		fmt.Println(fmt.Sprintf("connect influxdb err:%v", err))
		return
	}

	tick := time.Tick(time.Second)

	for {
		select {
		case <-tick:
			writesPoints(gopsutil_demo.GetCpuInfo())
		}
	}

	// 获取10条数据并展示
	//qs := fmt.Sprintf("SELECT * FROM %s LIMIT %d", "cpu_usage", 10)
	//res, err := queryDB(conn, qs)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, row := range res[0].Series[0].Values {
	//	for j, value := range row {
	//		log.Printf("j:%d value:%v\n", j, value)
	//	}
	//}
}

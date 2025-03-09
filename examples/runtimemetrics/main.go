package main

import (
	"fmt"
	"runtime/metrics"
)

func main() {
	// 获取所有可用指标的描述
	descs := metrics.All()
	for _, desc := range descs {
		fmt.Println(desc.Name) // 这将打印出所有可用指标的名称，包括 /sched/latencies:seconds
	}
}

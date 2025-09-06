package main

import (
	"log"
	"system-monitor/monitor"
	"system-monitor/ui"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 初始化监控器
	sysMonitor := monitor.NewSystemMonitor()
	if err := sysMonitor.Initialize(); err != nil {
		log.Fatalf("Failed to initialize system monitor: %v", err)
	}
	defer sysMonitor.Close()

	// 启动数据收集
	go sysMonitor.StartMonitoring()

	// 启动UI
	if err := ui.StartUI(sysMonitor); err != nil {
		log.Fatalf("Failed to start UI: %v", err)
	}
}

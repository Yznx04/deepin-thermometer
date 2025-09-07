package main

import (
	"deepin-thermometer/internal/monitor"
	"deepin-thermometer/internal/ui"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 初始化监控器
	sysMonitor := monitor.NewSystemMonitor()
	if err := sysMonitor.Initialize(); err != nil {
		log.Printf("Warning: Failed to initialize system monitor: %v", err)
		// 继续运行，使用模拟数据
	}
	defer sysMonitor.Close()

	// 启动数据收集
	go sysMonitor.StartMonitoring()

	// 启动UI
	if err := ui.StartUI(sysMonitor); err != nil {
		log.Fatalf("Failed to start UI: %v", err)
	}
}

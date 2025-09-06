package monitor

import (
	"fmt"
	"sync"
	"time"
)

type SystemData struct {
	CPUTemp     float64
	GPUTemp     float64
	CPUUsage    float64
	MemoryUsage float64
	PowerUsage  float64
	Timestamp   time.Time
}

type SystemMonitor struct {
	data     SystemData
	mu       sync.RWMutex
	running  bool
	stopChan chan struct{}
	sensors  *SensorReader
}

func NewSystemMonitor() *SystemMonitor {
	return &SystemMonitor{
		data:     SystemData{},
		stopChan: make(chan struct{}),
		sensors:  NewSensorReader(),
	}
}

func (sm *SystemMonitor) Initialize() error {
	return sm.sensors.Initialize()
}

func (sm *SystemMonitor) GetData() SystemData {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.data
}

func (sm *SystemMonitor) StartMonitoring() {
	sm.running = true
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if !sm.running {
				return
			}
			sm.updateData()
		case <-sm.stopChan:
			return
		}
	}
}

func (sm *SystemMonitor) StopMonitoring() {
	sm.running = false
	close(sm.stopChan)
}

func (sm *SystemMonitor) Close() {
	sm.StopMonitoring()
	if sm.sensors != nil {
		sm.sensors.Close()
	}
}

func (sm *SystemMonitor) updateData() {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// 更新温度数据
	cpuTemp, gpuTemp := sm.sensors.GetTemperatures()
	sm.data.CPUTemp = cpuTemp
	sm.data.GPUTemp = gpuTemp

	// 更新CPU和内存使用率
	cpuUsage, memUsage := sm.getSystemUsage()
	sm.data.CPUUsage = cpuUsage
	sm.data.MemoryUsage = memUsage

	// 估算功耗（简化计算）
	sm.data.PowerUsage = sm.estimatePowerUsage(cpuUsage, memUsage)

	sm.data.Timestamp = time.Now()
}

func (sm *SystemMonitor) getSystemUsage() (cpu float64, memory float64) {
	// 这里使用简化的系统使用率计算
	// 实际应用中可以读取 /proc/stat 和 /proc/meminfo
	cpu = 25.0 + (50.0 * float64(time.Now().Unix()%100) / 100.0)    // 修正：添加float64转换
	memory = 45.0 + (30.0 * float64(time.Now().Unix()%100) / 100.0) // 修正：添加float64转换
	return
}

func (sm *SystemMonitor) estimatePowerUsage(cpuUsage, memUsage float64) float64 {
	// 简化的功耗估算公式
	// 实际功耗计算需要更复杂的硬件特定算法
	basePower := 15.0 // 基础功耗
	cpuPower := cpuUsage * 0.3
	memPower := memUsage * 0.1
	return basePower + cpuPower + memPower
}

func (sm *SystemMonitor) FormatData() string {
	data := sm.GetData()
	return fmt.Sprintf(
		"CPU: %.1f°C (%.1f%%)\n"+
			"GPU: %.1f°C\n"+
			"MEM: %.1f%%\n"+
			"POWER: %.1fW",
		data.CPUTemp, data.CPUUsage,
		data.GPUTemp,
		data.MemoryUsage,
		data.PowerUsage,
	)
}

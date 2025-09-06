package ui

import (
	"system-monitor/monitor"
	"time"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

type MainWindow struct {
	window       *ui.Window
	label        *ui.Label
	monitor      *monitor.SystemMonitor
	updateTicker *time.Ticker
}

var mainWindow *MainWindow

func StartUI(sysMonitor *monitor.SystemMonitor) error {
	err := ui.Main(func() {
		mainWindow = &MainWindow{
			monitor: sysMonitor,
		}
		mainWindow.createWindow()
		mainWindow.startUpdates()
	})

	return err
}

func (mw *MainWindow) createWindow() {
	// 创建主窗口
	mw.window = ui.NewWindow("System Monitor", 200, 120, true) // 改为有边框以便于操作

	// 创建标签显示系统信息
	mw.label = ui.NewLabel("Initializing...")

	// 设置窗口内容
	box := ui.NewVerticalBox()
	box.SetPadded(true)
	box.Append(mw.label, true)
	mw.window.SetChild(box)

	// 显示窗口
	mw.window.Show()

	// 绑定事件
	mw.bindEvents()
}

func (mw *MainWindow) bindEvents() {
	// 窗口关闭事件
	mw.window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
}

func (mw *MainWindow) startUpdates() {
	// 使用ticker定期更新界面
	mw.updateTicker = time.NewTicker(1 * time.Second)

	go func() {
		for range mw.updateTicker.C {
			ui.QueueMain(func() {
				mw.updateDisplay()
			})
		}
	}()
}

func (mw *MainWindow) updateDisplay() {
	if mw.monitor != nil && mw.label != nil {
		displayText := mw.monitor.FormatData()
		mw.label.SetText(displayText)
	}
}

func (mw *MainWindow) stopUpdates() {
	if mw.updateTicker != nil {
		mw.updateTicker.Stop()
	}
}

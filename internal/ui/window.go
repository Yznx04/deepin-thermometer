package ui

import (
	"image/color"
	"system-monitor/internal/monitor"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	windowWidth  = 200
	windowHeight = 120
)

type Game struct {
	monitor     *monitor.SystemMonitor
	dragging    bool
	dragOffsetX int
	dragOffsetY int
	windowX     int
	windowY     int
	lastClick   time.Time
}

func StartUI(sysMonitor *monitor.SystemMonitor) error {
	// 启动监控
	go sysMonitor.StartMonitoring()

	// 创建游戏实例
	game := &Game{
		monitor: sysMonitor,
		windowX: 1000,
		windowY: 50,
	}

	// 设置窗口属性
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("System Monitor")
	ebiten.SetWindowDecorated(false) // 无边框
	ebiten.SetWindowFloating(true)   // 置顶

	// 运行游戏循环
	if err := ebiten.RunGame(game); err != nil {
		return err
	}

	return nil
}

func (g *Game) Update() error {
	// 获取鼠标位置
	mouseX, mouseY := ebiten.CursorPosition()

	// 检查鼠标左键按下
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// 检查是否在窗口内点击
		if mouseX >= 0 && mouseX < windowWidth && mouseY >= 0 && mouseY < windowHeight {
			// 检测双击
			now := time.Now()
			if now.Sub(g.lastClick) < 300*time.Millisecond {
				// 双击关闭程序
				return ebiten.Termination
			}
			g.lastClick = now

			g.dragging = true
			g.dragOffsetX = mouseX
			g.dragOffsetY = mouseY
		}
	}

	// 检查鼠标左键释放
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.dragging = false
	}

	// 检查右键关闭
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		return ebiten.Termination
	}

	// 拖动窗口
	if g.dragging {
		currentX, currentY := ebiten.WindowPosition()
		deltaX := mouseX - g.dragOffsetX
		deltaY := mouseY - g.dragOffsetY
		ebiten.SetWindowPosition(currentX+deltaX, currentY+deltaY)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 绘制背景
	screen.Fill(color.RGBA{0, 0, 0, 200})

	// 绘制边框
	for i := 0; i < windowWidth; i++ {
		screen.Set(i, 0, color.RGBA{100, 100, 100, 255})
		screen.Set(i, windowHeight-1, color.RGBA{100, 100, 100, 255})
	}
	for i := 0; i < windowHeight; i++ {
		screen.Set(0, i, color.RGBA{100, 100, 100, 255})
		screen.Set(windowWidth-1, i, color.RGBA{100, 100, 100, 255})
	}

	// 绘制系统信息
	var text string
	if g.monitor != nil {
		text = g.monitor.FormatData()
	} else {
		text = "Initializing..."
	}

	// 绘制文本
	ebitenutil.DebugPrintAt(screen, text, 10, 10)

	// 绘制提示信息
	ebitenutil.DebugPrintAt(screen, "L:Drag R:Close", 10, windowHeight-20)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

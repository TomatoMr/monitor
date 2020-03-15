package main

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	//初始化图形数据
	bc := widgets.NewBarChart()
	bc.Title = "monitor"
	bc.SetRect(5, 5, 100, 25)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	//启动http反向代理服务
	go Proxy()
	//更新访问数据
	go func() {
		for {
			select {
			case label := <-ch:
				_, has := labels[label]
				if has {
					labels[label]++
				} else {
					labels[label] = 1
				}
			}
		}
	}()

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents: //退出事件
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker: //定时事件
			l := []string{}
			data := []float64{}
			for k, v := range labels {
				l = append(l, k)
				data = append(data, v)
			}
			if len(l) <= 0 {
				l = []string{"/"}
				data = []float64{1}
			}
			bc.Labels = l
			bc.Data = data
			ui.Render(bc)
		}
	}
}

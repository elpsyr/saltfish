package app

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/systray"
	x "fyne.io/x/fyne/widget"
	"github.com/elpsyr/saltfish/internal/job"
	"log"
	"net/url"
	"time"
)

//go:generate fyne bundle -o bundled.go ../../images/xianyu.svg
//go:generate fyne bundle -o bundled.go -append ../../images/xianyu.ico
//go:generate fyne bundle -o bundled.go -append ../../images/moyu.gif
func Run() {
	myApp := app.New()
	w := myApp.NewWindow("salt fish @elpsyr")

	w.SetIcon(resourceXianyuIco)
	w.Resize(fyne.Size{
		Width: 300,
	})

	// 点击关闭进行隐藏
	w.SetCloseIntercept(func() {
		w.Hide()
	})

	manager := job.NewManager()
	//manager.ResizeWindow()

	parse, err := url.Parse("https://github.com/elpsyr/saltfish")
	if err != nil {
		fmt.Println(err)
	}
	hyperlink := widget.NewHyperlink("How to use", parse)

	str := binding.NewString()
	str.Set(fmt.Sprintf("Rewards : %d Fishing : %d", manager.GetCountReward(), manager.GetCountFish()))
	workLabel := widget.NewLabelWithData(str)

	timeLabel := widget.NewLabel("Run Time : 00:00:00")

	// gif
	//gif, err := x.NewAnimatedGif(storage.NewFileURI("./images/moyu.gif"))
	gif, err := x.NewAnimatedGifFromResource(resourceMoyuGif)
	gif.Show()
	gif.SetMinSize(fyne.Size{
		Width:  20,
		Height: 20,
	})
	gif.Start()

	slider := widget.NewSlider(1, 255)
	alpha := binding.NewFloat()
	alpha.Set(255)
	slider.Bind(alpha)
	slider.OnChanged = func(f float64) {
		manager.AlphaWindow(int(f))
	}

	box := container.NewVBox(
		//gif,
		widget.NewButton("hide", func() {
			manager.HideMode()
		}),
		widget.NewButton("show", func() {
			manager.ShowMode()
		}),
		widget.NewButton("reward", func() {
			go manager.SetCallBack(func() {
				str.Set(fmt.Sprintf("Rewards : %d Fishing : %d", manager.GetCountReward(), manager.GetCountFish()))
			}).GetReward()
		}),
		widget.NewButton("fishing", func() {
			go manager.SetCallBack(func() {
				str.Set(fmt.Sprintf("Rewards : %d Fishing : %d", manager.GetCountReward(), manager.GetCountFish()))
			}).GetFish()
		}),

		container.New(layout.NewStackLayout(), slider),
		container.New(layout.NewHBoxLayout(), layout.NewSpacer(), workLabel, layout.NewSpacer()),
		container.New(layout.NewHBoxLayout(), layout.NewSpacer(), timeLabel, layout.NewSpacer()),
		container.New(layout.NewHBoxLayout(), layout.NewSpacer(), gif, hyperlink, layout.NewSpacer()),
	)
	w.SetContent(box)

	menu := fyne.NewMenu("MyApp",
		fyne.NewMenuItem("open", func() {
			log.Println("Tapped show")
			w.Show()
		}),
		fyne.NewMenuItem("show", func() {
			manager.ShowMode()
		}),
		fyne.NewMenuItem("hide", func() {
			manager.HideMode()
		}),
		fyne.NewMenuItem("reward", func() {
			go manager.GetReward()
		}),
	)

	if desk, ok := myApp.(desktop.App); ok {
		desk.SetSystemTrayIcon(resourceXianyuIco)
		desk.SetSystemTrayMenu(menu)
	}

	drv := fyne.CurrentApp().Driver()
	if drv, ok := drv.(desktop.Driver); ok {
		fmt.Println(ok)

		_w := drv.CreateSplashWindow()
		spGif, err := x.NewAnimatedGifFromResource(resourceMoyuGif)
		if err != nil {
		}
		spGif.Show()
		spGif.SetMinSize(fyne.Size{
			Width:  150,
			Height: 150,
		})
		spGif.Start()

		_w.SetContent(spGif)

		_w.Show()
		go func() {
			time.Sleep(time.Second * 3)
			_w.Close()
			spGif.Stop()
			w.Show()
		}()
	}

	go updateTimeLabel(timeLabel)
	GetReward2Hour(manager, str) // 注册
	GetFish8Hour(manager, str)   // 注册

	//systray.Run(onReady, onExit)
	//w.ShowAndRun()
	myApp.Run()
}

func onReady() {

	systray.SetIcon(resourceXianyuIco.Content())
	systray.SetTitle("salt fish")
	systray.SetTooltip("salt fish")

	mOpen := systray.AddMenuItem("open", "")
	mQuit := systray.AddMenuItem("quit", "")

	go func() {
		for {
			select {
			case <-mOpen.ClickedCh:
				fmt.Println("open")
			case <-mQuit.ClickedCh:
				fmt.Println("quit")
			}
		}
	}()
}

func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		drv := fyne.CurrentApp().Driver()
		if drv, ok := drv.(desktop.Driver); ok {
			fmt.Println(ok)

			w := drv.CreateSplashWindow()
			w.SetContent(widget.NewLabelWithStyle("Hello World!\n\nMake a splash!",
				fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
			//go func() {
			//	time.Sleep(time.Second * 3)
			//	w.Close()
			//}()
			w.Show()
		}
	})
}

func onExit() {
	// clean up here
}

func elapsedTime(startTime time.Time) string {
	elapsedTime := time.Since(startTime)
	hours := int(elapsedTime.Hours())
	minutes := int(elapsedTime.Minutes()) % 60
	seconds := int(elapsedTime.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func updateTimeLabel(label *widget.Label) {
	startTime := time.Now()

	for {
		currentElapsedTime := elapsedTime(startTime)
		time.Sleep(time.Second)
		label.SetText("Run Time : " + currentElapsedTime)
	}
}

func GetReward2Hour(m *job.Manager, str binding.String) {
	// 创建一个每隔4小时触发一次的Ticker
	ticker := time.NewTicker(4 * time.Hour)
	//ticker := time.NewTicker(30 * time.Second)

	// 启动一个goroutine来处理Ticker触发的事件
	go func() {
		for {
			select {
			case <-ticker.C:
				// 在Ticker触发时调用方法A
				go m.SetCallBack(func() {
					str.Set(fmt.Sprintf("Rewards : %d Fishing : %d", m.GetCountReward(), m.GetCountFish()))
				}).GetReward()
			}
		}
	}()

}

func GetFish8Hour(m *job.Manager, str binding.String) {
	// 创建一个每隔8小时触发一次的Ticker
	ticker := time.NewTicker(8*time.Hour + time.Minute)

	// 启动一个goroutine来处理Ticker触发的事件
	go func() {
		for {
			select {
			case <-ticker.C:
				// 在Ticker触发时调用方法A
				go m.SetCallBack(func() {
					str.Set(fmt.Sprintf("Rewards : %d Fishing : %d", m.GetCountReward(), m.GetCountFish()))
				}).GetFish()
			}
		}
	}()

}

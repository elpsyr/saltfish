package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"fyne.io/systray"
	"github.com/elpsyr/saltfish/internal/job"
	"github.com/elpsyr/saltfish/pkg/win"
	"log"
	"os"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("salt fish")
	resource, err := fyne.LoadResourceFromPath("./images/fish.svg")
	if err != nil {
		fmt.Println("LoadResourceFromPath ERROR:", err)
	}
	w.SetIcon(resource)
	w.Resize(fyne.Size{
		Width: 300,
	})

	// 点击关闭进行隐藏
	w.SetCloseIntercept(func() {
		w.Hide()
	})

	manager := job.Manager{}

	//hello := widget.NewLabel("👑")
	w.SetContent(container.NewVBox(
		//hello,
		widget.NewButton("hide", func() {
			//hello.SetText("🦈️")
			hwnd := win.GetHwndByTitle("咸鱼之王")
			win.SetTopWindow(hwnd)
			win.HideWindow(hwnd)
		}),
		widget.NewButton("show", func() {
			//hello.SetText("🐟")
			hwnd := win.GetHwndByTitle("咸鱼之王")
			win.ShowWindow(hwnd)
			win.SetTopWindow(hwnd)
		}),
		widget.NewButton("reward", func() {
			//hello.SetText("🪙")
			hwnd := win.GetHwndByTitle("咸鱼之王")
			go manager.GetReward(hwnd)
		}),
		widget.NewButton("fishing", func() {
			hwnd := win.GetHwndByTitle("咸鱼之王")
			go manager.GetFish(hwnd)
		}),
	))

	menu := fyne.NewMenu("MyApp",
		fyne.NewMenuItem("Show", func() {
			log.Println("Tapped show")
			w.Show()
		}))

	if desk, ok := myApp.(desktop.App); ok {
		resourceIco, err := fyne.LoadResourceFromPath("./images/fish.ico")
		if err != nil {
			fmt.Println("LoadResourceFromPath ERROR:", err)
		}
		desk.SetSystemTrayIcon(resourceIco)
		desk.SetSystemTrayMenu(menu)
	}

	//systray.Run(onReady, onExit)
	w.ShowAndRun()
}

func onReady() {

	// 使用 ioutil.ReadFile 读取图片文件内容
	imageBytes, err := os.ReadFile("./images/fish.ico")
	if err != nil {
		fmt.Println("无法读取图片文件:", err)
		return
	}
	systray.SetIcon(imageBytes)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome超级棒")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Sets the icon of a menu item.
	mQuit.SetIcon(imageBytes)
}

func onExit() {
	// clean up here
}

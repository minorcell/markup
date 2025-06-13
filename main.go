package main

import (
	"markedit/internal/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	// 创建应用实例，并提供一个唯一的ID来支持快捷键等功能
	myApp := app.NewWithID("com.github.markedit")

	// 设置自定义的GitHub主题
	myApp.Settings().SetTheme(ui.NewGitHubTheme())

	// 创建主窗口
	myWindow := myApp.NewWindow("MarkUp")
	myWindow.Resize(fyne.NewSize(1200, 800))

	// 创建主界面控制器
	controller := ui.NewGuiController()

	// 设置窗口内容
	content := controller.BuildUI(myWindow)
	myWindow.SetContent(content)

	// 设置窗口关闭时的回调
	myWindow.SetCloseIntercept(func() {
		controller.OnWindowClose()
		myApp.Quit()
	})

	// 显示窗口并运行应用
	myWindow.ShowAndRun()
}

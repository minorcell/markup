package ui

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"markup/internal/core"
	"markup/internal/markdown"
)

// GuiController 主界面控制器
type GuiController struct {
	window     fyne.Window
	mdRenderer *markdown.Renderer
	appState   *core.AppState

	// UI 组件
	editorEntry *widget.Entry

	// 状态
	isEditing bool // 是否处于编辑模式
}

// NewGuiController 创建新的主控制器
func NewGuiController() *GuiController {
	return &GuiController{
		mdRenderer: markdown.NewRenderer(),
		appState:   core.NewAppState(),
		isEditing:  false,
	}
}

// BuildUI 构建用户界面
func (sc *GuiController) BuildUI(window fyne.Window) fyne.CanvasObject {
	sc.window = window

	if !sc.isEditing {
		// 显示启动界面
		return sc.buildStartupUI()
	} else {
		// 显示编辑界面
		return sc.buildEditorUI()
	}
}

// buildStartupUI 构建启动界面
func (sc *GuiController) buildStartupUI() fyne.CanvasObject {
	// 标题
	title := widget.NewLabel("MarkUp")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	// 新建按钮
	newBtn := widget.NewButton("新建文件", func() {
		sc.createNewFile()
	})
	newBtn.Resize(fyne.NewSize(200, 50))

	// 打开按钮
	openBtn := widget.NewButton("打开文件", func() {
		sc.openFile()
	})
	openBtn.Resize(fyne.NewSize(200, 50))

	// 创建垂直布局
	content := container.NewVBox(
		widget.NewLabel(""), // 空白
		widget.NewLabel(""), // 空白
		title,
		widget.NewLabel(""), // 空白
		container.NewCenter(newBtn),
		widget.NewLabel(""), // 间距
		container.NewCenter(openBtn),
		widget.NewLabel(""), // 空白
		widget.NewLabel(""), // 空白
	)

	return container.NewCenter(content)
}

// buildEditorUI 构建编辑界面
func (sc *GuiController) buildEditorUI() fyne.CanvasObject {
	// 初始化编辑器
	sc.editorEntry = widget.NewMultiLineEntry()
	sc.editorEntry.Wrapping = fyne.TextWrapWord
	sc.editorEntry.SetPlaceHolder("在此输入 Markdown 内容...")

	// 设置文本变化事件
	sc.editorEntry.OnChanged = func(content string) {
		sc.appState.SetCurrentContent(content)
	}

	// 创建工具栏
	toolbar := sc.createEditorToolbar()

	// 创建主布局
	return container.NewBorder(
		toolbar,                             // top
		nil,                                 // bottom
		nil,                                 // left
		nil,                                 // right
		container.NewScroll(sc.editorEntry), // center
	)
}

// createEditorToolbar 创建编辑器工具栏
func (sc *GuiController) createEditorToolbar() *fyne.Container {
	// 保存按钮
	saveBtn := widget.NewButton("保存", func() {
		sc.saveFile()
	})

	return container.NewHBox(
		saveBtn,
	)
}

// createNewFile 创建新文件
func (sc *GuiController) createNewFile() {
	sc.isEditing = true
	sc.appState.SetCurrentFile("")
	sc.appState.SetCurrentContent("# 新文档\n\n开始编写您的内容...")
	sc.appState.SetOriginalContent("")

	// 重新构建UI
	sc.window.SetContent(sc.BuildUI(sc.window))

	// 设置编辑器内容
	if sc.editorEntry != nil {
		sc.editorEntry.SetText(sc.appState.GetCurrentContent())
	}
}

// openFile 打开文件
func (sc *GuiController) openFile() {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil || reader == nil {
			return
		}
		defer reader.Close()

		// 检查文件扩展名
		filename := reader.URI().Name()
		if !strings.HasSuffix(strings.ToLower(filename), ".md") &&
			!strings.HasSuffix(strings.ToLower(filename), ".markdown") {
			dialog.ShowInformation("错误", "请选择 Markdown 文件（.md 或 .markdown）", sc.window)
			return
		}

		// 读取文件内容
		content, err := sc.appState.LoadFile(reader.URI().Path())
		if err != nil {
			dialog.ShowError(err, sc.window)
			return
		}

		// 设置状态
		sc.isEditing = true
		sc.appState.SetCurrentFile(reader.URI().Path())
		sc.appState.SetCurrentContent(content)
		sc.appState.SetOriginalContent(content)

		// 重新构建UI
		sc.window.SetContent(sc.BuildUI(sc.window))

		// 设置编辑器内容
		if sc.editorEntry != nil {
			sc.editorEntry.SetText(content)
		}
	}, sc.window)
}

// saveFile 保存文件
func (sc *GuiController) saveFile() {
	currentFile := sc.appState.GetCurrentFile()
	content := sc.appState.GetCurrentContent()

	if currentFile == "" {
		// 另存为
		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil || writer == nil {
				return
			}
			defer writer.Close()

			// 确保文件扩展名
			filename := writer.URI().Name()
			if !strings.HasSuffix(strings.ToLower(filename), ".md") &&
				!strings.HasSuffix(strings.ToLower(filename), ".markdown") {
				filename += ".md"
			}

			// 写入文件
			_, err = writer.Write([]byte(content))
			if err != nil {
				dialog.ShowError(err, sc.window)
				return
			}

			// 更新状态
			sc.appState.SetCurrentFile(writer.URI().Path())
			sc.appState.SetOriginalContent(content)

			dialog.ShowInformation("保存成功", "文件已保存", sc.window)
		}, sc.window)
	} else {
		// 直接保存
		err := sc.appState.SaveFile(currentFile, content)
		if err != nil {
			dialog.ShowError(err, sc.window)
			return
		}

		sc.appState.SetOriginalContent(content)
		dialog.ShowInformation("保存成功", "文件已保存", sc.window)
	}
}

// OnWindowClose 窗口关闭时的处理
func (sc *GuiController) OnWindowClose() {
	// 自动保存
	if sc.isEditing && sc.appState.HasUnsavedChanges() {
		currentFile := sc.appState.GetCurrentFile()
		content := sc.appState.GetCurrentContent()

		if currentFile != "" {
			// 如果有文件路径，直接保存
			sc.appState.SaveFile(currentFile, content)
		}
		// 如果是新文件且没有路径，则不自动保存（避免弹出对话框）
	}
}

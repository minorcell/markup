package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// GitHubTheme 结构体，用于实现自定义主题
type GitHubTheme struct{}

// 断言 GitHubTheme 实现了 fyne.Theme 接口
var _ fyne.Theme = (*GitHubTheme)(nil)

// NewGitHubTheme 返回一个新的 GitHubTheme 实例
func NewGitHubTheme() *GitHubTheme {
	return &GitHubTheme{}
}

// Color 返回指定名称和变体的颜色
func (t *GitHubTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	// 忽略亮色/暗色变体，始终使用我们定义的颜色
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0xe6, G: 0xea, B: 0xed, A: 0xff} // GitHub 更深的背景灰
	case theme.ColorNameButton:
		return color.NRGBA{R: 0xd0, G: 0xd7, B: 0xde, A: 0xff} // 按钮灰色
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 0x95, G: 0x9d, B: 0xa5, A: 0xff}
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 0xf6, G: 0xf8, B: 0xfa, A: 0x80}
	case theme.ColorNameError:
		return color.NRGBA{R: 0xcb, G: 0x24, B: 0x31, A: 0xff}
	case theme.ColorNameFocus:
		return color.NRGBA{R: 0x03, G: 0x66, B: 0xd6, A: 0xff} // GitHub 链接/焦点蓝
	case theme.ColorNameForeground:
		return color.NRGBA{R: 0x24, G: 0x29, B: 0x2e, A: 0xff} // GitHub 主要文本黑
	case theme.ColorNameHover:
		return color.NRGBA{R: 0xc8, G: 0xd3, B: 0xde, A: 0xff} // hover效果
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff} // 输入框白色背景
	case theme.ColorNameInputBorder:
		return color.NRGBA{R: 0xd1, G: 0xd5, B: 0xda, A: 0xff} // 输入框边框
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0x6a, G: 0x73, B: 0x7d, A: 0xff}
	case theme.ColorNamePressed:
		return color.NRGBA{R: 0xb0, G: 0xb7, B: 0xbe, A: 0xff} // 按下效果
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 0x03, G: 0x66, B: 0xd6, A: 0xff}
	case theme.ColorNameScrollBar:
		return color.NRGBA{R: 0xd1, G: 0xd5, B: 0xda, A: 0xff}
	case theme.ColorNameSelection:
		return color.NRGBA{R: 0x03, G: 0x66, B: 0xd6, A: 0x40} // 选中文本背景
	case theme.ColorNameSeparator:
		return color.NRGBA{R: 0xe1, G: 0xe4, B: 0xe8, A: 0xff} // 分隔线
	case theme.ColorNameShadow:
		return color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x20}
	case theme.ColorNameSuccess:
		return color.NRGBA{R: 0x2e, G: 0xa4, B: 0x4f, A: 0xff} // 成功色
	case theme.ColorNameWarning:
		return color.NRGBA{R: 0xd1, G: 0x9a, B: 0x66, A: 0xff} // 警告色
	case theme.ColorNameMenuBackground:
		return color.NRGBA{R: 0xd0, G: 0xd7, B: 0xde, A: 0xff} // 菜单背景
	case theme.ColorNameOverlayBackground:
		return color.NRGBA{R: 0xd0, G: 0xd7, B: 0xde, A: 0xff} // 覆盖层背景
	default:
		// 对于未指定的颜色，回退到默认的亮色主题
		return theme.DefaultTheme().Color(name, theme.VariantLight)
	}
}

// Font 返回指定样式和大小的字体
func (t *GitHubTheme) Font(style fyne.TextStyle) fyne.Resource {
	// 暂不覆盖字体，使用Fyne的默认字体
	return theme.DefaultTheme().Font(style)
}

// Icon 返回指定名称的图标
func (t *GitHubTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	// 暂不覆盖图标，使用Fyne的默认图标
	return theme.DefaultTheme().Icon(name)
}

// Size 返回指定名称的大小
func (t *GitHubTheme) Size(name fyne.ThemeSizeName) float32 {
	// 暂不覆盖尺寸，使用Fyne的默认尺寸
	return theme.DefaultTheme().Size(name)
}

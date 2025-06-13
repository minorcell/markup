# MarkUp 编辑器

一个基于 Fyne 框架开发的轻量级、本地运行的 Markdown 编辑器。专注于提供简洁的界面和高效的编辑体验。

## 安装和运行

### 前置要求
- Go 1.21 或更高版本
- 支持 GUI 的操作系统（Windows、macOS、Linux）

### 编译和运行
```bash
# 克隆项目
git clone https://github.com/minorell/markup
cd markup

# 下载依赖
go mod tidy

# 编译并运行
go run main.go
```

### 构建可执行文件
```bash
# 构建当前平台的可执行文件
go build -o markup .

# 跨平台构建示例
# Windows
GOOS=windows GOARCH=amd64 go build -o markup.exe .

# macOS
GOOS=darwin GOARCH=amd64 go build -o markup-mac .

# Linux
GOOS=linux GOARCH=amd64 go build -o markup-linux .
```

### 支持的 Markdown 语法
- **标题**：`# H1`, `## H2`, `### H3` 等
- **文本格式**：`**粗体**`, `*斜体*`, `~~删除线~~`
- **列表**：有序和无序列表
- **链接**：`[文本](URL)`
- **代码**：行内代码和代码块
- **引用**：`> 引用内容`
- **表格**：支持标准 Markdown 表格语法

## 技术架构

### 框架和库
- **UI 框架**：[Fyne v2](https://fyne.io/) - 跨平台 Go GUI 框架
- **Markdown 解析**：[gomarkdown](https://github.com/gomarkdown/markdown) - Go Markdown 解析器

### 架构设计
- **模块化设计**：清晰的包结构和职责分离
- **状态管理**：集中式的应用状态管理
- **事件驱动**：基于用户交互的事件处理
- **响应式更新**：实时的内容同步和界面更新

## 许可证

本项目采用 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。

## 致谢

感谢以下开源项目和社区：
- [Fyne](https://fyne.io/) - 优秀的 Go GUI 框架
- [gomarkdown](https://github.com/gomarkdown/markdown) - 强大的 Markdown 解析器
- Go 语言社区的支持和贡献

---

**MarkUp** - 让 Markdown 编辑更简单！

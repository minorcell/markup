# MarkUp - Markdown 编辑器

一个基于 Fyne 框架开发的轻量级、本地运行的 Markdown 编辑器。专注于提供简洁的界面和高效的编辑体验。

## 特性

### 🚀 核心功能
- **实时编辑**：流畅的 Markdown 文本编辑体验
- **即时预览**：实时渲染 Markdown 内容
- **模式切换**：编辑和预览模式灵活切换
- **文件管理**：支持文件的新建、保存和导出

### 🎨 用户界面
- **简洁设计**：专注于内容创作的界面设计
- **响应式布局**：自适应窗口大小调整
- **工具栏**：便捷的功能按钮访问
- **快捷键支持**：提高操作效率

### 📄 文件操作
- **新建文件**：快速创建新的 Markdown 文档
- **保存文件**：将内容保存为 .md 文件
- **HTML 导出**：将 Markdown 内容导出为格式化的 HTML 文件
- **文件夹浏览**：（计划功能）浏览和管理文件目录

## 安装和运行

### 前置要求
- Go 1.21 或更高版本
- 支持 GUI 的操作系统（Windows、macOS、Linux）

### 编译和运行
```bash
# 克隆项目
git clone <repository-url>
cd mark

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

## 使用说明

### 基本操作
1. **新建文档**：点击 "新建文件" 按钮开始编写
2. **编辑内容**：在编辑器区域输入 Markdown 内容
3. **预览效果**：点击 "预览" 按钮查看渲染结果
4. **保存文件**：点击 "保存" 按钮将内容保存为文件
5. **导出HTML**：点击 "导出HTML" 按钮生成网页文件

### 键盘快捷键
| 快捷键 | 功能 |
|--------|------|
| `Ctrl+E` | 切换到编辑模式 |
| `Ctrl+P` | 切换到预览模式 |
| `Ctrl+S` | 保存当前文件 |
| `Ctrl+O` | 打开文件夹 |

### 支持的 Markdown 语法
- **标题**：`# H1`, `## H2`, `### H3` 等
- **文本格式**：`**粗体**`, `*斜体*`, `~~删除线~~`
- **列表**：有序和无序列表
- **链接**：`[文本](URL)`
- **代码**：行内代码和代码块
- **引用**：`> 引用内容`
- **表格**：支持标准 Markdown 表格语法

## 项目结构

```
markup/
├── main.go                 # 应用入口点
├── go.mod                  # Go 模块定义
├── example.md              # 示例 Markdown 文件
├── README.md               # 项目说明文档
├── docs/
│   └── product_design.md   # 产品设计文档
└── internal/               # 内部包
    ├── core/               # 核心逻辑
    │   └── state.go        # 应用状态管理
    ├── filemanager/        # 文件管理
    │   └── manager.go      # 文件操作功能
    ├── markdown/           # Markdown 处理
    │   └── renderer.go     # Markdown 渲染器
    └── ui/                 # 用户界面
        ├── controller.go   # 主界面控制器
        └── simple.go       # 简化界面控制器
```

## 技术架构

### 框架和库
- **UI 框架**：[Fyne v2](https://fyne.io/) - 跨平台 Go GUI 框架
- **Markdown 解析**：[gomarkdown](https://github.com/gomarkdown/markdown) - Go Markdown 解析器
- **HTML 清理**：[bluemonday](https://github.com/microcosm-cc/bluemonday) - HTML 安全处理

### 架构设计
- **模块化设计**：清晰的包结构和职责分离
- **状态管理**：集中式的应用状态管理
- **事件驱动**：基于用户交互的事件处理
- **响应式更新**：实时的内容同步和界面更新

## 开发路线图

### 当前版本 (v1.0)
- [x] 基本的 Markdown 编辑功能
- [x] 实时预览
- [x] 文件保存和导出
- [x] 简洁的用户界面

### 计划功能 (v2.0)
- [ ] 文件树浏览器
- [ ] 文档大纲导航
- [ ] 语法高亮
- [ ] 主题支持
- [ ] 文件搜索功能

### 未来展望
- [ ] 插件系统
- [ ] Git 集成
- [ ] 多标签页支持
- [ ] 协作编辑功能

## 贡献指南

欢迎贡献代码、报告问题或提出改进建议！

### 开发环境设置
1. 确保安装了 Go 1.21+
2. Fork 本项目
3. 创建功能分支
4. 提交更改
5. 创建 Pull Request

### 代码规范
- 遵循 Go 代码风格指南
- 添加适当的注释（使用中文）
- 编写单元测试
- 保持代码简洁和可读性

## 许可证

本项目采用 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。

## 致谢

感谢以下开源项目和社区：
- [Fyne](https://fyne.io/) - 优秀的 Go GUI 框架
- [gomarkdown](https://github.com/gomarkdown/markdown) - 强大的 Markdown 解析器
- Go 语言社区的支持和贡献

---

**MarkUp** - 让 Markdown 编辑更简单！

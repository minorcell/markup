package markdown

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"

	"markup/internal/core"
)

// Renderer Markdown渲染器
type Renderer struct {
	parser    *parser.Parser     // Markdown解析器
	htmlFlags html.Flags         // HTML渲染标志
	policy    *bluemonday.Policy // HTML清理策略
}

// NewRenderer 创建新的Markdown渲染器
func NewRenderer() *Renderer {
	// 配置Markdown解析器
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock

	p := parser.NewWithExtensions(extensions)

	// 配置HTML渲染选项
	htmlFlags := html.CommonFlags | html.HrefTargetBlank

	// 创建HTML清理策略
	policy := bluemonday.UGCPolicy()

	return &Renderer{
		parser:    p,
		htmlFlags: htmlFlags,
		policy:    policy,
	}
}

// RenderToHTML 将Markdown内容渲染为HTML
func (r *Renderer) RenderToHTML(mdContent string) string {
	// 解析Markdown
	doc := r.parser.Parse([]byte(mdContent))

	// 创建HTML渲染器
	renderer := html.NewRenderer(html.RendererOptions{Flags: r.htmlFlags})

	// 渲染为HTML
	htmlBytes := markdown.Render(doc, renderer)

	// 清理HTML（防止XSS攻击）
	safeHTML := r.policy.SanitizeBytes(htmlBytes)

	return string(safeHTML)
}

// RenderToHTMLWithTemplate 将Markdown内容渲染为带模板的完整HTML页面
func (r *Renderer) RenderToHTMLWithTemplate(mdContent, title string) string {
	// 获取HTML内容
	htmlContent := r.RenderToHTML(mdContent)

	// HTML模板
	tmpl := `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 800px;
            margin: 0 auto;
            padding: 2rem;
            background-color: #fff;
        }
        h1, h2, h3, h4, h5, h6 {
            margin-top: 2rem;
            margin-bottom: 1rem;
            color: #2c3e50;
        }
        h1 { border-bottom: 2px solid #eaecef; padding-bottom: 0.3rem; }
        h2 { border-bottom: 1px solid #eaecef; padding-bottom: 0.3rem; }
        p { margin-bottom: 1rem; }
        pre {
            background: #f6f8fa;
            border-radius: 6px;
            padding: 16px;
            overflow: auto;
            line-height: 1.45;
        }
        code {
            background: #f6f8fa;
            padding: 0.2em 0.4em;
            border-radius: 3px;
            font-size: 85%;
        }
        blockquote {
            border-left: 4px solid #dfe2e5;
            padding-left: 1rem;
            color: #6a737d;
            margin: 1rem 0;
        }
        table {
            border-collapse: collapse;
            width: 100%;
            margin: 1rem 0;
        }
        th, td {
            border: 1px solid #dfe2e5;
            padding: 0.6rem 1rem;
            text-align: left;
        }
        th {
            background-color: #f6f8fa;
            font-weight: 600;
        }
        a {
            color: #0366d6;
            text-decoration: none;
        }
        a:hover {
            text-decoration: underline;
        }
        img {
            max-width: 100%;
            height: auto;
        }
        .toc {
            background: #f8f9fa;
            border: 1px solid #e1e4e8;
            border-radius: 6px;
            padding: 1rem;
            margin: 1rem 0;
        }
        .toc ul {
            list-style-type: none;
            padding-left: 1rem;
        }
        .toc > ul {
            padding-left: 0;
        }
        .toc a {
            color: #586069;
        }
    </style>
</head>
<body>
    {{.Content}}
</body>
</html>`

	// 创建模板
	t, err := template.New("html").Parse(tmpl)
	if err != nil {
		return htmlContent // 如果模板解析失败，返回原始HTML
	}

	// 准备模板数据
	data := struct {
		Title   string
		Content template.HTML
	}{
		Title:   title,
		Content: template.HTML(htmlContent),
	}

	// 渲染模板
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return htmlContent // 如果模板执行失败，返回原始HTML
	}

	return buf.String()
}

// RenderToRichText 将Markdown内容渲染为Fyne的RichText格式
func (r *Renderer) RenderToRichText(mdContent string) string {
	// 对于Fyne的RichText组件，我们可以直接返回Markdown内容
	// 因为Fyne的RichText支持Markdown语法
	return mdContent
}

// ExtractOutline 从Markdown内容中提取大纲
func (r *Renderer) ExtractOutline(mdContent string) []core.OutlineEntry {
	var outline []core.OutlineEntry

	// 按行分割内容
	lines := strings.Split(mdContent, "\n")

	// 正则表达式匹配标题
	headerRegex := regexp.MustCompile(`^(#{1,6})\s+(.+)$`)

	for lineNum, line := range lines {
		line = strings.TrimSpace(line)

		// 匹配标题行
		matches := headerRegex.FindStringSubmatch(line)
		if len(matches) == 3 {
			level := len(matches[1]) // #的数量就是标题级别
			title := strings.TrimSpace(matches[2])

			// 创建大纲条目
			entry := core.OutlineEntry{
				Title: title,
				Level: level,
				Line:  lineNum + 1, // 行号从1开始
			}

			outline = append(outline, entry)
		}
	}

	return outline
}

// GenerateTableOfContents 生成目录HTML
func (r *Renderer) GenerateTableOfContents(outline []core.OutlineEntry) string {
	if len(outline) == 0 {
		return ""
	}

	var buf bytes.Buffer
	buf.WriteString(`<div class="toc">`)
	buf.WriteString(`<h3>目录</h3>`)

	currentLevel := 0
	for _, entry := range outline {
		// 处理级别变化
		if entry.Level > currentLevel {
			// 开启新的列表级别
			for i := currentLevel; i < entry.Level; i++ {
				buf.WriteString(`<ul>`)
			}
		} else if entry.Level < currentLevel {
			// 关闭多余的列表级别
			for i := entry.Level; i < currentLevel; i++ {
				buf.WriteString(`</ul>`)
			}
		}

		// 生成锚点ID（移除特殊字符，用连字符替换空格）
		anchorID := strings.ToLower(entry.Title)
		anchorID = regexp.MustCompile(`[^\w\s-]`).ReplaceAllString(anchorID, "")
		anchorID = regexp.MustCompile(`\s+`).ReplaceAllString(anchorID, "-")
		anchorID = strings.Trim(anchorID, "-")

		// 添加目录项
		buf.WriteString(`<li><a href="#`)
		buf.WriteString(anchorID)
		buf.WriteString(`">`)
		buf.WriteString(template.HTMLEscapeString(entry.Title))
		buf.WriteString(`</a></li>`)

		currentLevel = entry.Level
	}

	// 关闭所有未关闭的列表
	for i := 0; i < currentLevel; i++ {
		buf.WriteString(`</ul>`)
	}

	buf.WriteString(`</div>`)
	return buf.String()
}

// RenderToHTMLWithTOC 渲染Markdown并包含目录
func (r *Renderer) RenderToHTMLWithTOC(mdContent, title string) string {
	// 提取大纲
	outline := r.ExtractOutline(mdContent)

	// 生成目录
	toc := r.GenerateTableOfContents(outline)

	// 渲染主要内容
	htmlContent := r.RenderToHTML(mdContent)

	// 将目录插入到内容前面
	contentWithTOC := toc + htmlContent

	// 使用模板渲染完整页面
	return r.RenderToHTMLWithTemplate(contentWithTOC, title)
}

// ValidateMarkdown 验证Markdown语法
func (r *Renderer) ValidateMarkdown(mdContent string) []string {
	var warnings []string

	// 检查常见的Markdown语法问题
	lines := strings.Split(mdContent, "\n")

	for lineNum, line := range lines {
		lineNum++ // 行号从1开始

		// 检查标题格式
		if strings.HasPrefix(line, "#") {
			if !regexp.MustCompile(`^#{1,6}\s+.+`).MatchString(line) {
				warnings = append(warnings,
					fmt.Sprintf("第%d行: 标题格式不正确，# 后面应该有空格", lineNum))
			}
		}

		// 检查链接格式
		if strings.Contains(line, "](") {
			if !regexp.MustCompile(`\[.*?\]\(.*?\)`).MatchString(line) {
				warnings = append(warnings,
					fmt.Sprintf("第%d行: 链接格式可能不正确", lineNum))
			}
		}
	}

	return warnings
}

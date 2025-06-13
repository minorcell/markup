package core

import (
	"os"
	"sync"
)

// OutlineEntry 代表大纲中的一个条目
type OutlineEntry struct {
	Title string // 标题文本
	Level int    // 标题级别 (1-6)
	Line  int    // 行号
}

// AppState 应用状态管理
type AppState struct {
	mutex           sync.RWMutex   // 读写锁
	currentFile     string         // 当前打开的文件路径
	currentContent  string         // 当前文件内容
	originalContent string         // 原始文件内容（用于检测变更）
	outline         []OutlineEntry // 文档大纲
}

// NewAppState 创建新的应用状态实例
func NewAppState() *AppState {
	return &AppState{
		outline: make([]OutlineEntry, 0),
	}
}

// GetCurrentFile 获取当前文件路径
func (s *AppState) GetCurrentFile() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.currentFile
}

// SetCurrentFile 设置当前文件路径
func (s *AppState) SetCurrentFile(filePath string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.currentFile = filePath
}

// GetCurrentContent 获取当前内容
func (s *AppState) GetCurrentContent() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.currentContent
}

// SetCurrentContent 设置当前内容
func (s *AppState) SetCurrentContent(content string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.currentContent = content
}

// GetOriginalContent 获取原始内容
func (s *AppState) GetOriginalContent() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.originalContent
}

// SetOriginalContent 设置原始内容
func (s *AppState) SetOriginalContent(content string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.originalContent = content
}

// GetOutline 获取文档大纲
func (s *AppState) GetOutline() []OutlineEntry {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	// 返回副本以避免并发修改
	outline := make([]OutlineEntry, len(s.outline))
	copy(outline, s.outline)
	return outline
}

// SetOutline 设置文档大纲
func (s *AppState) SetOutline(outline []OutlineEntry) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.outline = outline
}

// HasUnsavedChanges 检查是否有未保存的变更
func (s *AppState) HasUnsavedChanges() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.currentContent != s.originalContent
}

// LoadFile 加载文件内容
func (s *AppState) LoadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// SaveFile 保存文件内容
func (s *AppState) SaveFile(filePath, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}

// Reset 重置状态
func (s *AppState) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.currentFile = ""
	s.currentContent = ""
	s.originalContent = ""
	s.outline = make([]OutlineEntry, 0)
}

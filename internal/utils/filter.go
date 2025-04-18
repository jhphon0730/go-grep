package utils

import "strings"

// 필터링할 확장자 or 경로 키워드
var skipKeywords = []string{
	".git", "node_modules", "vendor", ".idea", ".vscode",
	".exe", ".png", ".jpg", ".jpeg", ".gif", ".svg",
	".mp3", ".mp4", ".avi", ".mov",
	".pdf", ".doc", ".docx", ".xls", ".xlsx",
	".ppt", ".pptx", ".zip", ".tar", ".gz", ".rar",
	".7z", ".iso",
}

// 이 함수가 true면, 해당 파일/디렉토리는 grep 대상에서 제외됨
func ShouldSkipPath(path string) bool {
	path = strings.ToLower(path)
	for _, keyword := range skipKeywords {
		if strings.Contains(path, keyword) {
			return true
		}
	}
	return false
}

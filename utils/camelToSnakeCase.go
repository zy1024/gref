package utils

import (
	"regexp"
	"strings"
)

// CamelToSnakeCase converts a camelCase string to snake_case
// Example: "CamelCaseString" -> "camel_case_string"
// Example: "HTTPRequest" -> "http_request"
func CamelToSnakeCase(camel string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(camel, "${1}_${2}")
	// 处理字符串开头的边界条件
	snake = strings.ToLower(snake)
	return snake
}

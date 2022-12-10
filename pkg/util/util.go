package util

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"yqsl.com/go/seed/conf"
	"yqsl.com/go/seed/pkg/i18n"
)

func CheckIDs(con conf.App, ctx context.Context, ids []uint64) (err error) {
	if len(ids) <= 0 {
		err = errors.New(i18n.T(ctx, "ids不能为空"))
	} else if len(ids) > int(con.MaxBatch) {
		err = errors.New(i18n.Tf(ctx, "ids的数量不能超过最大值%d", con.MaxBatch))
	}
	return
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

//CamelString 蛇形转驼峰
//xx_yy to XxYx  xx_y_y to XxYY
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
func TrimHtml(src string) string {
	// 将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	// 去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	// 去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	// 去除所有尖括号内的HTML代码，并换成空格
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, " ")
	// 去除连续的空格
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, " ")
	return strings.TrimSpace(src)
}

func CutStr(src string, length int) string {
	if length <= 0 {
		return ""
	}
	// 去除HTML标签
	snippet := TrimHtml(src)
	// 字符串长度截取，无论中英文都算作一个字符
	snippetRune := []rune(snippet)
	if len(snippetRune) > length {
		snippet = string(snippetRune[:length]) + "..."
	}
	return snippet
}

// Camel2Snake 驼峰转蛇形
func Camel2Snake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	// 大写统一转小写
	return strings.ToLower(string(data[:]))
}

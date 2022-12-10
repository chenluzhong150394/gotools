package util

import (
	"regexp"
)

func MathPattern(p string, text string) (result []string, num int) {

	reg := regexp.MustCompile(p)

	result = reg.FindAllString(text, -1)

	num = len(result)

	return

}

func IsMathPattern(p string, text string) (result bool, err error) {
	result, err = regexp.MatchString(p, text)
	return
}

func RePlaceAllSpaceString(pattern string, src string) (d string) {
	reg := regexp.MustCompile(pattern)
	d = reg.ReplaceAllString(src, "")
	//fmt.Println(d)
	return
}

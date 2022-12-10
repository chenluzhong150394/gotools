package test

import (
	"fmt"
	"github.com/chenluzhong150394/gotools/pkg/util"
	"testing"
)

func TestXml(t *testing.T) {

	type Iterm struct {
		LogisticsCode string `xml:"logisticsCode"`
		LogisticsName string `xml:"logisticsName"`
	}

	iterm := Iterm{
		LogisticsCode: "123",
		LogisticsName: "123",
	}

	b := make([]interface{}, 0)
	b = append(b, iterm)

	output, err := util.ParseXml(b)
	if err != nil {
		return
	}
	fmt.Println(string(output))

}

func TestRegx(t *testing.T) {

	result, num := util.MathPattern(`(?U)\b.+\b`, "Hello 世界！123 Go.")

	fmt.Println(result, num)
}

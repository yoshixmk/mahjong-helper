package webapi

import (
	"bytes"
	"io"
	"os"
)

type ApiData struct {
	// 数据更新时间戳
	Timestamp int `json:"timestamp"`

	// 自家手牌 一个长度为 34 的整数数组
	Counts []int `json:"counts"`

	// 手牌危险度 一个长度为 34 的浮点数组
	RiskTable []float64 `json:"risk"`

	DiscardTiles []int `json:"discardTiles"`
	DiscardTiles1 []int `json:"discardTiles1"`
	DiscardTiles2 []int `json:"discardTiles2"`
	DiscardTiles3 []int `json:"discardTiles3"`

	output_buffer bytes.Buffer
}

func (data *ApiData) Init() {
	data.output_buffer.Reset()
} 

// implement the io.Writer interface
var _ io.Writer = (*ApiDataConvertor)(nil)
type ApiDataConvertor struct {
	*ApiData
}

func (writer ApiDataConvertor) Write(p []byte) (n int, err error) {
	n, e := writer.output_buffer.Write(p)
	if e != nil {
		return n, e
	}
	return os.Stdout.Write(p)
}




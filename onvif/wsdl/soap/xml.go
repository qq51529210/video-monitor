package soap

import (
	"encoding/xml"
	"fmt"
	"strings"
)

const (
	Namespace = "http://www.w3.org/2003/05/soap-envelope"
)

// Header 可以用于 Envelope 的 Header 字段
type Header[Data any] struct {
	XMLName xml.Name `xml:"env:Header"`
	Data    Data
}

// Body 可以用于 Envelope 的 Body 字段
type Body[Data any] struct {
	XMLName xml.Name `xml:"env:Body"`
	Data    Data
	Fault   *Fault `xml:",omitempty"`
}

// Envelope 表示整个 xml 消息
type Envelope struct {
	XMLName xml.Name    `xml:"env:Envelope"`
	Attr    []*xml.Attr `xml:",attr"`
}

// Fault 表示 Envelope.Body 中的错误
type Fault struct {
	XMLName xml.Name     `xml:"env:Fault"`
	Code    *FaultCode   `xml:"env:Code"`
	Reason  *FaultReason `xml:"env:Reason"`
	Detail  *FaultDetail `xml:"env:Detail"`
}

func (c *Fault) Error() string {
	if c.Detail != nil {
		return fmt.Sprintf("code: %s, %s",
			c.Detail.ErrorCode, c.Detail.Description)
	}
	if c.Reason != nil {
		return c.Reason.Text
	}
	if c.Code != nil {
		var str strings.Builder
		fmt.Fprintf(&str, "code: %s", c.Code.Value)
		if c.Code.Subcode != nil {
			fmt.Fprintf(&str, " sub code: %s", c.Code.Subcode.Value)
		}
		return str.String()
	}
	return "unknown fault"
}

// FaultCode 表示 Fault 的 Code 字段
type FaultCode struct {
	Value   string     `xml:"Value"`
	Subcode *FaultCode `xml:"Subcode,omitempty"`
}

// FaultReason 表示 Fault 的 Reason 字段
type FaultReason struct {
	Text string `xml:"Text"`
}

// FaultDetail 表示 Fault 的 Detail 字段
type FaultDetail struct {
	ErrorCode   string `xml:"ErrorCode"`
	Description string `xml:"Description"`
}

// NewNamespaceAttr 返回命名空间属性
func NewNamespaceAttr() *xml.Attr {
	return &xml.Attr{
		Name: xml.Name{
			Local: "xmlns:env",
		},
		Value: Namespace,
	}
}

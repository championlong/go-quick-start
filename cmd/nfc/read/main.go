package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/ebfe/scard"
)

const (
	cla     = 0xFF
	insRead = 0xB0
	p1      = 0x00
)

func main() {
	ctx, err := scard.EstablishContext()
	if err != nil {
		log.Fatal("无法建立上下文:", err)
	}
	defer ctx.Release()

	readers, err := ctx.ListReaders()
	if err != nil {
		log.Fatal("无法列出读卡器:", err)
	}
	if len(readers) == 0 {
		log.Fatal("未找到读卡器")
	}

	card, err := ctx.Connect(readers[0], scard.ShareShared, scard.ProtocolAny)
	if err != nil {
		log.Fatal("连接读卡器失败:", err)
	}
	defer card.Disconnect(scard.LeaveCard)

	// 读取用户数据区（页0x04到页0x86）
	var data []byte
	for page := 0x04; page <= 0x86; page++ {
		cmd := []byte{cla, insRead, p1, byte(page), 0x04}
		rsp, err := card.Transmit(cmd)
		if err != nil {
			log.Printf("页 0x%02X 读取失败: %v", page, err)
			continue
		}

		// 检查响应状态
		if len(rsp) < 2 || rsp[len(rsp)-2] != 0x90 || rsp[len(rsp)-1] != 0x00 {
			log.Printf("页 0x%02X 状态异常: %X", page, rsp)
			continue
		}
		fmt.Print(string(rsp[:len(rsp)-2]))
		data = append(data, rsp[:len(rsp)-2]...)
	}

	// 解析NDEF数据
	if url := parseNDEF(data); url != "" {
		fmt.Println("提取的URL:", url)
	} else {
		fmt.Println("未找到URL记录")
	}
}

// NDEF解析逻辑
func parseNDEF(data []byte) string {
	// 查找NDEF消息头 D1 01（绝对URI记录）

	startIndex := bytes.Index(data, []byte{0xD1, 0x01})
	if startIndex == -1 {
		return ""
	}
	startIndex += 2 // 跳过D1 01

	// 边界检查
	if startIndex+2 > len(data) {
		return ""
	}

	// 解析关键字段
	payloadLength := int(data[startIndex])   // 负载长度（包含URI标识符）
	uriIdentifier := data[startIndex+2]      // URI前缀代码
	urlStart := startIndex + 3               // URL内容起始位置
	urlEnd := urlStart + (payloadLength - 1) // 计算URL结束位置（减去标识符）

	// 边界保护
	if urlEnd > len(data) {
		urlEnd = len(data)
	}
	urlBytes := data[urlStart:urlEnd]

	// 截断填充符（0xFE或0x00）
	if end := bytes.IndexByte(urlBytes, 0xFE); end != -1 {
		urlBytes = urlBytes[:end]
	}
	if end := bytes.IndexByte(urlBytes, 0x00); end != -1 {
		urlBytes = urlBytes[:end]
	}

	// 构建完整URL
	return buildURL(uriIdentifier, urlBytes)
}

// URL构建逻辑
func buildURL(identifier byte, urlBytes []byte) string {
	var prefix string
	switch identifier {
	case 0x01:
		prefix = "http://www."
	case 0x02:
		prefix = "https://www."
	case 0x03:
		prefix = "http://"
	case 0x04:
		prefix = "https://"
	default:
		prefix = ""
	}

	// 直接转换为字符串（写入时保证UTF-8）
	return prefix + strings.TrimSpace(string(urlBytes))
}

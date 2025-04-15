package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ebfe/scard"
	"github.com/xuri/excelize/v2"
)

// CGO_ENABLED=1 go build
// https://www.acs.com.hk/cn/drivers/
// 使用NFC读取器ACR122U

const (
	cla           = 0xFF
	insWrite      = 0xD6 // UPDATE BINARY指令
	userPageStart = 0x04 // 用户数据起始页
	pageSize      = 4    // 每页4字节
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("请通过命令行指定Excel文件路径，例如: ./nfc-writer urls.xlsx")
	}

	// 从Excel读取URL列表
	excelPath := os.Args[1]
	urls, err := readURLsFromExcel(excelPath)
	if err != nil {
		log.Fatal("读取Excel失败:", err)
	}
	if len(urls) == 0 {
		log.Fatal("Excel文件中未找到有效URL")
	}

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
	reader := readers[0]

	input := bufio.NewReader(os.Stdin)

	for _, url := range urls {
		fmt.Printf("\n=== 准备写入: %s ===\n", url)
		fmt.Println("请将新的NFC标签放置在读卡器上，然后按回车键继续...")
		_, _ = input.ReadString('\n') // 等待用户回车

		// 每次循环重新连接以确保识别新标签
		card, err := ctx.Connect(reader, scard.ShareExclusive, scard.ProtocolAny)
		if err != nil {
			log.Printf("连接读卡器失败: %v", err)
			continue
		}

		// 构建并写入NDEF数据
		if err := writeURLToTag(card, url); err != nil {
			log.Printf("写入失败: %v", err)
		} else {
			fmt.Printf("成功写入: %s\n", url)
		}

		card.Disconnect(scard.LeaveCard)
	}
}

// 写入单个标签
func writeURLToTag(card *scard.Card, url string) error {
	// 构建NDEF数据
	ndefData := buildNDEF(url)
	dataBuffer := bytes.NewBuffer(ndefData)

	// 分页写入
	page := userPageStart
	for dataBuffer.Len() > 0 {
		pageData := dataBuffer.Next(pageSize)
		if len(pageData) < pageSize {
			pageData = append(pageData, bytes.Repeat([]byte{0x00}, pageSize-len(pageData))...)
		}

		cmd := []byte{
			cla,
			insWrite,
			0x00,       // P1 (页地址高字节)
			byte(page), // P2 (页地址低字节)
			byte(len(pageData)),
		}
		cmd = append(cmd, pageData...)

		rsp, err := card.Transmit(cmd)
		if err != nil {
			return fmt.Errorf("传输失败: %w", err)
		}

		// 检查APDU响应状态
		if len(rsp) < 2 || rsp[len(rsp)-2] != 0x90 || rsp[len(rsp)-1] != 0x00 {
			return fmt.Errorf("页0x%02X写入异常: %X", page, rsp)
		}

		page++
	}
	return nil
}

// 构建NDEF字节流
func buildNDEF(url string) []byte {
	// 提取URI标识符和内容
	prefixCode, trimmedURL := parseURI(url)
	urlBytes := []byte(trimmedURL)

	// 构建NDEF记录
	ndefHeader := []byte{
		0x03,                    // TLV类型（NDEF消息）
		byte(len(urlBytes) + 5), // TLV长度（D1 01 + 标识符 + URL）
		0xD1,                    // NDEF Record Header (MB=1, ME=1, TNF=0x01)
		0x01,                    // 类型长度
		byte(len(urlBytes) + 1), // 负载长度
		'U',                     // 类型（URI）
		prefixCode,              // URI标识符
	}
	ndefData := append(ndefHeader, urlBytes...)

	// 添加终止TLV（0xFE）
	terminator := []byte{0xFE}
	return append(ndefData, terminator...)
}

// URI解析函数（同之前示例）
func parseURI(url string) (byte, string) {
	switch {
	case strings.HasPrefix(url, "https://www."):
		return 0x02, strings.TrimPrefix(url, "https://www.")
	case strings.HasPrefix(url, "http://www."):
		return 0x01, strings.TrimPrefix(url, "http://www.")
	case strings.HasPrefix(url, "https://"):
		return 0x04, strings.TrimPrefix(url, "https://")
	case strings.HasPrefix(url, "http://"):
		return 0x03, strings.TrimPrefix(url, "http://")
	default:
		return 0x00, url
	}
}

// 从Excel文件读取URL列表（假设URL在Sheet1的A列）
func readURLsFromExcel(path string) ([]string, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// 获取Sheet1的A列数据
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	var urls []string
	for _, row := range rows {
		if len(row) > 0 && strings.HasPrefix(row[0], "http") {
			urls = append(urls, row[0])
		}
	}
	return urls, nil
}

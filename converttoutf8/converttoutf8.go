/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2024-04-19 10:41:49
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2024-04-19 10:43:10
 * @FilePath: \xlsx_del\pkg\converttoutf8\converttoutf8.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package converttoutf8

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// 将字节切片转换为UTF-8编码的字符串
func ConvertBytesToUTF8(input []byte) (string, error) {
	// 创建转换器
	utf8Reader := transform.NewReader(strings.NewReader(string(input)), unicode.UTF8.NewDecoder())

	// 读取并转换内容为UTF-8编码
	output, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// 将文本转换为UTF-8编码
func ConvertStringsToUTF8(text string) (string, error) {
	// 将文本从GBK转换为UTF-8
	gbkReader := transform.NewReader(strings.NewReader(text), simplifiedchinese.GBK.NewDecoder())

	// 读取并转换内容为UTF-8编码
	utf8Bytes, err := ioutil.ReadAll(gbkReader)
	if err != nil {
		return "", fmt.Errorf("转换编码失败：%v", err)
	}

	return string(utf8Bytes), nil
}

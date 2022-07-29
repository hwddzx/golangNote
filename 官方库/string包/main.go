package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is an example of a string"

	var resInt int
	var resBool bool
	var resString string

	// 前缀
	resBool = strings.HasPrefix(str, "Th")
	fmt.Println(resBool)
	// 后缀
	resBool = strings.HasSuffix(str, "string")
	fmt.Println(resBool)
	// 包含
	resBool = strings.Contains(str, "an")
	fmt.Println(resBool)

	// 索引
	resInt = strings.Index(str, "is")
	fmt.Println(resInt)
	resInt = strings.LastIndex(str, "is")
	fmt.Println(resInt)

	// 替换
	str = "你好世界,这个世界真好"
	newStr := "地球"
	oldStr := "世界"
	resString = strings.Replace(str, oldStr, newStr, 1)
	fmt.Println(resString)

	// 统计
	resInt = strings.Count(str, "好")
	fmt.Println(resInt)

	// 大小写转换
	var orig = "How are you?"
	fmt.Println(strings.ToUpper(orig))
	fmt.Println(strings.ToLower(orig))

	// 修剪
	fmt.Printf("%q\n", strings.Trim(" !!! Golang !!! ", "! "))
	fmt.Printf("%q\n", strings.Trim(" !!! Golang !!! ", " ! "))
	fmt.Printf("%q\n", strings.Trim(" !!! Golang !!! ", " !"))
	fmt.Printf("%q\n", strings.Trim(" !!! Golang !!! ", "!"))

	// 分隔
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))

	// 插入字符 join
	str = "The quick brown fox jumps over the lazy dog 中文"
	strSli := strings.Fields(str) // 等同于 strings.Split(str, " ")
	str2 := strings.Join(strSli, ",")
	fmt.Println(str2)
}

package main

import "fmt"

func main() {
	s := `aA你2`
	fmt.Println("字符串长度:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	s += "世界"
	fmt.Println(s)

	s = `我是中国人`

	// 遍历
	fmt.Println(`遍历`)
	for i := 0; i < len(s); i++ {
		fmt.Printf(`%c`, s[i])
	}
	fmt.Println()
	for _, v := range s {
		fmt.Printf(`%c`, v)
	}
	fmt.Println()


	// 字符串修改
	fmt.Println(`字符串修改`)
	// 修改字符串中的字节
	s = `Hello 世界!`
	b := []byte(s)
	fmt.Println(b)
	b[5] = ','
	fmt.Println(b)
	c := string(b)
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", b)
	fmt.Printf("%s\n", c)

	// 修改字符串中的字符
	r := []rune(s)
	fmt.Println(r)
	r[6] = '中'
	r[7] = '国'
	fmt.Println(s)
	fmt.Println(string(r))

}

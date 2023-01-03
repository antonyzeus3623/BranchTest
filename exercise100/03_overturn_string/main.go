// 请实现⼀个算法，在不使⽤【额外数据结构和储存空间】的情况下，翻转⼀个给定的字符串(可以使⽤单个过程变量)。
// 给定⼀个string，请返回⼀个string，为翻转后的字符串。保证字符串的⻓度⼩于等于5000。
// 思路：翻转字符串其实是将⼀个字符串以中间字符为轴，前后翻转，即将str[len]赋值给str[0],将str[0] 赋值 str[len]。

package main

import (
	"fmt"
	"log"
)

func main() {
	s := "abcdefg123456"
	_s, err := overTurnString(s)
	if err != nil {
		log.Println("转换失败, err: ", err)
	}
	log.Printf("字符串%s转换后的结果为:%s", s, _s)
}

func overTurnString(s string) (string, error) {
	str := []rune(s)
	l := len(str)
	if l > 5000 {
		return "", fmt.Errorf("字符串的长度超过了5000")
	}
	for i := 0; i < l/2; i++ {
		// tmp := str[i]
		// str[i] = str[l-i-1]
		// str[l-i-1] = tmp
		str[i], str[l-i-1] = str[l-i-1], str[i]
	}
	return string(str), nil
}

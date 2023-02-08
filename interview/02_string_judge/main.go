package main

import (
	"log"
	"strings"
)

// 判断字符串中字符是否全都不同
//请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。
//给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。
//字符串的长度小于等于【3000】。

func main() {
	str1 := "abcdfrhon"
	ret1 := isStringUnique1(str1)
	log.Println("是否唯一：", ret1)

	str2 := "abcdfrhona"
	ret2 := isStringUnique2(str2)
	log.Println("是否唯一：", ret2)

}

func isStringUnique1(str string) bool {
	n := strings.Count(str, "")
	if n > 3000 {
		return false
	}
	for _, v := range str {
		if v > 127 { // ASCII字符字符一共有256个，其中128个是常用字符，可以在键盘上输入。128之后的是键盘上无法找到的
			return false
		}
		if strings.Count(str, string(v)) > 1 {
			return false
		}
	}
	return true
}

func isStringUnique2(str string) bool {
	n := strings.Count(str, "")
	if n > 3000 {
		return false
	}
	for k, v := range str {
		if v > 127 {
			return false
		}
		if strings.Index(str, string(v)) != k {
			return false
		}
	}
	return true
}

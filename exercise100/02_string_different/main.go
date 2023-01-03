// 请实现⼀个算法，确定⼀个字符串的所有字符【是否全都不同】。这⾥我们要求【不允许使⽤额外的存储结构】。
// 给定⼀个string，请返回⼀个bool值,true代表所有字符全都不同，false代表存在相同的字符。
// 保证字符串中的字符为【ASCII字符】。字符串的⻓度⼩于等于【3000】。
package main

import (
	"log"
	"strings"
)

func main() {
	s := "ABCDEFGJA"
	b := isUniqueString(s)
	log.Printf("字符串%s, 是否全不相同：%v\n", s, b)

}

func isUniqueString(s string) bool {
	if strings.Count(s, " ") > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

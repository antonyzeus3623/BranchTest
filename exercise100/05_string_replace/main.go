// 请编写⼀个⽅法，将字符串中的空格全部替换为“%20”。
// 假定该字符串有⾜够的空间存放新增的字符，并且知道字符串的真实⻓度(⼩于等于1000)，同时保证字符串由【⼤⼩写的英⽂字⺟组成】。
// 给定⼀个string为原始的串，返回替换后的string。
package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

func main() {
	s := "ABD efd1"
	result, err := replaceString(s)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)

}

func replaceString(s string) (string, error) {
	if len([]rune(s)) > 1000 {
		return s, fmt.Errorf("string长度超过了1000")
	}
	for _, v := range s {
		if !unicode.IsSpace(v) && !unicode.IsLetter(v) {
			return s, fmt.Errorf("string包含非字母元素")
		}
	}
	return strings.Replace(s, " ", "%20", -1), nil
}

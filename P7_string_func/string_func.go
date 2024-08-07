package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字串比大小 
	// strings.Compare(a, b)
	// a == b: return 0, a > b: return 1, a < b: return -1
	// 直接用運算子比較快～
	println(strings.Compare("addd", "b"))
	println(strings.Compare("aqqq", "aqqq"))
	println(strings.Compare("z", "x"))
	println("z" < "x")

	// 字串包含
	// strings.Contains(a, b)
	// b in a: return true, b not in a return false
	println(strings.Contains("hello world", "llo"))
	println(strings.Contains("hello world", "a"))

	// strings.Index(str, sub_str)
	// 回傳 sub_str 在 str 的起始位置, 沒有則 return -1
	println(strings.Index("abcdefg", "cd"))
	println(strings.Index("abcdefg", "z"))

	// strings.Join([]string,  sep string)
	// 將 sep 加入每個元素之間, 並返回成一個字串
	str_array := []string{"標籤1", "標籤2", "標籤3"}
	println(strings.Join(str_array, ", "))

	// strings.Count(str, substr)
	// 計算 substr 在 str 出現過幾次, 若 substr 是空字串, 則回傳 str 有幾個字 + 1
	println(strings.Count("你好嗎, 我很好", "好"))
	println(strings.Count("你好嗎", ""))

	// strings.Replace(str, old_str, new_str, n)
	// replace old_str to new_str in str, if n < 0: replace no limit, else replace n times
	// n < 0 = strings.ReplaceAll(str, old_str, new_str)
	println(strings.Replace("oink oink oink", "oink", "moo", 2))
	println(strings.Replace("oink oink oink", "oink", "moo", -1))
	println(strings.ReplaceAll("oink oink oink", "oink", "moo"))

	// strings.Split(str, sep string)
	fmt.Printf("%q\n", strings.Split("標籤1, 標籤2, 標籤3", ", "))

	println(strings.ToUpper("abcdefg"))
	println(strings.ToLower("ZXCVB"))

	// strings.TrimSpace(str)
	// 去除頭尾的空格
	println(strings.TrimSpace("   as  sdsa \n"))
}
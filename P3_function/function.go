package main

import (
	"fmt"
	"math/rand"
	"os"
)

var env = os.Getenv("ENV")

func init() {
	if env == "" {
		panic("no value for $ENV")
	}
}

func throwsPanic(f func()) (b bool) {
    defer func() {
        if x := recover(); x != nil {
            b = true
        }
    }()
    f() //執行函式 f，如果 f 中出現了 panic，那麼就可以恢復回來
    return
}

func myGoto() {
	var i int = 0
Here: // 這行的第一個詞，以冒號結束作為標籤
	fmt.Println(i)
	i++

	if i == 5 {
		return
	}

	goto Here // 跳轉到 Here 去
}

func myFor() {
	for i := 0; i < 5; i++ {
		fmt.Println(i * i)
	}

	sum := 1

	// for ; sum < 1000; sum++
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	var numbers map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}

	// like for k, v in numbers.items()
	for k, v := range numbers {
		fmt.Println(k, v)
	}

	// go 如果有沒使用到的變數會報錯, 可用 _ 丟棄
	for _, v := range numbers {
		fmt.Println(v)
	}
}

func mySwitch() {
	i := rand.Intn(4)

	switch i {
	case 0:
		fmt.Println("i is equal to 0")
	case 1:
		fmt.Println("i is equal to 1")
	case 2:
		fmt.Println("i is equal to 2")
	default:
		fmt.Println(i)
	}
}

// output 部分可不宣告, 但型別一定要宣告
func myFunc(input1 int, input2 string) (output1 int, output2 string) {
	return input1, input2
}

func max(a, b int) (int) {
	if a > b {
		return a
	}

	return b
}

// 可變參數
func myArgs(arg ...int) {
	for _, v := range arg {
		fmt.Println(v)
	}
}

func myPointer(a *int) (int) {
	*a = *a + 1
	return *a
}

type testInt func(int) bool // 宣告一個函式型別

func isOdd(integer int) bool {
	return integer % 2 != 0
}

func isEven(integer int) bool {
	return integer % 2 == 0
}

func filterInt(slice []int, function testInt) []int {
	var result []int

	for _, value := range slice {
		if function(value) {
			result = append(result, value)
		}
	}

	return result
}


func main() {
	randInt := rand.Intn(100)

	if randInt < 50 {
		fmt.Printf("random value is less than 50\n")
	} else {
		fmt.Printf("random value is greater than 50\n")
	}

	myGoto()
	myFor()
	mySwitch()
	fmt.Println(myFunc(123, "abc"))
	myArgs(1, 2, 3, 4, 5)

	slice := []int {1, 2, 3, 4, 5, 6, 7}
	odd := filterInt(slice, isOdd)
	fmt.Println(odd)
}

package main

import (
	"fmt"
)

var intValue int
var stringValue string
var boolValue bool

var intValue1 int = 1

var intValue2, stringValue2, boolValue2 = 2, "2", true

var s1, s2, s3 string = "1", "2", "3"

const (
	i = 100
	pi = 3.14
	prefix = `P2_`
)

const (
    x = iota // x == 0
    y = iota // y == 1
    z = iota // z == 2
    w        // 常數宣告省略值時，預設和之前一個值的字面相同。這裡隱含的說 w = iota，因此 w == 3。其實上面 y 和 z 可同樣不用"= iota"
)

func main() {
	// 短變數宣告 := 只能在 function 內部使用
	// vname1, vname2, vname3 := "v1", "v2", "v3"
	// fmt.Printf(vname1 + vname2 + vname3)
	// s1 = `s1
	// 	s2
	// 		s3`
	// fmt.Printf(s1)

	// err := errors.New(`some error`)

	// if err != nil {
	// 	fmt.Print(err)
	// }

	// fmt.Println(s1, s2, s3)

	fmt.Println(x, y, z, w)

	// 每遇到一個 const 關鍵字，iota 就會重置，此時 v == 0
	const v = iota
	fmt.Println(x, y, z, w, v)

	const (
		h, i, j = iota, iota, iota //h=0,i=0,j=0 iota 在同一行值相同
	)

	const (
		a       = iota //a=0
		b       = "B"
		c       = iota             //c=2
		d, e, f = iota, iota, iota //d=3,e=3,f=3
		g       = iota             //g = 4
	)

	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)

	var arr [10]int
	arr[0] = 1111
	arr[1] = 2222
	fmt.Printf("The first element of the array is %d\n", arr[0])
	// 最後一個元素, 但未賦值, return 0
	fmt.Printf("The last element of the array is %d\n", arr[9])

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// 宣告了一個長度為 10 的 int 陣列，其中前三個元素初始化為 1、2、3，其它預設為 0
	arr3 := [10]int{1, 2, 3}
	fmt.Println(arr3)

	// 可以省略長度而採用`...`的方式，Go 會自動根據元素個數來計算長度
	arr4 := [...]int{1, 2, 3, 4}
	fmt.Println(arr4)

	// 宣告了一個二維陣列，該陣列以兩個陣列作為元素，其中每個陣列中又有 4 個 int 型別的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	// 上面的宣告可以簡化，直接忽略內部的型別
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(doubleArray, easyArray)

	// 動態陣列, like python list, slice 與 array 的差別在於有沒有寫明陣列的長度, 或使用 `...` 自動計算
	var fslice []int
	fmt.Println(fslice)

	slice := []byte {'a', 'b', 'c', 'd'}
	fmt.Println(slice)

	// 宣告一個含有 10 個元素元素型別為 byte 的陣列
	var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	// 宣告兩個含有 byte 的 slice
	var aslice, bslice []byte

	// a 指向陣列的第 3 個元素開始，併到第五個元素結束，
	aslice = ar[2:5]
	//現在 a 含有的元素: ar[2]、ar[3]和 ar[4]

	// b 是陣列 ar 的另一個 slice
	bslice = ar[3:5]
	// b 的元素是：ar[3]和 ar[4]

	fmt.Println(aslice, bslice)

	// slice[0:n] = slice[:n]
	// slice[n:len(slice)] = slice[n:]
	// 宣告一個陣列
	var array = [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	// 宣告兩個 slice
	var aSlice, bSlice []string

	// 示範一些簡便操作
	aSlice = array[:3] // 等價於 aSlice = array[0:3] aSlice 包含元素: a,b,c
	fmt.Println(aSlice)
	aSlice = array[5:] // 等價於 aSlice = array[5:10] aSlice 包含元素: f,g,h,i,j
	fmt.Println(aSlice)
	aSlice = array[:]  // 等價於 aSlice = array[0:10] 這樣 aSlice 包含了全部的元素
	fmt.Println(aSlice)

	// 從 slice 中取得 slice
	aSlice = array[3:7]  // aSlice 包含元素: d,e,f,g，len=4，cap=7
	fmt.Println(aSlice)
	bSlice = aSlice[1:3] // bSlice 包含 aSlice[1], aSlice[2] 也就是含有: e,f
	fmt.Println(bSlice)
	bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
	fmt.Println(bSlice)
	bSlice = aSlice[0:5] // 對 slice 的 slice 可以在 cap 範圍內擴充套件，此時 bSlice 包含：d,e,f,g,h
	fmt.Println(bSlice)
	bSlice = aSlice[:]   // bSlice 包含所有 aSlice 的元素: d,e,f,g
	fmt.Println(bSlice)

	// map = python's dict, format is map[keyType]valueType

	var numbers map[string]int
	// 另一種 map 的宣告方式
	numbers = make(map[string]int)
	numbers["one"] = 1  //賦值
	numbers["ten"] = 10 //賦值
	numbers["three"] = 3
	fmt.Println("第三個數字是: ", numbers["three"]) // 讀取資料

	// 初始化一個字典
	rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
	// map 有兩個回傳值，第二個回傳值，如果不存在 key，那麼 ok 為 false，如果存在 ok 為 true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C")  // 刪除 key 為 C 的元素
	fmt.Println(rating)

	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut"  // 現在 m["hello"]的值已經是 Salut 了
	fmt.Println(m)

	// 關於“零值”，所指並非是空值，而是一種“變數未填充前”的預設值，通常為 0。 此處羅列 部分類型 的 “零值”
	// int     0
	// int8    0
	// int32   0
	// int64   0
	// uint    0x0
	// rune    0 //rune 的實際型別是 int32
	// byte    0x0 // byte 的實際型別是 uint8
	// float32 0 //長度為 4 byte
	// float64 0 //長度為 8 byte
	// bool    false
	// string  ""
}
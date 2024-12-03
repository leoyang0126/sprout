package testAclass

import (
	"fmt"
	"sort"
	"strings"
)

//type Student struct {
//	string   name
//	int8     age
//}

func Test() {
	println("世界你好")
}

func Login(login string, passwd string) {
	go println(login, passwd)
	println("======")
}

func TestStr() {
	var strArray = strings.Split("as,qwe,qq", ",")
	fmt.Println("===TestStr===", strArray)

	str := strings.Join(strArray, "=")
	fmt.Println("===TestStr 拼接===", str)

	//language := language2.Chinese
	//collator := collate.New(language)
	//sort.Slice(strArray, func(i, j int) bool {
	//
	//})
	arrays := []int{2, 4, 6, 1, 3, 5}
	sort.Slice(arrays, func(i, j int) bool {
		return arrays[i] < arrays[j]
	})
	fmt.Println("arrays is ", arrays)
}

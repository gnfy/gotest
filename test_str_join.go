package main

import (
	"fmt"
	"sort"
)

func strJoin(m1 map[string]string) string {
	var strSlice []string
	// 将map中的key拼接成字符串
	for k, _ := range m1 {
		strSlice = append(strSlice, k)
	}
	// 对key进行升序
	sort.Strings(strSlice)
	// 拼接字符串
	var str string
	for _, v := range strSlice {
		str += fmt.Sprintf("%v=%v&", v, m1[v])
	}
	return str
}

func main() {
	var m1 = make(map[string]string)
	m1["name"] = "tom"
	m1["age"] = "20"
	m1["address"] = "beijing"
	fmt.Println("strJoin(m1): ", strJoin(m1))
}

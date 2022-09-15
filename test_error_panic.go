/* package main

import "errors"

// 读取文件
func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("read file error") // 返回一个error类型的变量
	}
}




func main() {
	defer func() {
		err := recover()
		if err != nil {
			println("read file error")
		}
	}()
	err := readFile("abc.go")
	if err != nil {
		panic(err)
	}
}
*/
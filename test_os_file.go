package main

import (
	"fmt"
	"os"
)

func readFile() {
	str, err := os.ReadFile("go.mod")
	if err != nil {
		fmt.Printf("err: %v", err)
	} else {
		fmt.Printf("str: %v\n", string(str[:]))
	}

}

func writeFile() {
	s := "hello world"
	err := os.WriteFile("test.txt", []byte(s), 0666)
	if err != nil {
		fmt.Printf("err: %v", err)
	} else {
		fmt.Printf("write success\n")
	}
}

func openFile() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("err: %v", err)
	} else {
		fmt.Printf("file: %v", file)
		file.Close()
	}
	f, err := os.OpenFile("test1.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("f: ", f)
		f.Close()
	}
}

func readOps() {
	// f, _ := os.Open("test.txt")
	// for {
	// 	buf := make([]byte, 3)
	// 	n, err := f.Read(buf)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println("n: ", n)
	// 	fmt.Printf("buf: %v\n", string(buf))
	// }
	// f.Close()
	f, _ := os.Open("test.txt")
	f.Seek(5, 0)
	buf := make([]byte, 10)
	// n, _ := f.ReadAt(buf, 2)
	n, _ := f.Read(buf)
	fmt.Println("n: ", n)
	fmt.Printf("buf at: %v\n", string(buf))
	f.Close()
}

func main() {
	readFile()
	writeFile()
	openFile()
	readOps()
}

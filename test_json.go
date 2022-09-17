/* package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID   int
	Name string
	Age  int
}

// 结构体标签
type Student2 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 嵌套结构体
type Class struct {
	Title    string
	Students []Student
}

func main() {
	var str = `{ "ID": 1, "Name": "Tom", "Age": 18 }`
	var stu Student
	json.Unmarshal([]byte(str), &stu)
	fmt.Println(stu)
	fmt.Printf("%#v", stu)
	var stu2 = Student2{
		ID:   1,
		Name: "Tom",
		Age:  18,
	}
	fmt.Printf("%#v\n", stu2)
	data, _ := json.Marshal(stu2)
	fmt.Println(string(data))

	c := Class{
		Title: "101",
		Students: []Student{
			{1, "Tom", 18},
			{2, "Jack", 19},
		},
	}
	fmt.Printf("c: %v\n", c)
	c2 := Class{
		Title:    "102",
		Students: make([]Student, 0),
	}

	for i := 0; i < 5; i++ {
		s := Student{
			ID:   i,
			Name: fmt.Sprintf("stu%02d", i),
			Age:  18 + i,
		}
		c2.Students = append(c2.Students, s)
	}
	fmt.Printf("c2: %v\n", c2)
	data2, _ := json.Marshal(c2)
	fmt.Println(string(data2))
}
*/
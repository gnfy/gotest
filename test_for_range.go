package main

/* import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	/* for i, v := range s {
		println(i, v)
	} */
	for _, v := range s {
		println(v)
	}
	m := make(map[string]string, 0)
	m["name"] = "Go"
	m["age"] = "18"
	for k, v := range m {
		println(k, v)
	}
	str := "hello"
	for _, v := range str {
		fmt.Printf("v: %c\n", v)
	}
}
 */
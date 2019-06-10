package main

import "fmt"

func main() {
	fmt.Println("Creating map...")
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println(m2 == nil) // false
	fmt.Println(m3 == nil) // true

	fmt.Println("Traversing map...")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values...")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	//errorname := m["errorname"]
	//fmt.Println(errorname)
	if errorname, ok := m["errorname"]; ok {
		fmt.Println(errorname)
	} else {
		fmt.Println("Key does not exist!")
	}

	fmt.Println("Deleting values...")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok, len(m))
}

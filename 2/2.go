package main

import "fmt"

type Animal struct {
	Name  string
	Class string
	Age   int8
}

func main() {
	moyu := Animal{
		Name:  "摸鱼",
		Class: "猫猫",
		Age:   2,
	}

	fmt.Printf("姓名:%s,种类:%s"+
		",年龄:%d\n", moyu.Name, moyu.Class, moyu.Age)
}

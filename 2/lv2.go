package main

import "fmt"

type WorkInter interface {
	Work()
}

type Student struct {
	Name string
	Age  uint8
	Job  string
}

func (t *Student) Work() {
	fmt.Println(t.Name + "的任务是" + t.Job)
}

func main() {
	var t WorkInter = &Student{
		Name: "大一学生",
		Age:  18,
		Job:  "好好学习",
	}
	t.Work()
}

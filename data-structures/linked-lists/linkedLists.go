package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type student struct {
	name string
	age  int
	ssn  string
	next *student
}

func reverse(curr *student) *student {
	if curr.next == nil {
		return curr.next
	}
	temp := reverse(curr.next)
	temp.next.next = temp
	temp.next = nil
	return temp
}

func run() {
	students := new(student)

	students.next = nil

	if len(os.Args) < 2 {
		fmt.Println("there is no file")
		return
	}

	filename := os.Args[1]

	data, _ := ioutil.ReadFile(filename)

	for _, line := range strings.Split(string(data), "\n") {

		if line == "" {
			continue
		}

		td := strings.Split(string(line), " ")

		ts := &student{
			name: td[0],
			ssn:  td[2],
			next: students.next,
		}

		ts.age, _ = strconv.Atoi(td[1])

		students.next = ts

	}

	for s := students.next; s != nil; s = s.next {

		fmt.Println(s.name, s.age, s.ssn)

	}
}

func main() {
	newList := LinkedList{nil, nil}
	newList.Insert(5)
	newList.Insert(6)
	newList.Insert(7)

	newList.Print()
}

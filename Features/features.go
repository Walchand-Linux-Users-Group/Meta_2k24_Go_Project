package features

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

type Student struct {
	Name   string
	RollNo string
	Marks  int
}

type Class struct {
	Students []Student
}

func (c *Class) NewClass() {
	// var studentsArray []Student
	// fmt.Println("Enter your first name")
	// var name string
	// fmt.Scanln(&name)

	// fmt.Println("Enter your Roll No")
	// var roll string
	// fmt.Scanln(&roll)

	// fmt.Println("Enter your marks")
	// var marks int
	// fmt.Scanln(&marks)

	// studentsArray = append(studentsArray, Student{Name: name, RollNo: roll, Marks: marks,})

	// // fmt.Println(studentsArray)
	// c.Students=studentsArray

	var studentsArray []Student
	file, err := os.Open("db.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		marks, _ := strconv.Atoi(fields[2])
		studentsArray = append(studentsArray, Student{Name: fields[0], RollNo: fields[1], Marks: marks})
	}
	// fmt.Println(studentsArray)

	c.Students = studentsArray

}

func (c *Class) AddStudent(){
	fmt.Println("Enter your first name")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter your Roll No")
	var roll string
	fmt.Scanln(&roll)

	fmt.Println("Enter your marks")
	var marks int
	fmt.Scanln(&marks)

	entry:=Student{
		Name: name,
		RollNo: roll,
		Marks: marks,
	}

	c.Students=append(c.Students, entry)

	//Updata entries in db.txt

	fmt.Println("Data Added Successfully")
	fmt.Println()
}

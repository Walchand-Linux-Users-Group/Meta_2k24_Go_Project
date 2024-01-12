package features

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

type Student struct {
	Name string
	RollNo string
	Marks int
}

type Class struct {
	Students []Student
}

func (c *Class) NewClass(){

	

	// fmt.Println(entry)
	var studentsArray []Student
	file,err:=os.Open("db.txt")
	if err !=nil {
		panic(err)
	}
	scanner:=bufio.NewScanner(file)

	for scanner.Scan() {
		line:=scanner.Text()
		fields:=strings.Split(line," ")
		marks,_:=strconv.Atoi(fields[2])

		entry:=Student{
			Name: fields[0],
			RollNo: fields[1],
			Marks: marks,
		}

		studentsArray=append(studentsArray, entry)

	}

	// fmt.Println(studentsArray)
	c.Students=studentsArray

}

// CRUD: Create Read Update and Delete

func (c *Class) AddStudent(){

	fmt.Println("Enter your first name")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter your roll no")
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

	fmt.Println("Student Added Successfully")
	fmt.Println()

	//Update this data in the db.txt
}

func (c *Class) ShowStudents(){
	for _,student:=range(c.Students){
		name,roll,marks:=student.Name,student.RollNo,student.Marks
		fmt.Printf("Name: %s\nRoll No: %s\nMarks: %d\n",name,roll,marks)
		fmt.Println()
	}
}

func (c *Class) UpdateStudent(){
	var roll string
	fmt.Println("Enter the roll no")
	fmt.Scanln(&roll)
	var i int
	for idx,student:=range(c.Students) {
		if(student.RollNo==roll){
			i=idx
			break
		}
	}

	fmt.Println("Enter your first name")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter your roll no")
	var rollNo string
	fmt.Scanln(&rollNo)

	fmt.Println("Enter your marks")
	var marks int
	fmt.Scanln(&marks)

	c.Students[i]=Student{
		Name: name,
		RollNo: rollNo,
		Marks: marks,
	}

	fmt.Println("Data updated successfully")
	fmt.Println()

	//Update it in db.txt
}

func (c *Class) DeleteStudent(){
	var roll string
	fmt.Println("Enter the roll no")
	fmt.Scanln(&roll)
	var i int
	for idx,student:=range(c.Students) {
		if(student.RollNo==roll){
			i=idx
			break
		}
	}

	dummyArr:=c.Students

	dummyArr=append(dummyArr[:i],dummyArr[i+1:]... )
	c.Students=dummyArr
	fmt.Println("Data deleted successfully")
	fmt.Println()

	//Update the data in db.txt

}


package main

import (
	"fmt"
	"time"
)

type Grader interface {
	Grade() string
}

type Student struct {
	name  string
	marks int
}

type Result struct {
	name  string
	grade string
	marks int
}

func panicprocess(marks int, name string) {

	defer func() {
		r := recover()
		fmt.Println("recovered: ", r)
	}()

	panic(fmt.Sprintf("%d is an invalid marks for %s", marks, name))
}

func (student Student) Grade() string {

	result := ""
	switch m := student.marks; {
	case m >= 90 && m <= 100:
		result = "A"
	case m >= 75 && m <= 89:
		result = "B"
	case m >= 60 && m <= 74:
		result = "C"
	case m >= 40 && m <= 59:
		result = "D"
	case m < 40:
		result = "F"
	default:
		panicprocess(m, student.name)
	}

	return result
}

func processstudent(students []Student, resultchn chan Result) {

	for _, student := range students {
		grade := student.Grade()
		if grade != "" {
			resultchn <- Result{student.name, grade, student.marks}
		}
	}

}

func studentexamproccesssystem() {

	s1 := Student{"abhay", 60}
	s2 := Student{"vaibhav", 90}
	s3 := Student{"jay", 160}
	s4 := Student{"hero", 30}
	s5 := Student{"suraj", 95}

	students := []Student{s1, s2, s3, s4, s5}

	resultchn := make(chan Result, len(students))

	go processstudent(students[0:2], resultchn)
	go processstudent(students[2:3], resultchn)
	go processstudent(students[3:5], resultchn)

	time.Sleep(time.Second)
	close(resultchn)

	result := make(map[string]string)
	highscore := 0
	highscorer := ""
	sum := 0

	fmt.Println()
	for studentresult := range resultchn {
		fmt.Printf("%s -> %s\n", studentresult.name, studentresult.grade)
		result[studentresult.name] = studentresult.grade
		if studentresult.marks > highscore {
			highscore = studentresult.marks
			highscorer = studentresult.name
		}
		sum += studentresult.marks
	}

	avgmarks := sum / len(result)
	count := [6]int{}
	for _, grade := range result {
		count[rune(grade[0])-65]++
	}

	fmt.Printf("\nAverage Marks: %d\n", avgmarks)

	fmt.Println("\nHighest Scorer:")
	fmt.Printf("%s (%d)\n", highscorer, highscore)

	fmt.Println("\nGrade Counts:")
	for grade, freq := range count {
		fmt.Printf("%c: %d\n", grade+65, freq)
	}

}

package Student

import "fmt"

type Student struct {
	name  string
	score int
}

func GetStudent(name string, score int) *Student {
	return &Student{name, score}
}

func (s *Student) String() string {
	return fmt.Sprintf("Student(name: %s, score: %d)", s.name, s.score)
}

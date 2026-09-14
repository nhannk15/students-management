package student

type Student struct {
	Id        int
	Name      string
	ClassName string
	Grades    []int
}

func (student Student) CalculateAverageGrades() float32 {
	var totalGrades float32

	for _, grade := range student.Grades {
		totalGrades += float32(grade)
	}

	return totalGrades / float32(len(student.Grades))
}

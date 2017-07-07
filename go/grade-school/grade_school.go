package school

import "sort"

const testVersion = 1

type Grade struct {
	grade    int
	students []string
}

type School map[int]*Grade

func New() *School {
	school := make(School, 0)
	return &school
}

func (s *School) Add(name string, grade int) {
	if 0 < grade && grade < 10 {
		if (*s)[grade] == nil {
			(*s)[grade] = &Grade{grade, make([]string, 0)}
		}
		(*s)[grade].students = append((*s)[grade].students, name)
		sort.StringSlice((*s)[grade].students).Sort()
		(*s)[grade].grade = grade
	}
}

func (s *School) Grade(g int) []string {
	if (*s)[g] == nil {
		return []string{}
	}
	sort.StringSlice((*s)[g].students).Sort()
	return (*s)[g].students
}

type GradeSlice []Grade

func (g GradeSlice) Len() int           { return len(g) }
func (g GradeSlice) Less(i, j int) bool { return g[i].grade < g[j].grade }
func (g GradeSlice) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }

func (s *School) Enrollment() []Grade {
	grades := make([]Grade, 0)
	for _, g := range *s {
		grades = append(grades, *g)
	}
	sort.Sort(GradeSlice(grades))
	return grades
}

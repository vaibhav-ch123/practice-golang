package employe

type employe struct {
	name    string
	salary  int
	jobpost string
}

func New(name string, salary int, jobpost string) employe {
	e := employe{name, salary, jobpost}
	return e
}

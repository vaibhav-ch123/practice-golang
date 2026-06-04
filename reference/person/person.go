package person

type person struct {
	name  string
	age   int
	hobby string
}

func New(name string, age int, hobby string) person {

	p := person{name, age, hobby}
	return p
}

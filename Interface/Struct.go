package main

type cat struct {
	weight float64
}

func (c cat) howWeight() float64 {
	return c.weight
}

func (c cat) feedShareMass() float64 {
	return c.weight * 7
}

func (c cat) String() string {
	return "Cat"
}

type dog struct {
	weight float64
}

func (d dog) howWeight() float64 {
	return d.weight
}

func (d dog) feedShareMass() float64 {
	return d.weight / 5.0 * 10.0
}

func (d dog) String() string {
	return "Dog"
}

type cow struct {
	weight float64
}

func (c cow) howWeight() float64 {
	return c.weight
}

func (c cow) feedShareMass() float64 {
	return c.weight * 25
}

func (c cow) String() string {
	return "Cow"
}

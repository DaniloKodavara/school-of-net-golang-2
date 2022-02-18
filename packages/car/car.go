package car

type Car struct {
	Name  string
	color string
}

func (c Car) Start() string {
	c.color = "Black"
	return c.Name + " color " + c.color + " has been started"
}

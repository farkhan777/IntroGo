package main

import "fmt"

type vehicle struct {
	doors int
	colors string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main()  {
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			colors: "Yellow",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 2,
			colors: "Black",
		},
		luxury: true,
	}

	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Println(t1.colors, t1.doors, t1.fourWheel)
	fmt.Println(s1.doors, s1.colors, s1.luxury)
}

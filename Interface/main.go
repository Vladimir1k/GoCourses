package main

import "fmt"

type howAnimalsEat interface {
	feedShareMass() float64
	howWeight() float64
	fmt.Stringer
}

func main() {
	animals := []howAnimalsEat{
		cat{weight: 4.5},
		cat{weight: 7},
		dog{weight: 9},
		dog{weight: 11},
		dog{weight: 15},
		cow{weight: 300},
		cow{weight: 150},
	}
	allFeed := 0.0
	for _, a := range animals {
		fmt.Printf("\n Тварина %v має вагу %vкг, та потребує %vкг корму на місяць",
			howAnimalsEat(a), a.howWeight(), a.feedShareMass())
		//fmt.Println(a.String(), a.feedShareMass())
		allFeed += a.feedShareMass()
	}
	fmt.Printf("\nВсього потрібно %vкг корму на місяць", allFeed)
}

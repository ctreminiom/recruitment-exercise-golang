package main

import (
	"github.com/ctreminiom/recruitment-exercise-golang/factory"
)

const carsAmount = 100

func main() {
	assemblesFactory := factory.New()

	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id
	assemblesFactory.StartAssemblingProcess(carsAmount)
}

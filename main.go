package main

import (
	"fmt"
	"github.com/ctreminiom/recruitment-exercise-golang/factory"
)

const carsAmount = 100

func main() {
	assemblesFactory := factory.New()

	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id

	//////////// Possible Solution //////////////////////
	// 1. I've implemented a buffered channel using a custom struct created under the factory package
	// 2. The structure VehicleLoggerScheme contains the information required on this "Hint" comment method
	// 3. The StartAssemblingProcess returns the status of each iteration and it stores the metadata on the channel
	// 4. Runs the StartAssemblingProcess method with the channels concurrently.
	// 5. Execute for loop with the select clause in order to print the status of each vehicle integration.
	// TODO: Format the message printed.
	out := make(chan *factory.VehicleLoggerScheme, carsAmount)
	go assemblesFactory.StartAssemblingProcess(carsAmount, out)

	for {
		select {
		default:
			fmt.Println(<-out)
		}
	}
}

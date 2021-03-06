package main

import (
	"github.com/ctreminiom/recruitment-exercise-golang/factory"
	log "github.com/sirupsen/logrus"
)

const carsAmount = 100

func main() {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	assemblesFactory := factory.New()

	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id

	//////////// Possible Solution //////////////////////
	// 1. I've implemented a buffered channel using a custom struct created under the factory package
	// 2. The structure VehicleLoggerScheme contains the information required on this "Hint" comment method
	// 3. The StartAssemblingProcess returns the status of each iteration and it stores the metadata on the channel
	// 4. Runs the StartAssemblingProcess method with the channels concurrently.
	// 5. Execute for loop with the select clause in order to print the status of each vehicle integration.
	listener := make(chan *factory.VehicleLoggerScheme)
	chError := make(chan error)
	isCompleted := make(chan bool)

	go assemblesFactory.StartAssemblingProcess(carsAmount, listener, chError, isCompleted)

	for {
		select {

		case event := <-listener:
			log.WithFields(log.Fields{
				"vehicle-id":    event.ID,
				"testing-logs":  event.History,
				"assemble-logs": event.AssemblyStatus,
			}).Info("Vehicle Assembled :)")

		case appError := <-chError:
			panic(appError)

		case <-isCompleted:
			return
		}
	}
}

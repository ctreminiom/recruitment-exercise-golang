package factory

import (
	"fmt"
	"github.com/ctreminiom/recruitment-exercise-golang/assemblyspot"
	"github.com/ctreminiom/recruitment-exercise-golang/vehicle"
	"github.com/gammazero/workerpool"
	"log"
)

const assemblySpots int = 5

type Factory struct {
	AssemblingSpots chan *assemblyspot.AssemblySpot
}

func New() *Factory {
	factory := &Factory{
		AssemblingSpots: make(chan *assemblyspot.AssemblySpot, assemblySpots),
	}

	totalAssemblySpots := 0

	for {
		factory.AssemblingSpots <- &assemblyspot.AssemblySpot{}

		totalAssemblySpots++

		if totalAssemblySpots >= assemblySpots {
			break
		}
	}

	return factory
}

type VehicleLoggerScheme struct {
	ID             int
	History        string
	AssemblyStatus string
}

// StartAssemblingProcess
//HINT: this function is currently not returning anything, make it return right away every single vehicle once assembled,
//(Do not wait for all of them to be assembled to return them all, send each one ready over to main)
/////// Possible Solution ////////
// As the AssembleVehicle method can been the internal method concurrently using mutex's and waitgroups
// The method needs 7 seconds to process the vehicle, we can reduce the execution time using a workerpool
// I tried to create a out-the-box workerpool (documented under the worker method), but I as could not implement it at time,
// I used a third-party called workerpool ("github.com/gammazero/workerpool") and now the program takes 7 seconds to process 5 vehicles.
func (f *Factory) StartAssemblingProcess(amountOfVehicles int, out chan<- *VehicleLoggerScheme) {

	var (
		vehicleList = f.generateVehicleLots(amountOfVehicles)
		worker      = workerpool.New(assemblySpots)
	)

	for _, vehicle := range vehicleList {

		vehicle := vehicle

		worker.Submit(func() {
			fmt.Println("Assembling vehicle...")

			idleSpot := <-f.AssemblingSpots
			idleSpot.SetVehicle(&vehicle)
			vehicle, err := idleSpot.AssembleVehicle()

			if err != nil {
				log.Println(err)
				return
			}

			vehicle.TestingLog = f.testCar(vehicle)
			vehicle.AssembleLog = idleSpot.GetAssembledLogs()

			idleSpot.SetVehicle(nil)
			f.AssemblingSpots <- idleSpot

			out <- &VehicleLoggerScheme{
				ID:             vehicle.Id,
				History:        vehicle.TestingLog,
				AssemblyStatus: vehicle.AssembleLog,
			}
		})

	}

	worker.StopWait()
}

func worker(cars []vehicle.Car, spots <-chan *assemblyspot.AssemblySpot, factory *Factory, result chan<- *VehicleLoggerScheme) {

	for spot := range spots {

		for _, car := range cars {
			fmt.Println("Assembling vehicle...")

			spot.SetVehicle(&car)
			vehicle, err := spot.AssembleVehicle()

			if err != nil {
				continue
			}

			vehicle.TestingLog = factory.testCar(vehicle)
			vehicle.AssembleLog = spot.GetAssembledLogs()

			result <- &VehicleLoggerScheme{
				ID:             vehicle.Id,
				History:        vehicle.TestingLog,
				AssemblyStatus: vehicle.AssembleLog,
			}
		}
	}
}

func (Factory) generateVehicleLots(amountOfVehicles int) []vehicle.Car {
	var vehicles []vehicle.Car
	var index = 0

	for {
		vehicles = append(vehicles, vehicle.Car{
			Id:            index,
			Chassis:       "NotSet",
			Tires:         "NotSet",
			Engine:        "NotSet",
			Electronics:   "NotSet",
			Dash:          "NotSet",
			Sits:          "NotSet",
			Windows:       "NotSet",
			EngineStarted: false,
		})

		index++

		if index >= amountOfVehicles {
			break
		}
	}

	return vehicles
}

func (f *Factory) testCar(car *vehicle.Car) string {
	logs := ""

	log, err := car.StartEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnLeft()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnRight()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.StopEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	return logs
}

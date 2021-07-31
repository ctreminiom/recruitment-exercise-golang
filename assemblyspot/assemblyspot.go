package assemblyspot

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ctreminiom/recruitment-exercise-golang/vehicle"
)

type AssemblySpot struct {
	vehicleToAssemble *vehicle.Car
	assemblyLog       string
	mutex             sync.Mutex
}

func (s *AssemblySpot) SetVehicle(v *vehicle.Car) {
	s.vehicleToAssemble = v
}

func (s *AssemblySpot) GetAssembledVehicle() *vehicle.Car {
	return s.vehicleToAssemble

}

func (s *AssemblySpot) GetAssembledLogs() string {
	return s.assemblyLog
}

type AssembleVehicleResult struct {
	Vehicle *vehicle.Car
	Err     error
}

//hint: improve this function to execute this process concurrenlty
func (s *AssemblySpot) AssembleVehicle(result chan<- *vehicle.Car, chErr chan<- error) {

	if s.vehicleToAssemble == nil {
		chErr <- errors.New("no vehicle set to start assembling")
	}

	wg := sync.WaitGroup{}
	wg.Add(7)

	go s.assembleChassis(&wg)
	go s.assembleTires(&wg)
	go s.assembleEngine(&wg)
	go s.assembleElectronics(&wg)
	go s.assembleDash(&wg)
	go s.assembleSeats(&wg)
	go s.assembleWindows(&wg)

	wg.Wait()

	result <- s.vehicleToAssemble

	//return s.vehicleToAssemble, nil
}

func (s *AssemblySpot) assembleChassis(wg *sync.WaitGroup) {
	s.mutex.Lock()
	s.vehicleToAssemble.Chassis = "Assembled"
	s.mutex.Unlock()

	time.Sleep(1 * time.Second)

	s.mutex.Lock()
	s.assemblyLog += fmt.Sprintf("Chassis at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	s.mutex.Unlock()

	wg.Done()
}

func (s *AssemblySpot) assembleTires(wg *sync.WaitGroup) {
	s.mutex.Lock()
	s.vehicleToAssemble.Tires = "Assembled"
	s.mutex.Unlock()

	time.Sleep(1 * time.Second)

	s.mutex.Lock()
	s.assemblyLog += fmt.Sprintf("Tires at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	s.mutex.Unlock()

	wg.Done()
}

func (s *AssemblySpot) assembleEngine(wg *sync.WaitGroup) {
	s.mutex.Lock()
	s.vehicleToAssemble.Engine = "Assembled"
	s.mutex.Unlock()

	time.Sleep(1 * time.Second)

	s.mutex.Lock()
	s.assemblyLog += fmt.Sprintf("Engine at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	s.mutex.Unlock()

	wg.Done()
}

func (s *AssemblySpot) assembleElectronics(wg *sync.WaitGroup) {
	s.mutex.Lock()
	s.vehicleToAssemble.Electronics = "Assembled"
	s.mutex.Unlock()

	// Simulate work
	time.Sleep(1 * time.Second)

	s.mutex.Lock()
	s.assemblyLog += fmt.Sprintf("Electronics at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	s.mutex.Unlock()

	wg.Done()
}

func (s *AssemblySpot) assembleDash(wg *sync.WaitGroup) {
	s.mutex.Lock()
	s.vehicleToAssemble.Dash = "Assembled"
	s.mutex.Unlock()

	// Simulate work
	time.Sleep(1 * time.Second)

	s.mutex.Lock()
	s.assemblyLog += fmt.Sprintf("Dash at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	s.mutex.Unlock()

	wg.Done()
}

func (s *AssemblySpot) assembleSeats(wg *sync.WaitGroup) {
	s.mutex.Lock()
	s.vehicleToAssemble.Sits = "Assembled"
	s.mutex.Unlock()

	// Simulate work
	time.Sleep(1 * time.Second)

	s.mutex.Lock()
	s.assemblyLog += fmt.Sprintf("Sits at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	s.mutex.Unlock()

	wg.Done()
}

func (s *AssemblySpot) assembleWindows(wg *sync.WaitGroup) {
	s.mutex.Lock()
	s.vehicleToAssemble.Windows = "Assembled"
	s.mutex.Unlock()

	// Simulate work
	time.Sleep(1 * time.Second)

	s.mutex.Lock()
	s.assemblyLog += fmt.Sprintf("Windows at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	s.mutex.Unlock()
	wg.Done()
}

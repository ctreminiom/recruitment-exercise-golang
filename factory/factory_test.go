package factory

import (
	"github.com/ctreminiom/recruitment-exercise-golang/vehicle"
	"reflect"
	"testing"
)

func TestFactory_generateVehicleLots(t *testing.T) {

	f := New()
	var vehiclesMockedWithValues = f.generateVehicleLots(100)

	type args struct {
		amountOfVehicles int
	}

	tests := []struct {
		name    string
		args    args
		want    []vehicle.Car
		wantErr bool
	}{
		{
			name: "generateVehicleLotsWhenTheParameterAreCorrect",
			args: args{
				amountOfVehicles: 100,
			},
			want:    vehiclesMockedWithValues,
			wantErr: false,
		},

		{
			name: "generateVehicleLotsWhenTheVehiclesCreatedIsIncorrect",
			args: args{
				amountOfVehicles: 50,
			},
			want:    vehiclesMockedWithValues,
			wantErr: true,
		},

		{
			name: "generateVehicleLotsWhenTheVehiclesCreatedIsDoesNotContainsValues",
			args: args{
				amountOfVehicles: 0,
			},
			want:    vehiclesMockedWithValues,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := New()

			if tt.wantErr {

				if got := f.generateVehicleLots(tt.args.amountOfVehicles); reflect.DeepEqual(got, tt.want) {
					t.Errorf("generateVehicleLots() = %v, want %v", got, tt.want)
				}

			} else {
				if got := f.generateVehicleLots(tt.args.amountOfVehicles); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("generateVehicleLots() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestFactory_testCar(t *testing.T) {

	type args struct {
		car *vehicle.Car
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testCarWhenTheParametersAreCorrect",
			args: args{
				car: &vehicle.Car{
					Id:            20,
					Chassis:       "",
					Tires:         "",
					Engine:        "",
					Electronics:   "",
					Dash:          "",
					Sits:          "",
					Windows:       "",
					EngineStarted: false,
					TestingLog:    "",
					AssembleLog:   "",
				},
			},
			want: "Engine Started!, Cannot move with stopped engine, Cannot move with stopped engine, Cannot turn left " +
				"with stopped engine, Cannot turn right with stopped engine, Cannot stop engine already stopped, ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := New()

			idleSpot := <-f.AssemblingSpots
			idleSpot.SetVehicle(tt.args.car)

			result := make(chan *vehicle.Car)
			chError := make(chan error)

			go idleSpot.AssembleVehicle(result, chError)

			select {

			case vehicle := <-result:
				vehicle.TestingLog = f.testCar(vehicle)
				vehicle.AssembleLog = idleSpot.GetAssembledLogs()
				idleSpot.SetVehicle(nil)
				f.AssemblingSpots <- idleSpot

				if got := f.testCar(vehicle); got != tt.want {
					t.Errorf("testCar() = %v, want %v", got, tt.want)
				}

			case assembleVehicleError := <-chError:
				t.Error(assembleVehicleError)
			}

		})
	}
}

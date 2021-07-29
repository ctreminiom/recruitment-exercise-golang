package vehicle

import (
	"testing"
)

func TestCar_MoveBackwards(t *testing.T) {

	type fields struct {
		Id            int
		Chassis       string
		Tires         string
		Engine        string
		Electronics   string
		Dash          string
		Sits          string
		Windows       string
		EngineStarted bool
		TestingLog    string
		AssembleLog   string
	}
	type args struct {
		distance int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "MoveBackwardsWhenTheParametersAreCorrect",
			fields: fields{
				Id:            1,
				EngineStarted: true,
			},
			args: args{
				distance: 20,
			},
			want:    "Moved backwards 20 meters!",
			wantErr: false,
		},

		{
			name: "MoveBackwardsWhenTheEngineIsNotStarted",
			fields: fields{
				Id:            1,
				EngineStarted: false,
			},
			args: args{
				distance: 20,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := &Car{
				Id:            tt.fields.Id,
				Chassis:       tt.fields.Chassis,
				Tires:         tt.fields.Tires,
				Engine:        tt.fields.Engine,
				Electronics:   tt.fields.Electronics,
				Dash:          tt.fields.Dash,
				Sits:          tt.fields.Sits,
				Windows:       tt.fields.Windows,
				EngineStarted: tt.fields.EngineStarted,
				TestingLog:    tt.fields.TestingLog,
				AssembleLog:   tt.fields.AssembleLog,
			}

			got, err := c.MoveBackwards(tt.args.distance)

			if (err != nil) != tt.wantErr {
				t.Errorf("MoveBackwards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("MoveBackwards() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCar_MoveForwards(t *testing.T) {
	type fields struct {
		Id            int
		Chassis       string
		Tires         string
		Engine        string
		Electronics   string
		Dash          string
		Sits          string
		Windows       string
		EngineStarted bool
		TestingLog    string
		AssembleLog   string
	}
	type args struct {
		distance int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "MoveForwardsWhenTheParametersAreCorrect",
			fields: fields{
				Id:            1,
				EngineStarted: true,
			},
			args: args{
				distance: 20,
			},
			want:    "Moved forward 20 meters!",
			wantErr: false,
		},

		{
			name: "MoveForwardsWhenTheEngineIsNotStarted",
			fields: fields{
				Id:            1,
				EngineStarted: false,
			},
			args: args{
				distance: 20,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Car{
				Id:            tt.fields.Id,
				Chassis:       tt.fields.Chassis,
				Tires:         tt.fields.Tires,
				Engine:        tt.fields.Engine,
				Electronics:   tt.fields.Electronics,
				Dash:          tt.fields.Dash,
				Sits:          tt.fields.Sits,
				Windows:       tt.fields.Windows,
				EngineStarted: tt.fields.EngineStarted,
				TestingLog:    tt.fields.TestingLog,
				AssembleLog:   tt.fields.AssembleLog,
			}
			got, err := c.MoveForwards(tt.args.distance)
			if (err != nil) != tt.wantErr {
				t.Errorf("MoveForwards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MoveForwards() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCar_StartEngine(t *testing.T) {
	type fields struct {
		Id            int
		Chassis       string
		Tires         string
		Engine        string
		Electronics   string
		Dash          string
		Sits          string
		Windows       string
		EngineStarted bool
		TestingLog    string
		AssembleLog   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "StartEngineWhenTheParametersAreCorrect",
			fields: fields{
				Id:            1,
				EngineStarted: false,
			},
			want:    "Engine Started!",
			wantErr: false,
		},

		{
			name: "StartEngineWhenTheEngineIsAlreadyStarted",
			fields: fields{
				Id:            1,
				EngineStarted: true,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Car{
				Id:            tt.fields.Id,
				Chassis:       tt.fields.Chassis,
				Tires:         tt.fields.Tires,
				Engine:        tt.fields.Engine,
				Electronics:   tt.fields.Electronics,
				Dash:          tt.fields.Dash,
				Sits:          tt.fields.Sits,
				Windows:       tt.fields.Windows,
				EngineStarted: tt.fields.EngineStarted,
				TestingLog:    tt.fields.TestingLog,
				AssembleLog:   tt.fields.AssembleLog,
			}
			got, err := c.StartEngine()
			if (err != nil) != tt.wantErr {
				t.Errorf("StartEngine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StartEngine() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCar_StopEngine(t *testing.T) {
	type fields struct {
		Id            int
		Chassis       string
		Tires         string
		Engine        string
		Electronics   string
		Dash          string
		Sits          string
		Windows       string
		EngineStarted bool
		TestingLog    string
		AssembleLog   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "StopEngineWhenTheParametersAreCorrect",
			fields: fields{
				Id:            1,
				EngineStarted: true,
			},
			want:    "Engine Stopped!",
			wantErr: false,
		},

		{
			name: "StopEngineWhenTheEngineIsAlreadyStopped",
			fields: fields{
				Id:            1,
				EngineStarted: false,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Car{
				Id:            tt.fields.Id,
				Chassis:       tt.fields.Chassis,
				Tires:         tt.fields.Tires,
				Engine:        tt.fields.Engine,
				Electronics:   tt.fields.Electronics,
				Dash:          tt.fields.Dash,
				Sits:          tt.fields.Sits,
				Windows:       tt.fields.Windows,
				EngineStarted: tt.fields.EngineStarted,
				TestingLog:    tt.fields.TestingLog,
				AssembleLog:   tt.fields.AssembleLog,
			}
			got, err := c.StopEngine()
			if (err != nil) != tt.wantErr {
				t.Errorf("StopEngine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StopEngine() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCar_TurnLeft(t *testing.T) {
	type fields struct {
		Id            int
		Chassis       string
		Tires         string
		Engine        string
		Electronics   string
		Dash          string
		Sits          string
		Windows       string
		EngineStarted bool
		TestingLog    string
		AssembleLog   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "TurnLeftWhenTheParametersAreCorrect",
			fields: fields{
				Id:            1,
				EngineStarted: true,
			},
			want:    "Turned Left",
			wantErr: false,
		},

		{
			name: "TurnLeftWhenTheEngineIsNotStarted",
			fields: fields{
				Id:            1,
				EngineStarted: false,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Car{
				Id:            tt.fields.Id,
				Chassis:       tt.fields.Chassis,
				Tires:         tt.fields.Tires,
				Engine:        tt.fields.Engine,
				Electronics:   tt.fields.Electronics,
				Dash:          tt.fields.Dash,
				Sits:          tt.fields.Sits,
				Windows:       tt.fields.Windows,
				EngineStarted: tt.fields.EngineStarted,
				TestingLog:    tt.fields.TestingLog,
				AssembleLog:   tt.fields.AssembleLog,
			}
			got, err := c.TurnLeft()
			if (err != nil) != tt.wantErr {
				t.Errorf("TurnLeft() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TurnLeft() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCar_TurnRight(t *testing.T) {
	type fields struct {
		Id            int
		Chassis       string
		Tires         string
		Engine        string
		Electronics   string
		Dash          string
		Sits          string
		Windows       string
		EngineStarted bool
		TestingLog    string
		AssembleLog   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "TurnRightWhenTheParametersAreCorrect",
			fields: fields{
				Id:            1,
				EngineStarted: true,
			},
			want:    "Turned Right!",
			wantErr: false,
		},

		{
			name: "TurnRightWhenTheEngineIsNotStated",
			fields: fields{
				Id:            1,
				EngineStarted: false,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Car{
				Id:            tt.fields.Id,
				Chassis:       tt.fields.Chassis,
				Tires:         tt.fields.Tires,
				Engine:        tt.fields.Engine,
				Electronics:   tt.fields.Electronics,
				Dash:          tt.fields.Dash,
				Sits:          tt.fields.Sits,
				Windows:       tt.fields.Windows,
				EngineStarted: tt.fields.EngineStarted,
				TestingLog:    tt.fields.TestingLog,
				AssembleLog:   tt.fields.AssembleLog,
			}
			got, err := c.TurnRight()
			if (err != nil) != tt.wantErr {
				t.Errorf("TurnRight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TurnRight() got = %v, want %v", got, tt.want)
			}
		})
	}
}

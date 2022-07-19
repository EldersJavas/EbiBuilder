package tool

import (
	"testing"
)

func TestIsEbitenGame(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{name: "1", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := IsEbitenGame(); got != tt.want {
				t.Errorf("IsEbitenGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckGo(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"TestCheckGo", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckGo(); got != tt.want {
				t.Errorf("CheckGo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDebugPrint(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestDebugPrint", args{"TestDebugPrint"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DebugPrint(tt.args.s)
		})
	}
}

func TestErrorPrint(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestErrorPrint", args{"TestErrorPrint"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorPrint(tt.args.s)
		})
	}
}

func TestExecName(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestExecName", args{"TestExecName"}, "TestExecName.exe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecName(tt.args.s); got != tt.want {
				t.Errorf("ExecName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEbitenVer(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEbitenVer()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEbitenVer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetEbitenVer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfoPrint(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestInfoPrint", args{"TestInfoPrint"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InfoPrint(tt.args.s)
		})
	}
}

func TestStepPrint(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestStepPrint", args{"TestStepPrint"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StepPrint(tt.args.s)
		})
	}
}

func TestSuccessPrint(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestSuccessPrint", args{"TestSuccessPrint"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SuccessPrint(tt.args.s)
		})
	}
}

func TestWarnPrint(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestWarnPrint", args{"TestWarnPrint"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WarnPrint(tt.args.s)
		})
	}
}

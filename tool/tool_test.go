package tool

import (
	"testing"
)

func TestIsEbitenGame(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		struct{name string; want bool}{name: "1",want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEbitenGame(); got != tt.want {
				t.Errorf("IsEbitenGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

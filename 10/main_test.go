package main

import (
	"math"
	"testing"
)

func Test_angleBetween(t *testing.T) {
	type args struct {
		asteroid1 *asteroid
		asteroid2 *asteroid
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"above", args{&asteroid{x: 5, y: 5}, &asteroid{x: 5, y: 0}}, 0},
		{"below", args{&asteroid{x: 5, y: 5}, &asteroid{x: 5, y: 10}}, math.Pi},
		{"right", args{&asteroid{x: 5, y: 5}, &asteroid{x: 10, y: 5}}, math.Pi / 2},
		{"left", args{&asteroid{x: 5, y: 5}, &asteroid{x: 0, y: 5}}, 3 * math.Pi / 2},
		{"quadrant1", args{&asteroid{x: 5, y: 5}, &asteroid{x: 10, y: 0}}, math.Pi / 4},
		{"quadrant2", args{&asteroid{x: 5, y: 5}, &asteroid{x: 10, y: 10}}, 3 * math.Pi / 4},
		{"quadrant3", args{&asteroid{x: 5, y: 5}, &asteroid{x: 0, y: 10}}, 5 * math.Pi / 4},
		{"quadrant4", args{&asteroid{x: 5, y: 5}, &asteroid{x: 0, y: 0}}, 7 * math.Pi / 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := angleBetween(tt.args.asteroid1, tt.args.asteroid2); got != tt.want {
				t.Errorf("angleBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

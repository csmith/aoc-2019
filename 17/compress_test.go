package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_replace(t *testing.T) {
	type args struct {
		parts    []string
		function []string
		name     string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"replace entire set", args{[]string{"R", "1", "L", "2"}, []string{"R", "1", "L", "2"}, "A"}, []string{"A"}},
		{"replace prefix", args{[]string{"R", "1", "L", "2"}, []string{"R", "1", "L"}, "A"}, []string{"A", "2"}},
		{"replace middle", args{[]string{"R", "1", "L", "2"}, []string{"1", "L"}, "A"}, []string{"R", "A", "2"}},
		{"replace suffix", args{[]string{"R", "1", "L", "2"}, []string{"L", "2"}, "A"}, []string{"R", "1", "A"}},
		{"replace multi", args{[]string{"R", "1", "L", "R", "1", "2", "R", "1"}, []string{"R", "1"}, "A"}, []string{"A", "L", "A", "2", "A"}},
		{"replace adjacent", args{[]string{"R", "1", "R", "1"}, []string{"R", "1"}, "A"}, []string{"A", "A"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replace(tt.args.parts, tt.args.function, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compress(t *testing.T) {
	tests := []struct {
		name     string
		args     string
		wantMain string
		wantA    string
		wantB    string
		wantC    string
	}{
		{
			name:     "simple",
			args:     "R,L,1",
			wantMain: "A,B,C",
			wantA:    "R",
			wantB:    "L",
			wantC:    "1",
		},
		{
			name:     "reddit1",
			args:     "L,12,R,4,R,4,L,6,L,12,R,4,R,4,R,12,L,12,R,4,R,4,L,6,L,10,L,6,R,4,L,12,R,4,R,4,L,6,L,12,R,4,R,4,R,12,L,10,L,6,R,4,L,12,R,4,R,4,R,12,L,10,L,6,R,4,L,12,R,4,R,4,L,6",
			wantMain: "A,B,A,C,A,B,C,B,C,A",
			wantA:    "L,12,R,4,R,4,L,6",
			wantB:    "L,12,R,4,R,4,R,12",
			wantC:    "L,10,L,6,R,4",
		},
		{
			name:     "reddit2",
			args:     "L,6,R,12,L,4,L,6,R,6,L,6,R,12,R,6,L,6,R,12,L,6,L,10,L,10,R,6,L,6,R,12,L,4,L,6,R,6,L,6,R,12,L,6,L,10,L,10,R,6,L,6,R,12,L,4,L,6,R,6,L,6,R,12,L,6,L,10,L,10,R,6",
			wantMain: "A,B,B,C,A,B,C,A,B,C",
			wantA:    "L,6,R,12,L,4,L,6",
			wantB:    "R,6,L,6,R,12",
			wantC:    "L,6,L,10,L,10,R,6",
		},
		{
			name:     "reddit3",
			args:     "L,12,L,12,R,4,R,10,R,6,R,4,R,4,L,12,L,12,R,4,R,6,L,12,L,12,R,10,R,6,R,4,R,4,L,12,L,12,R,4,R,10,R,6,R,4,R,4,R,6,L,12,L,12,R,6,L,12,L,12,R,10,R,6,R,4,R,4",
			wantMain: "A,B,A,C,B,A,B,C,C,B",
			wantA:    "L,12,L,12,R,4",
			wantB:    "R,10,R,6,R,4,R,4",
			wantC:    "R,6,L,12,L,12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMain, gotA, gotB, gotC := compress(strings.Split(tt.args, ","))
			if gotMain != tt.wantMain {
				t.Errorf("compress() gotMain = %v, want %v", gotMain, tt.wantMain)
			}
			if gotA != tt.wantA {
				t.Errorf("compress() gotA = %v, want %v", gotA, tt.wantA)
			}
			if gotB != tt.wantB {
				t.Errorf("compress() gotB = %v, want %v", gotB, tt.wantB)
			}
			if gotC != tt.wantC {
				t.Errorf("compress() gotC = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func Test_prefixes(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want [][]string
	}{
		{"simple", []string{"1", "2", "3"}, [][]string{{"1", "2", "3"}, {"1", "2"}, {"1"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixes(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixes() = %v, want %v", got, tt.want)
			}
		})
	}
}

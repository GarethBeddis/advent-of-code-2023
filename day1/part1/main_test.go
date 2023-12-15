package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_processCalibrationDocument(t *testing.T) {
	type args struct {
		inputFile io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				inputFile: strings.NewReader(`a1b2c34f
5test6`),
			},
			want: 70,
		},
		{
			name: "2",
			args: args{
				inputFile: strings.NewReader(`testing10abc
10ten10`),
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processCalibrationDocument(tt.args.inputFile); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCalibrationValues(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				r: strings.NewReader(`a1b2c34f
5test6`),
			},
			want: []int{14, 56},
		},
		{
			name: "2",
			args: args{
				r: strings.NewReader(`testing10abc
10ten10`),
			},
			want: []int{10, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCalibrationValues(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCalibrationValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNumericCharacters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "test 1",
			input: "a1b2c3",
			want:  "123",
		},
		{
			name:  "test 2",
			input: "ab33bb22cc11",
			want:  "332211",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumericCharacters(tt.input); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumArray(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{

			name: "test 1",
			args: args{
				a: []int{1, 2, 3},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumArray(tt.args.a); got != tt.want {
				t.Errorf("sumArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

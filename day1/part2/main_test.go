package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_getFirstDigit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "start alphabetical digit",
			args: args{
				s: `nine223`,
			},
			want: "9",
		},
		{
			name: "enclosed alphabetical digit",
			args: args{
				s: `abfive1`,
			},
			want: "5",
		},
		{
			name: "end alphabetical digit",
			args: args{
				s: `abcdethree`,
			},
			want: "3",
		},
		{
			name: "numerical digit",
			args: args{
				s: `abc4six5two`,
			},
			want: "4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotD := getFirstDigit(tt.args.s); gotD != tt.want {
				t.Errorf("\n\ngot = %v, want %v\n\n", gotD, tt.want)
			}
		})
	}
}

func Test_startsWithAlphabeticalDigit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "9",
			args: args{
				s: "nine",
			},
			want: "nine",
		},
		{
			name: "5",
			args: args{
				s: "five",
			},
			want: "five",
		},
		{
			name: "1",
			args: args{
				s: "onetwothree",
			},
			want: "one",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := startsWithAlphabeticalDigit(tt.args.s)
			if got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_endsWithAlphabeticalDigit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "alphabetical - alphanumeric string",
			args: args{
				s: "55nine",
			},
			want: "nine",
		},
		{
			name: "alphabetical - multiple alphabetical digits",
			args: args{
				s: "1onetwothree23four",
			},
			want: "four",
		},
		{
			name: "last alphabetical digit in middle of string",
			args: args{
				s: "aoneb2c3fouraa",
			},
			want: "",
		},
		{
			name: "numerical",
			args: args{
				s: "aoneb2c3",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := endsWithAlphabeticalDigit(tt.args.s)
			if got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLastdigit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "end numeric digit",
			args: args{
				s: `nine223three3`,
			},
			want: "3",
		},
		{
			name: "enclosed alphabetical digit",
			args: args{
				s: `abfive1`,
			},
			want: "1",
		},
		{
			name: "end alphabetical digit",
			args: args{
				s: `abcde5three`,
			},
			want: "3",
		},
		{
			name: "numerical digit",
			args: args{
				s: `abc4six5two`,
			},
			want: "2",
		},
		{
			name: "numerical digit",
			args: args{
				s: `three2five1six4nine`,
			},
			want: "9",
		},
		{
			name: "last alphabetical digit in middle of string",
			args: args{
				s: `aoneb2c3fourf`,
			},
			want: "4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotD := getLastdigit(tt.args.s); gotD != tt.want {
				t.Errorf("got %v, want %v", gotD, tt.want)
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
				r: strings.NewReader(`aoneb2c3fourf
five1test2six`),
			},
			want: []int{14, 56},
		},
		{
			name: "2",
			args: args{
				r: strings.NewReader(`testingoneabonec
oneabcone`),
			},
			want: []int{11, 11},
		},
		{
			name: "3",
			args: args{
				r: strings.NewReader(`five666
twofoursixeight`),
			},
			want: []int{56, 28},
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
				inputFile: strings.NewReader(`two1nine
				eightwothree
				abcone2threexyz
				xtwone3four
				4nineeightseven2
				zoneight234
				7pqrstsixteen`),
			},
			want: 281,
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

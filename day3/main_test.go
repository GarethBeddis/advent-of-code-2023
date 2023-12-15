package main

import (
	_ "embed"
	"testing"
)

func Test_parseEngineSchematic(t *testing.T) {
	type args struct {
		schematic string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "aoc test input",
			args: args{
				schematic: `467..114..
			...*......
			..35..633.
			......#...
			617*......
			.....+.58.
			..592.....
			......755.
			...$.*....
			.664.598..`,
			},
			want: 4361,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := parseEngineSchematic(tt.args.schematic)

			if tt.want != sum {
				t.Errorf("getGameState() gotGameId = %v, want %v", sum, tt.want)
			}
		})
	}
}

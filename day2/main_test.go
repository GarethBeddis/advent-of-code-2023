package main

import (
	_ "embed"
	"testing"
)

func Test_getPossibleGameCount(t *testing.T) {
	type args struct {
		games string
		bag   map[string]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "aoc test input",
			args: args{
				games: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
				bag: map[string]int{
					"red":   12,
					"green": 13,
					"blue":  14,
				},
			},
			want: 8,
		},
		{
			name: "aoc test input",
			args: args{
				games: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
				bag: map[string]int{
					"red":   12,
					"green": 13,
					"blue":  14,
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := getPossibleGameCount(tt.args.games, tt.args.bag); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGameState(t *testing.T) {
	type args struct {
		game string
		bag  map[string]int
	}
	tests := []struct {
		name         string
		args         args
		wantGameId   int
		wantPossible bool
		wantPow      int
	}{
		{
			name: "aoc input 1",
			args: args{
				game: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green`,
				bag: map[string]int{
					"red":   12,
					"green": 13,
					"blue":  14,
				},
			},
			wantGameId:   1,
			wantPossible: true,
			wantPow:      4 * 6 * 2,
		},
		{
			name: "aoc input 2",
			args: args{
				game: `Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue`,
				bag: map[string]int{
					"red":   12,
					"green": 13,
					"blue":  14,
				},
			},
			wantGameId:   2,
			wantPossible: true,
			wantPow:      1 * 3 * 4,
		},
		{
			name: "aoc input 3",
			args: args{
				game: `Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red`,
				bag: map[string]int{
					"red":   12,
					"green": 13,
					"blue":  14,
				},
			},
			wantGameId:   3,
			wantPossible: false,
			wantPow:      20 * 13 * 6,
		},
		{
			name: "aoc input 3",
			args: args{
				game: `Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red`,
				bag: map[string]int{
					"red":   12,
					"green": 13,
					"blue":  14,
				},
			},
			wantGameId:   3,
			wantPossible: false,
			wantPow:      20 * 13 * 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGameId, gotPossible, gotPow := getGameState(tt.args.game, tt.args.bag)
			if gotGameId != tt.wantGameId {
				t.Errorf("getGameState() gotGameId = %v, want %v", gotGameId, tt.wantGameId)
			}
			if gotPossible != tt.wantPossible {
				t.Errorf("getGameState() gotPossible = %v, want %v", gotPossible, tt.wantPossible)
			}
			if gotPow != tt.wantPow {
				t.Errorf("getGameState() gotPow = %v, want %v", gotPow, tt.wantPow)
			}
		})
	}
}

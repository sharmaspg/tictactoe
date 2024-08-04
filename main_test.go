package main

import "testing"

func TestFindWinner(t *testing.T) {
	DiagonalWinner := [3][3]string{
		{"X", "-", "X"},
		{"O", "X", "O"},
		{"O", "O", "X"},
	}

	DiagonalWinner2 := [3][3]string{
		{"X", "-", "O"},
		{"X", "O", "X"},
		{"O", "X", "O"},
	}
	winner := FindWinner(DiagonalWinner)
	if winner != "X" {
		t.Errorf("Got %s , wanted 'X'", winner)
	}
	winner1 := FindWinner(DiagonalWinner2)
	if winner1 != "O" {
		t.Errorf("Got %s , wanted 'O'", winner1)

	}

}

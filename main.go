package main

import (
	"fmt"
	"math/rand"
)

var table = [3][3]string{
	{"-", "-", "-"},
	{"-", "-", "-"},
	{"-", "-", "-"},
}

func TakeInputPreference() string {
	var preference string
	fmt.Printf("Please select your Input: X or any Character: ")
	fmt.Scanln(&preference)
	if preference == "O" {
		fmt.Println("This is used by the computer. please choose another character.")
		TakeInputPreference()
	}
	fmt.Printf("You have selected %s. \n", preference)
	return preference
}
func addEntry(row, col int, prefer string) {
	table[row][col] = prefer

}
func TakeMatrixEntry(prefer string) {
	fmt.Printf("row then column , seperated by space. Ex:2x2 ")
	var row, col int
	fmt.Scan(&row, &col)
	if row > 3 || row < 1 {
		fmt.Println("Please enter valid number between 1-3")
		TakeMatrixEntry(prefer)
	}
	if col > 3 || col < 1 {
		fmt.Println("Please enter valid number between 1-3")
		TakeMatrixEntry(prefer)
	}

	if table[row-1][col-1] != "-" {
		fmt.Println("This cell is not empty. please enter again")
		TakeMatrixEntry(prefer)
	}
	addEntry(row-1, col-1, prefer)
}

func isMatrixFull() bool {
	for rowindex, row := range table {
		for colIndex, _ := range row {
			if table[rowindex][colIndex] == "-" {
				return false
			}
		}
	}
	return true
}
func main() {
	prefer := TakeInputPreference()
	fmt.Println("Where do you want to place (please refere this matix-\n[[1 1] [1 2] [1 3]\n[2 1] [2 2] [2 3]\n[3 1] [3 2] [3 3])")
	for !isMatrixFull() {
		TakeMatrixEntry(prefer)
		AddRandomEntry()
		printTable()
		winner := FindWinner(table)
		if winner != "-" {
			fmt.Printf("Found Winner : %s\n", winner)
			break
		}

	}

	// fmt.Println(prefer)
}

func printTable() {
	for rowIndex, row := range table {
		for colIndex, column := range row {
			if colIndex == 0 {
				fmt.Printf("%s |", column)
				continue
			}
			if colIndex == len(row)-1 {
				fmt.Printf(" %s", column)
				continue

			} else {
				fmt.Printf(" %s |", column)

			}

		}
		fmt.Println()
		if rowIndex < len(table)-1 {
			fmt.Println("---------")
		}

	}
}

func AddRandomEntry() {
	i := rand.Intn(3)
	j := rand.Intn(3)

	for table[i][j] != "-" {
		i = rand.Intn(3)
		j = rand.Intn(3)
		// AddRandomEntry()
	}
	fmt.Printf("Adding Entry in [%d:%d]\n", i, j)
	table[i][j] = "O"

}

func findDiagonalMatch(prefer string, table [3][3]string) bool {
	if table[0][0] == prefer && table[1][1] == prefer && table[2][2] == prefer {
		return true
	}
	if table[0][2] == prefer && table[1][1] == prefer && table[2][0] == prefer {
		return true
	}
	return false
}

func findVerticalMatch(row, col int, prefer string, table [3][3]string) bool {
	if row == 0 {
		if table[row][col] == prefer && table[row+1][col] == prefer && table[row+2][col] == prefer {
			return true
		}
	}
	if row == 1 {
		if table[row-1][col] == prefer && table[row][col] == prefer && table[row+1][col] == prefer {
			return true
		}
	}
	if row == 2 {
		if table[row-2][col] == prefer && table[row-1][col] == prefer && table[row][col] == prefer {
			return true
		}
	}
	return false
}

func findHorizontalMatch(row, col int, prefer string, table [3][3]string) bool {
	if col == 0 {
		if table[row][col] == prefer && table[row][col+1] == prefer && table[row][col+2] == prefer {
			return true
		}
	}
	if col == 1 {
		if table[row][col-1] == prefer && table[row][col] == prefer && table[row][col+1] == prefer {
			return true
		}
	}
	if col == 2 {
		if table[row][col-2] == prefer && table[row][col-1] == prefer && table[row][col] == prefer {
			return true
		}
	}
	return false
}

func FindWinner(table [3][3]string) string {
	for rowIndex, row := range table {
		for colIndex, cell := range row {
			if cell == "-" {
				continue
			}
			// fmt.Println("Finding match for ", cell)

			// fmt.Println(rowIndex, colIndex)
			if findVerticalMatch(rowIndex, colIndex, cell, table) || findHorizontalMatch(rowIndex, colIndex, cell, table) {
				return cell
			}
			if rowIndex == colIndex {
				if findDiagonalMatch(cell, table) {
					return cell
				}
			}
		}
	}
	return "-"
}

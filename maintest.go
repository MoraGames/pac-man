/*
package main

import "fmt"

func main() {
	layer := [31][28]rune{
		{'╔', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '╤', '╤', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '╗'},
		{'║', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '│', '│', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '║'},
		{'║', ' ', '┌', '─', '─', '┐', ' ', '┌', '─', '─', '─', '┐', ' ', '│', '│', ' ', '┌', '─', '─', '─', '┐', ' ', '┌', '─', '─', '┐', ' ', '║'},
		{'║', ' ', '│', ' ', ' ', '│', ' ', '│', ' ', ' ', ' ', '│', ' ', '│', '│', ' ', '│', ' ', ' ', ' ', '│', ' ', '│', ' ', ' ', '│', ' ', '║'},
		{'║', ' ', '└', '─', '─', '┘', ' ', '└', '─', '─', '─', '┘', ' ', '└', '┘', ' ', '└', '─', '─', '─', '┘', ' ', '└', '─', '─', '┘', ' ', '║'},
		{'║', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '║'},
		{'║', ' ', '┌', '─', '─', '┐', ' ', '┌', '┐', ' ', '┌', '─', '─', '─', '─', '─', '─', '┐', ' ', '┌', '┐', ' ', '┌', '─', '─', '┐', ' ', '║'},
		{'║', ' ', '└', '─', '─', '┘', ' ', '│', '│', ' ', '└', '─', '─', '┐', '┌', '─', '─', '┘', ' ', '│', '│', ' ', '└', '─', '─', '┘', ' ', '║'},
		{'║', ' ', ' ', ' ', ' ', ' ', ' ', '│', '│', ' ', ' ', ' ', ' ', '│', '│', ' ', ' ', ' ', ' ', '│', '│', ' ', ' ', ' ', ' ', ' ', ' ', '║'},
		{'╚', '═', '═', '═', '═', '╗', ' ', '│', '└', '─', '─', '┐', ' ', '│', '│', ' ', '┌', '─', '─', '┘', '│', ' ', '╔', '═', '═', '═', '═', '╝'},
		{' ', ' ', ' ', ' ', ' ', '║', ' ', '│', '┌', '─', '─', '┘', ' ', '└', '┘', ' ', '└', '─', '─', '┐', '│', ' ', '║', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', '║', ' ', '│', '│', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '│', '│', ' ', '║', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', '║', ' ', '│', '│', ' ', '╔', '═', '═', '─', '─', '═', '═', '╗', ' ', '│', '│', ' ', '║', ' ', ' ', ' ', ' ', ' '},
		{'═', '═', '═', '═', '═', '╝', ' ', '└', '┘', ' ', '║', ' ', ' ', ' ', ' ', ' ', ' ', '║', ' ', '└', '┘', ' ', '╚', '═', '═', '═', '═', '═'},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '║', ' ', ' ', ' ', ' ', ' ', ' ', '║', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{'═', '═', '═', '═', '═', '╗', ' ', '┌', '┐', ' ', '║', ' ', ' ', ' ', ' ', ' ', ' ', '║', ' ', '┌', '┐', ' ', '╔', '═', '═', '═', '═', '═'},
		{' ', ' ', ' ', ' ', ' ', '║', ' ', '│', '│', ' ', '╚', '═', '═', '═', '═', '═', '═', '╝', ' ', '│', '│', ' ', '║', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', '║', ' ', '│', '│', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '│', '│', ' ', '║', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', '║', ' ', '│', '│', ' ', '┌', '─', '─', '─', '─', '─', '─', '┐', ' ', '│', '│', ' ', '║', ' ', ' ', ' ', ' ', ' '},
		{'╔', '═', '═', '═', '═', '╝', ' ', '└', '┘', ' ', '└', '─', '─', '┐', '┌', '─', '─', '┘', ' ', '└', '┘', ' ', '╚', '═', '═', '═', '═', '╗'},
		{'║', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '│', '│', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '║'},
		{'║', ' ', '┌', '─', '─', '┐', ' ', '┌', '─', '─', '─', '┐', ' ', '│', '│', ' ', '┌', '─', '─', '─', '┐', ' ', '┌', '─', '─', '┐', ' ', '║'},
		{'║', ' ', '└', '─', '┐', '│', ' ', '└', '─', '─', '─', '┘', ' ', '└', '┘', ' ', '└', '─', '─', '─', '┘', ' ', '│', '┌', '─', '┘', ' ', '║'},
		{'║', ' ', ' ', ' ', '│', '│', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '│', '│', ' ', ' ', ' ', '║'},
		{'╟', '─', '┐', ' ', '│', '│', ' ', '┌', '┐', ' ', '┌', '─', '─', '─', '─', '─', '─', '┐', ' ', '┌', '┐', ' ', '│', '│', ' ', '┌', '─', '╢'},
		{'╟', '─', '┘', ' ', '└', '┘', ' ', '│', '│', ' ', '└', '─', '─', '┐', '┌', '─', '─', '┘', ' ', '│', '│', ' ', '└', '┘', ' ', '└', '─', '╢'},
		{'║', ' ', ' ', ' ', ' ', ' ', ' ', '│', '│', ' ', ' ', ' ', ' ', '│', '│', ' ', ' ', ' ', ' ', '│', '│', ' ', ' ', ' ', ' ', ' ', ' ', '║'},
		{'║', ' ', '┌', '─', '─', '─', '─', '┘', '└', '─', '─', '┐', ' ', '│', '│', ' ', '┌', '─', '─', '┘', '└', '─', '─', '─', '─', '┐', ' ', '║'},
		{'║', ' ', '└', '─', '─', '─', '─', '─', '─', '─', '─', '┘', ' ', '└', '┘', ' ', '└', '─', '─', '─', '─', '─', '─', '─', '─', '┘', ' ', '║'},
		{'║', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '║'},
		{'╚', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '═', '╝'},
	}

	PrintLayer(layer)
}

func PrintLayer(l [31][28]rune) {
	for r := 0; r < len(l); r++ {
		for c := 0; c < len(l[r]); c++ {
			fmt.Printf("%c", l[r][c])
		}
		fmt.Println()
	}
}
*/
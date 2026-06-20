package main

import "fmt"

func main() {
	var chessGrid string
	for строка := 0; строка < 8; строка++ {
		for столбец := 0; столбец < 8; столбец++ {
			if (строка+столбец)%2 == 0 {
				chessGrid += " "
			} else {
				chessGrid += "#"
			}
		}
		chessGrid += "\n"

	}
	fmt.Print(chessGrid)
}

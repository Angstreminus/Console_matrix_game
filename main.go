package main

import "fmt"

func printMatrix(matr [][]string) {
	fmt.Println("  0   1    2")
	for i := 0; i < 3; i++ {
		fmt.Print(i)
		for j := 0; j < 3; j++ {
			fmt.Print(" [", matr[i][j], "] ")
		}
		fmt.Println()
	}

}

func validateMove(matr [][]string, dx, dy int) bool {
	//! coordinates must be in [0..3] diapasone"
	if (dx < 0 || dx > 3) && (dx < 0 || dx < 3) {
		fmt.Println("Coordinates must be in [0..3] diapasone")
		return false
	}

	//! field isn`t empty
	if matr[dx][dy] != "" {
		fmt.Println("Field isn`t empty")
		return false
	}

	//! all is okay
	return true
}

func playerMove(matr [][]string, dx, dy int, sym string) {
	matr[dx][dy] = sym
}

func checkFullFilled(matr [][]string) bool { //! matrix hasn`t empty strings ""
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (matr[i][j] != "X") && (matr[i][j] != "O") {
				return false
			}
		}
	}
	return true
}

//! check 1st or 2nd player wins
func checkCombinations(matr [][]string) (bool, bool) {
	cnt1, cnt2 := 0, 0
	// ! check lines
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matr[i][j] == "X" {
				cnt1++
			}
			if matr[i][j] == "O" {
				cnt2++
			}
		}
		if cnt1 == 3 {
			return true, false
		}
		if cnt2 == 3 {
			return false, true
		}
	}

	// ! check columns
	cnt1, cnt2 = 0, 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matr[j][i] == "X" {
				cnt1++
			}
			if matr[j][i] == "O" {
				cnt2++
			}
		}

		if cnt1 == 3 {
			return true, false
		}
		if cnt2 == 3 {
			return false, true
		}
	}

	//! check main diagonal
	cnt1, cnt2 = 0, 0
	for i := 0; i < 3; i++ {
		j := i
		if matr[i][j] == "X" {
			cnt1++
		}

		if matr[i][j] == "O" {
			cnt2++
		}
		j++
		if cnt1 == 3 {
			return true, false
		}
		if cnt2 == 3 {
			return false, true
		}
	}

	//! check sub diagonal

	cnt1, cnt2 = 0, 0
	for i := 2; i > 0; i-- {
		j := i
		if matr[i][j] == "X" {
			cnt1++
		}

		if matr[i][j] == "O" {
			cnt2++
		}
		j++
		if cnt1 == 3 {
			return true, false
		}
		if cnt2 == 3 {
			return false, true
		}
	}
	//! nobody wins
	return false, false
}

func main() {
	fmt.Println("" == "O")

	fmt.Println("Start game! ")

	//! initialize matrix
	arr := make([][]string, 3)
	for i := range arr {
		arr[i] = make([]string, 3)
	}

	printMatrix(arr)
	var (
		dx, dy int
	)
	for {
		var player1Win, player2Win bool

		fmt.Println("Player 1, input your coordinates: [X] [Y]")

		fmt.Scanf("%d %d", &dx, &dy)
		for {
			//! Force user unil data will be valid
			if validateMove(arr, dx, dy) {
				break
			} else {
				fmt.Println("Player 1, input your coordinates: [X] [Y]")
				fmt.Scanf("%d %d", &dx, &dy)
			}
		}
		playerMove(arr, dx, dy, "X")
		printMatrix(arr)

		player1Win, player2Win = checkCombinations(arr)

		if !player1Win && !player2Win && checkFullFilled(arr) {
			fmt.Println("NOBODY WINS!")
			break
		}

		if player1Win {
			fmt.Println("Player 1 WINS!")
			break
		}

		fmt.Println("Player 2, input your coordinates: [X] [Y]")
		fmt.Scanf("%d %d", &dx, &dy)
		//! Validating move
		for {
			//! Force user unil data will be valid
			if validateMove(arr, dx, dy) {
				break
			} else {
				fmt.Println("Player 1, input your coordinates: [X] [Y]")
				fmt.Scanf("%d %d", &dx, &dy)
			}
		}
		playerMove(arr, dx, dy, "O")
		printMatrix(arr)

		player1Win, player2Win = checkCombinations(arr)

		if !player1Win && !player2Win && checkFullFilled(arr) {
			fmt.Println("NOBODY WINS!")
			break
		}

		if player2Win {
			fmt.Println("Player 2 WINS!")
			break
		}
	}
}

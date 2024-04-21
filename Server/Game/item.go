package Game

import "strconv"

func addAppearnce(state BoardState) [][]item {
	selcX := state.xPos
	selcY := state.yPos
	board := state.Board
	squareX := selcX / 3
	squareY := selcY / 3
	square := squareX*3 + squareY
	nBoard := make([][]item, 9)
	for i := 0; i < 9; i++ {
		temp := make([]item, 9)
		for j := 0; j < 9; j++ {
			valueSel := board[selcX][selcY]
			if selcX == i && selcY == j {
				temp[j] = item{Value: board[i][j], Appearance: 3, TextColour: 1, Get: "click/" + strconv.Itoa(i)+ strconv.Itoa(j)}
			} else if inSquare(square, i, j) {
				temp[j] = item{Value: board[i][j], Appearance: ifSameElse(board[i][j], valueSel, 1), TextColour: getTextColour(state, i, j, valueSel), Get: "click/" + strconv.Itoa(i) + "/" + strconv.Itoa(j)}
			} else if i == selcX {
				temp[j] = item{Value: board[i][j], Appearance: ifSameElse(board[i][j], valueSel, 1), TextColour: getTextColour(state, i, j, valueSel), Get: "click/" + strconv.Itoa(i) + "/" + strconv.Itoa(j)}
			} else if j == selcY {
				temp[j] = item{Value: board[i][j], Appearance: ifSameElse(board[i][j], valueSel, 1), TextColour: getTextColour(state, i, j, valueSel), Get: "click/" + strconv.Itoa(i) + "/" + strconv.Itoa(j)}
			} else {
				temp[j] = item{Value: board[i][j], Appearance: ifSameElse(board[i][j], valueSel, 0), TextColour: getTextColour(state, i, j, valueSel), Get: "click/" + strconv.Itoa(i) + "/" + strconv.Itoa(j)}
			}
		}
		nBoard[i] = temp
	}
	return nBoard
}

func getTextColour(state BoardState, x int, y int, val int) int {
	if state.Board[x][y] == val && val != 0 {
		return 2
	}
	if state.OGBoad[x][y] != 0 {
		return 0
	} else if state.Board[x][y] != 0 {
		if state.FullBoard[x][y] == state.Board[x][y] {
			return 4
		} else {
			return 3
		}
	}
	return 1
}

func ifSameElse(val int, cVal int, els int) int {
	if val == cVal && val != 0 {
		return 2
	}
	return els
}

func inSquare(num int, col int, row int) bool {
	return ((col/3)*3)+(row/3) == num
}

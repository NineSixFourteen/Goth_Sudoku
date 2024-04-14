package src

import (
	"fmt"
	"math/rand/v2"
)

func Ten() int {
    return 10; 
}

type Pair struct {
    board [][]int
    complete bool
}

// A solver function that should only find one solution to the puzzle 
// Function is used to check if a board is solvable
func Solver(board [][]int) Pair {
    nBoard := clone(board)
    if !isLegal(board) {
        return Pair{board:[][]int{}, complete:false}
    }
    for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if nBoard[x][y] == 0 {
				for number := 1; number < 10; number++ {
					if possible(board, x, y, number) {
                        nBoard[x][y] = number;
                        nPair := Solver(nBoard)
                        if nPair.complete {
                            return nPair 
                        } else {
                            nBoard[x][y] = 0;
                        }
					}
				}
                return Pair{board:nBoard, complete: false} 
			}
		}
	}
    return Pair{board:nBoard, complete: true}
}

func isLegal(board [][] int) bool {
    if illegalRow(board){
        return false;
    }
    if illegalCol(board){
        return false;
    }
    if illegalSquare(board){
        return false;
    }
    return true;
}

func illegalRow(board [][]int) bool {
    for row := 0; row < 9; row++ {
        items := make([]int, 0);
        for rowItem := 0; rowItem < 9; rowItem++ {
            for i := 0; i < len(items);i++ {
                if board[row][rowItem] == items[i] {
                    return true;
                }
            }
            if board[row][rowItem] != 0 {
                items = append(items, board[row][rowItem]);
            }
        }
    }
    return false;
}

func RemoveSquares(board [][]int, num int) [][]int {
    nBoard := clone(board)
    for i := 0; i < num;i++ {
        x := rand.IntN(9);
        y := rand.IntN(9);
        temp := nBoard[x][y]
        nBoard[x][y] = 0;
        for !isUnique(nBoard) {
            nBoard[x][y] = temp;
            x := rand.IntN(9);
            y := rand.IntN(9);
            temp = nBoard[x][y];
            nBoard[x][y] = 0;
        }
    }
    return nBoard
}

var found bool = false
func isUnique(board[][]int) bool {
    found = false;
    return isUniqueHelper(board)
}

func isUniqueHelper(board [][]int) bool {
    nBoard := clone(board)
    for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if nBoard[x][y] == 0 {
				for number := 1; number < 10; number++ {
					if possible(board, x, y, number) {
                        nBoard[x][y] = number;
                        unq := isUniqueHelper(nBoard)
                        if(!unq){
                            return false    
                        }
                        nBoard[x][y] = 0;
					}
				}
                return true 
			}
		}
	}
    if !found {
        found = true
        return true
    }else {
        return false
    }
}

func illegalCol(board [][]int) bool {
    for col := 0; col < 9; col ++ {
    items := make([]int, 0);
        for colItem := 0; colItem < 9; colItem++ {
            for i := 0; i < len(items);i++ {
                if board[colItem][col] == items[i] {
                    return true;
                } 
            }
            if board[colItem][col] != 0 {
                items = append(items, board[colItem][col]);
            }
        }
    }
    return false;
}

func illegalSquare(board [][]int) bool {
    for grow := 0; grow < 3; grow ++{
        for gcol := 0; gcol <3; gcol ++ {
            items := make([]int , 0) 
            for srow := 0; srow < 3; srow++ {
                for scol := 0; scol < 3; scol++ {
                    val := board[grow * 3 + srow][gcol * 3 + scol]
                    for i := 0; i < len(items); i++ {
                        if val == items[i] {
                            return true;
                        }
                    }
                    if val != 0 {
                        items = append(items, val)
                    }
                }
            }
        }
    }
    return false 
}

func displayBoard(board [][]int) {
    fmt.Println("Board")
    for i := 0; i < 9; i++ {
        fmt.Println(board[i])
    }
}

func clone(board [][]int) [][] int {
    duplicate := make([][]int, 9)
    for i := range 9 {
        matrix := board[i]
        duplicate[i] = make([]int, 9)
        copy(duplicate[i], matrix)
    }
    return duplicate
}



func Solvable(board [][]int) bool {
    return Solver(board).complete
}

func possible(board [][]int, x int, y int, number int)  bool{
	if checkRow(board, x, number) {
		if checkCol(board, y, number) {
			if checkSquare(board, x, y, number) {
				return true
			}
		}
	}
	return false
}

func checkRow(board [][]int, x int, number int) bool {
    for ind := 0; ind < 9; ind++ {
        val := board[x][ind];
        if(val == number){
            return false;
        }
    }
    return true;
}
    

func checkCol(board [][]int, y int, number int) bool {
    for ind := 0; ind < 9; ind++ {
        val := board[ind][y];
        if(val == number){
            return false;
        }
    }
    return true;
}

func checkSquare(board [][]int,x int , y int, number int) bool { 
    for indx := (x / 3) * 3; indx < ((x / 3) + 3) + (2 - (x % 3)); indx++ {
        for indy := (y / 3) * 3; indy < ((y / 3) + 3) + (2 - (y % 3)); indy++ {
            if board[indx][indy] == number {
                return false;
            }
        }
    }
    return true
}

// start = (div * 3), end = num + (2 - rem) 
// 2, 7
// 2, start 0 * 3 = 0, end 2 + 0 = 2
// 7, start 2 * 3 = 6, end 7 + 1 = 8  
// 0, 8
// 0, start 0 * 3 = 0, end 0 + 2 = 2 
// 8, start 2 * 3 = 6, end 8 + 0 = 8 


    
// For Every Tile
//  If tile == 0 i.e. empty
//      for every possible number 1-9
//        If number is allowed there i.e. sudoku rules
//          place number there
//          start function again
//          if function ever returns undo move and try next number
//      return function if there are no possible moves
//  If there are no empty tiles return grid

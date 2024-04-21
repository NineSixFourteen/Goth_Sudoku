package Game

import (
	"math/rand/v2"
	"strconv"
)


func emptyBoard() [][]int {
    board := make([][]int, 9);
    for i:= 0; i < 9; i++{
        board[i] = make([]int, 9)
    }
    return board
}


func ToBoard(state BoardState) Board {
    board := addAppearnce(state)
    return Board{
        Board: [][]Square{
            {
                toSquare(
                []item{
                    board[0][0],board[0][1],board[0][2],
                    board[1][0],board[1][1],board[1][2],
                    board[2][0],board[2][1],board[2][2],
                }),
                toSquare(
                []item{
                    board[3][0],board[3][1],board[3][2],
                    board[4][0],board[4][1],board[4][2],
                    board[5][0],board[5][1],board[5][2],
                }),
                toSquare(
                []item {
                    board[6][0],board[6][1],board[6][2],
                    board[7][0],board[7][1],board[7][2],
                    board[8][0],board[8][1],board[8][2],
                }),
            },
            {
            toSquare(
                []item {
                    board[0][3],board[0][4],board[0][5],
                    board[1][3],board[1][4],board[1][5],
                    board[2][3],board[2][4],board[2][5],
                }),
                toSquare(
                []item{
                    board[3][3],board[3][4],board[3][5],
                    board[4][3],board[4][4],board[4][5],
                    board[5][3],board[5][4],board[5][5],
                }),
                toSquare(
                []item {
                    board[6][3],board[6][4],board[6][5],
                    board[7][3],board[7][4],board[7][5],
                    board[8][3],board[8][4],board[8][5],
                }),
            },
            {
          toSquare(
                []item {
                    board[0][6],board[0][7],board[0][8],
                    board[1][6],board[1][7],board[1][8],
                    board[2][6],board[2][7],board[2][8],
                }),
                toSquare(
                []item {
                    board[3][6],board[3][7],board[3][8],
                    board[4][6],board[4][7],board[4][8],
                    board[5][6],board[5][7],board[5][8],
                }),
                toSquare(
                []item {
                    board[6][6],board[6][7],board[6][8],
                    board[7][6],board[7][7],board[7][8],
                    board[8][6],board[8][7],board[8][8],
                }),
            },
        },
    }
}
    
func GetGame(game GameState) Game{
    return Game{
        Board: ToBoard(game.Board),
        UI: ToUI(game.UI),
    }
}

func ToUI(ui UIState) UI {
    return UI {
        Mistakes: strconv.Itoa(ui.Mistakes) + "/" + getMistakes(ui.MaxMistakes),
        Hints: strconv.Itoa(ui.Hints) + "/" + getMistakes(ui.MaxHints),
        Tiles: ui.RTiles,
    }
}

func getMistakes(num int) string {
    if num > 0 {
        return strconv.Itoa(num)
    } else {
        return "∞"
    }
}

func toSquare( nums []item) Square{
    return Square {
        Square: [][]item{
            {nums[0],nums[1],nums[2]},
            {nums[3],nums[4],nums[5]},
            {nums[6],nums[7],nums[8]},
        },
    }
}

func GetNewBoard(diff  int ) GameState{
    pieces := 0;
    mistakes := 0; 
    hints := 0;
    switch(diff) {
        case 0: 
            pieces = 51;
            mistakes = -1;
            hints = -1;
        case 1: 
            mistakes = 5;
            pieces = 60;
            hints = 3; 
        case 2: 
            mistakes = 3; 
            pieces = 65;
            hints = 0;
        default:
            hints = -1;
            mistakes = -1;
            pieces = 40;
    }
    fullBoard := makeNewBoard();
    board := RemoveSquares(fullBoard,pieces)
    boardS := BoardState{
        OGBoad: board,
        Board: Clone(board),
        FullBoard: fullBoard,
        xPos: 4,
        yPos: 4,
    }
    UI := UIState{
        Mistakes: 0,
        MaxMistakes: mistakes,
        Hints: 0,
        MaxHints: hints,
        RTiles: 81 - pieces,
    }
    return GameState{
        Board: boardS,
        UI: UI,
    }
}

func Clone(og [][]int ) [][]int {
    ret := make([][]int, 9);
    for i := 0; i < 9; i++ {
        temp := make([]int, 9)
        for j := 0; j < 9; j++ {
            temp[j] = og[i][j];
        }
        ret[i] = temp;
    }
    return ret;
}

func makeNewBoard() [][]int {
    board := emptyBoard();
    for x := 0; x < 9; x++ {
        for y := 0; y < 9; y++ {
            num := rand.IntN(9) + 1;
            board[x][y] =  num;
            for !Solvable(board){
                num = rand.IntN(9) + 1;
                board[x][y] = num;
            }
        }
    }
    return board;
}

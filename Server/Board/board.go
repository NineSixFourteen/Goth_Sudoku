package Board

import "math/rand/v2"

type item struct {
    Value int 
    Appearance int
}

type Square struct {
    Square [][]item
}

type Board struct {
    Board [][]Square
}

func NewBoard() Board {
    a := make([][]Square, 3)
    for i := 0; i < 3; i++ {
        a[i] = make([]Square, 3);
        for j := 0; j < 3;j++ {
            a[i][j] = newSquare();
        }
    }
    return Board{
        Board: a, 
    }
}

func newSquare() Square {
    a := make([][]item, 3)
    for i := 0; i < 3; i++ {
        a[i] = make([]item, 3);
        for j := 0; j < 3;j++ {
            a[i][j] = item{Value: 1, Appearance: 0} ;
        } 
    }
    return Square {
        Square: a,
    }
}

func emptyBoard() [][]int {
    board := make([][]int, 9);
    for i:= 0; i < 9; i++{
        board[i] = make([]int, 9)
    }
    return board
}


func ToBoard(board [][]item) Board {
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
                    board[5][6],board[5][7],board[6][8],
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

func toSquare( nums []item) Square{
    return Square {
        Square: [][]item{
            {nums[0],nums[1],nums[2]},
            {nums[3],nums[4],nums[5]},
            {nums[6],nums[7],nums[8]},
        },
    }
}

func GetNewBoard(diff  int) Board {
    pieces := 0;
    switch(diff) {
        case 0: 
            pieces = 51; 
        case 1: 
            pieces = 60;
        case 2: 
            pieces = 65;
        default: 
            pieces = 40;
    }
    board := ToBoard(addAppearnce(RemoveSquares(makeNewBoard(),pieces), 2,8))
    return board
}

func addAppearnce(board [][]int, selcX int, selcY int) [][]item {
    squareX := selcX /3;
    squareY := selcY /3;
    square := squareX * 3 + squareY;
    nBoard := make([][]item, 9);
    for i := 0; i < 9 ; i++ {
        temp := make([]item, 9);
            for j := 0; j < 9; j++ {
                if selcX == i && selcY == j {
                    temp[j] = item{Value: board[i][j], Appearance: 2} 
                } else if inSquare(square, i,j) {
                    temp[j] = item{Value: board[i][j], Appearance: 1};
                } else if i == selcX {
                    temp[j] = item{Value: board[i][j], Appearance:1} 
                } else if j == selcY {
                    temp[j] = item{Value: board[i][j], Appearance:1} 
                } else {
                    temp[j] = item{Value: board[i][j], Appearance:0} 
                }
            }
        nBoard[i] = temp
        }        
    return nBoard;
}

func inSquare(num int, col int, row int) bool {
    return ((col/ 3) * 3) + (row/3) == num;
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

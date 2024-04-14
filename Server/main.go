package main

import (
	"Server/main/src"
	"fmt"
	"html/template"
	"io"
	"math/rand/v2"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Square struct {
    Square [][]int
}


type Board struct {
    Board [][]Square
}

type Data struct{
    Button []Button
}

type Button struct {
    Get string 
    Message string
}

func createButtons() []Button{
    buttons := make([]Button, 4)
    buttons[0] = Button{
        Get: "/board",
        Message: "Easy",
    }
    buttons[1] = Button{
        Get: "/newGame/0",
        Message: "Medium",
    }
    buttons[2] = Button{
        Get: "",
        Message: "Hard",
    }
    buttons[3] = Button{
        Get: "",
        Message: "Other",
    }
    return buttons
}

func newBoard() Board {
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
    a := make([][]int , 3)
    for i := 0; i < 3; i++ {
        a[i] = make([]int, 3);
        for j := 0; j < 3;j++ {
            a[i][j] = 1 ;
        } 
    }
    return Square {
        Square: a,
    }
}

func newData() Data {
    return Data{
        Button: createButtons(),
    }
}

type Templates struct {
	templates *template.Template;
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("Templates/*.html")),
	}
}

var gameBoard Board = newBoard();
var logger echo.Logger = echo.New().Logger
func main() {
    fmt.Println("")
    e := echo.New();
    e.Renderer = NewTemplates();
    e.Static("/", "static/css")
    e.Use(middleware.Logger())
    e.GET("/", hello);
    e.GET("/board",board);
    e.GET("/newGame/:id",
        func (c echo.Context) error {
            id := c.QueryParam("id")
            ID,_ := strconv.Atoi(id);
            gameBoard = getNewBoard(ID)    
            return c.Render(http.StatusOK,"board", gameBoard)
        },
    )
    e.Logger.Fatal(e.Start(":1323"))
}

func getNewBoard(diff  int) Board {
    board := toBoard(makeNewBoard())
    return removeSquare(board, diff)
}

func makeNewBoard() [][]int {
    board := emptyBoard();
    for x := 0; x < 9; x++ {
        for y := 0; y < 9; y++ {
            num := rand.IntN(9) + 1;
            board[x][y] =  num;
            for !src.Solvable(board){
                num = rand.IntN(9) + 1;
                board[x][y] = num;
            }
        }
    }
    return board;
}

func emptyBoard() [][]int {
    board := make([][]int, 9);
    for i:= 0; i < 9; i++{
        board[i] = make([]int, 9)
    }
    return board
}



func toBoard(board [][]int) Board {
    return Board{
        Board: [][]Square{
            {
                toSquare(
                []int {
                    board[0][0],board[0][1],board[0][2],
                    board[1][0],board[1][1],board[1][2],
                    board[2][0],board[2][1],board[2][2],
                }),
                toSquare(
                []int {
                    board[3][0],board[3][1],board[3][2],
                    board[4][0],board[4][1],board[4][2],
                    board[5][0],board[5][1],board[5][2],
                }),
                toSquare(
                []int {
                    board[6][0],board[6][1],board[6][2],
                    board[7][0],board[7][1],board[7][2],
                    board[8][0],board[8][1],board[8][2],
                }),
            },
            {
            toSquare(
                []int {
                    board[0][3],board[0][4],board[0][5],
                    board[1][3],board[1][4],board[1][5],
                    board[2][3],board[2][4],board[2][5],
                }),
                toSquare(
                []int {
                    board[3][3],board[3][4],board[3][5],
                    board[4][3],board[4][4],board[4][5],
                    board[5][3],board[5][4],board[5][5],
                }),
                toSquare(
                []int {
                    board[6][3],board[6][4],board[6][5],
                    board[7][3],board[7][4],board[7][5],
                    board[8][3],board[8][4],board[8][5],
                }),
            },
            {
          toSquare(
                []int {
                    board[0][6],board[0][7],board[0][8],
                    board[1][6],board[1][7],board[1][8],
                    board[2][6],board[2][7],board[2][8],
                }),
                toSquare(
                []int {
                    board[3][6],board[3][7],board[3][8],
                    board[4][6],board[4][7],board[4][8],
                    board[5][6],board[5][7],board[6][8],
                }),
                toSquare(
                []int {
                    board[6][6],board[6][7],board[6][8],
                    board[7][6],board[7][7],board[7][8],
                    board[8][6],board[8][7],board[8][8],
                }),
            },
        },
    }
}

func toSquare( nums []int) Square{
    return Square {
        Square: [][]int{
            {nums[0],nums[1],nums[2]},
            {nums[3],nums[4],nums[5]},
            {nums[6],nums[7],nums[8]},
        },
    }
}


func removeSquare(board Board, diff int) Board {
    return board;
    //switch(diff){
      //  case 0:
        //    return removeSquareEasy(board);
        //case 1: 
         //   return removeSquareMedium(board);
       // case 2:
        //    return removeSquareHard(board);
       // default:
        //    return removeSquareOther(board);
    //}
}

func removeSquareEasy(board Board) Board {
    return newBoard();
}
func removeSquareMedium(board Board) Board {
    return newBoard();
}
func removeSquareHard(board Board) Board {
    return newBoard();
}
func removeSquareOther(board Board) Board {
    return newBoard();
}

func hello (c echo.Context) error {
    return c.Render(http.StatusOK,"index", newData())
}

func board(c echo.Context) error {
    gameBoard = toBoard(makeNewBoard())
    return c.Render(http.StatusOK,"board", gameBoard)
}



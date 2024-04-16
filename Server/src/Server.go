package src

import (
	"Server/main/Board"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)


type Button struct {
    Get string 
    Message string
}

type GameButton struct {
    First string
    Rest string
}

type PageState struct {
    DifficultyButtons []Button 
    SmallButtons      []GameButton
    BigButtons        []GameButton
    KeyPad            [][]int
    Board             Board.Board
}


func createButtons() []Button{
    buttons := make([]Button, 4)
    buttons[0] = Button{
        Get: "/newGame/0",
        Message: "Easy",
    }
    buttons[1] = Button{
        Get: "/newGame/1",
        Message: "Medium",
    }
    buttons[2] = Button{
        Get: "/newGame/2",
        Message: "Hard",
    }
    buttons[3] = Button{
        Get: "/newGame/5",
        Message: "Other",
    }
    return buttons
}

func createSmallButtons() []GameButton {
    buttons := make([]GameButton,2)
    buttons[0] = GameButton{
        First: "U",
        Rest: "ndo",
    }
    buttons[1] = GameButton{
        First: "R",
        Rest: "edo",
    }
    return buttons
}

func createBigButtons() []GameButton {
    buttons := make([]GameButton,2)
    buttons[0] = GameButton{
        First: "N",
        Rest: "otes",
    }
    buttons[1] = GameButton{
        First: "G",
        Rest: "ive up",
    }
    return buttons
}

func keyPad() [][]int {
    return [][]int{
        {1,2,3},
        {4,5,6},
        {7,8,9},
    }
}

func initPage() PageState {
    state = Board.GetNewBoard(0, state);
    return PageState{
        Board: Board.GetBoard(state),
        DifficultyButtons: createButtons(), 
        SmallButtons: createSmallButtons(), 
        BigButtons: createBigButtons(), 
        KeyPad: keyPad(),
    }
}

var logger = echo.New().Logger
var state Board.BoardState; 

func StartServer(){
    e := echo.New();
    logger = e.Logger;
    logger.SetLevel(log.DEBUG)
    e.Renderer = NewTemplates();
    e.Static("/", "static/css")
    e.Use(middleware.Logger())
    e.GET("/", hello);
    e.GET("/newGame/:id",board);
    e.GET("/empty/:id", LoggerRequest);
    e.GET("/key/:id", keyHandler);
    e.Logger.Info(e.Start(":1323"))
}



func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Templates struct {
	templates *template.Template;
}

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("Templates/*.html")),
	}
}

func hello (c echo.Context) error {
    return c.Render(http.StatusOK,"index", initPage())
}

func board(c echo.Context) error {
    diffs := c.Param("id")
    diff, _ := strconv.Atoi(diffs)
    logger.Info(c)
    state = Board.GetNewBoard(diff, state);
    return c.Render(http.StatusOK,"board", Board.GetBoard(state))  
}

func LoggerRequest(c echo.Context) error {
    logger.Info("C Botta ")
    logger.Info(c)
    return c.Render(http.StatusNoContent, "empty", []int{})
}

func keyHandler(c echo.Context) error {
    val := c.Param("id");
    vals, _ := strconv.Atoi(val)
    state = Board.Handle(vals,state);
    logger.Info(state)
    return c.Render(http.StatusOK, "board", Board.GetBoard(state))
}


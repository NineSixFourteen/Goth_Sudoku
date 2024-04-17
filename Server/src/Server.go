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
    KeyPad            [][]Key
    StartDiff         string
}

type Key struct {
    Value int 
    Addr string
}


func createButtons() []Button{
    buttons := make([]Button, 4)
    buttons[0] = Button{
        Get: "/new/0",
        Message: "Easy",
    }
    buttons[1] = Button{
        Get: "/new/1",
        Message: "Medium",
    }
    buttons[2] = Button{
        Get: "/new/2",
        Message: "Hard",
    }
    buttons[3] = Button{
        Get: "/new/5",
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

func keyPad() [][]Key {
    keys := make([][]Key, 3)
    for i := 0; i < 3;i++ {
        temp := make([]Key,3) 
        for j := 0; j < 3; j++ {
            temp[j] = Key{
                Value: i * 3 + j + 1,
                Addr: "key/" + strconv.Itoa(i * 3 + j + 1)}
        }
        keys[i] = temp;
    }
    return keys
}

func initPage() PageState {
    return PageState{
        DifficultyButtons: createButtons(), 
        SmallButtons: createSmallButtons(), 
        BigButtons: createBigButtons(), 
        KeyPad: keyPad(),
        StartDiff: "/newGame/1/",
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
    e.Static("/Images/", "static/Images")
    e.Use(middleware.Logger())
    e.GET("/", hello);
    e.GET("/new/:id", new);
    e.GET("/newGame/:id",board);
    e.GET("/empty/:id", LoggerRequest);
    e.GET("/key/:id", keyHandler);
    e.GET("/click/:x/:y", clickHandler);
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

func new(c echo.Context) error {
    diffs := c.Param("id");
    return c.Render(http.StatusOK, "loadtext", "/newGame/" + diffs)
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
    logger.Info(vals)
    logger.Info(state)
    return c.Render(http.StatusOK, "board", Board.GetBoard(state))
}

func clickHandler(c echo.Context) error {
    xPos := c.Param("x");
    yPos := c.Param("y");
    x, _ := strconv.Atoi(xPos);
    y, _ := strconv.Atoi(yPos);
    state = Board.ClickHandle(x, y, state);
    return c.Render(http.StatusOK, "board", Board.GetBoard(state))
}


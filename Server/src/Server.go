package src

import (
	"Server/main/Board"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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


func newData() Data {
    return Data{
        Button: createButtons(),
    }
}

func StartServer(){
    e := echo.New();
    e.Renderer = NewTemplates();
    e.Static("/", "static/css")
    e.Use(middleware.Logger())
    e.GET("/", hello);
    e.GET("/newGame/:id",board);
    e.Logger.Fatal(e.Start(":1323"))
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
    return c.Render(http.StatusOK,"index", newData())
}

func board(c echo.Context) error {
    diffs := c.QueryParam("id")
    diff, _ := strconv.Atoi(diffs)
    return c.Render(http.StatusOK,"board", Board.GetNewBoard(diff))  
}

package src

import (
	"Server/main/Game"
	"html/template"
	"io"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var logger = echo.New().Logger
var state Game.GameState; 

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

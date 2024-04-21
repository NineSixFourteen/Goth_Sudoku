package src

import (
	"Server/main/Game"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

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
    state = Game.GetNewBoard(diff);
    return c.Render(http.StatusOK,"Game", gameToData(Game.GetGame(state)))  
}

func LoggerRequest(c echo.Context) error {
    logger.Info("C Botta ")
    logger.Info(c)
    return c.Render(http.StatusNoContent, "empty", []int{})
}

func keyHandler(c echo.Context) error {
    val := c.Param("id");
    vals, _ := strconv.Atoi(val)
    state = Game.Handle(vals,state);
    logger.Info(vals)
    logger.Info(state)
    return c.Render(http.StatusOK, "Game", gameToData(Game.GetGame(state)))
}

func clickHandler(c echo.Context) error {
    xPos := c.Param("x");
    yPos := c.Param("y");
    x, _ := strconv.Atoi(xPos);
    y, _ := strconv.Atoi(yPos);
    state = Game.ClickHandle(x, y, state);
    return c.Render(http.StatusOK, "Game", gameToData(Game.GetGame(state)))
}

func gameToData(game Game.Game) Data {
    return Data{
        SmallButtons: createSmallButtons(),
        BigButtons: createBigButtons(),
        KeyPad: keyPad(),
        Difficulty: "",
        HasBoard: true,
        Board: game.Board,
        UI: game.UI,
    }
}


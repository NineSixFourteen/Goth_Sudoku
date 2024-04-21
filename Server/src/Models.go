package src

import "Server/main/Game"

type Data struct {
    SmallButtons      []GameButton
    BigButtons        []GameButton
    KeyPad            [][]Key 
    Difficulty        string
    HasBoard          bool 
    Board             Game.Board
    UI                Game.UI
}

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
    Data Data
    Other Data 
}

type Key struct {
    Value int 
    Addr string
 }

 type Stats struct {
    Mistakes string
    Hints string 
    Tiles int 
 }

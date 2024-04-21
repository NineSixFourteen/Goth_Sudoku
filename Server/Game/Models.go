package Game 

//Frontend

type Game struct {
    Board Board
    UI   UI 
}

type Board struct {
    Board [][]Square
}

type Square struct {
    Square [][]item
}

type item struct {
    Value int 
    Appearance int
    TextColour int
    Get string
}

type UI struct {
    Mistakes string 
    Hints string 
    Tiles int 
}

// Backend


type GameState struct {
    Board BoardState 
    UI    UIState
}

type BoardState struct {
    OGBoad [][]int 
    Board  [][]int
    FullBoard [][]int
    xPos   int 
    yPos   int
}

type UIState struct {
    Mistakes int 
    MaxMistakes int 
    Hints    int 
    MaxHints int 
    RTiles   int 
}

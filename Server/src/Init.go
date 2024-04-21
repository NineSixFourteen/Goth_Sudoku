package src

import (
	"Server/main/Game"
	"strconv"
)

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
        Other: Data {   
            SmallButtons: createSmallButtons(), 
            BigButtons: createBigButtons(), 
            KeyPad: keyPad(),
            Difficulty: "/newGame/1/",
            HasBoard: false,
            Board: Game.Board{},
            UI: Game.UI{
                Mistakes: "0/0", 
                Hints:    "0/0", 
                Tiles:    21,
            },
        },
    }
}

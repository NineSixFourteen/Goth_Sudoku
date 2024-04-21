package Game;  

func ClickHandle(xPos int, yPos int, state GameState) GameState {
    if xPos < 10 {
        state.Board.xPos = xPos;
        state.Board.yPos = yPos;
    } else {
        //KeyPad Stuff
    }
    return state;
} 

func Handle(key int, gState GameState) GameState {
    state := gState.Board; 
    if len(state.Board) > 0{
        if key < 10 {
            gState = AddNumber(gState, key);
        }
        switch(key) {
            case 10:
                if state.yPos > 0 && state.yPos <= 8{
                    state.yPos--;
                }
                break;
            case 11:
                if state.yPos >= 0 && state.yPos < 8{
                    state.yPos++;
                }   
                break;
            case 12: 
                if state.xPos > 0 && state.xPos <= 8{
                    state.xPos-- 
                }
                break;
            case 13:
             if state.xPos >= 0 && state.xPos < 8{
                state.xPos++;
            }
                break;
        }
    }
    gState.Board = state;
    return gState;
}

func AddNumber(state GameState,key int) GameState{
    x := state.Board.xPos; 
    y := state.Board.yPos;
    if state.Board.OGBoad[x][y] == 0 {
        state.Board.Board[x][y] = key
    }
    return state;
}


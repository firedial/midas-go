package controller

import(
    "github.com/firedial/midas-go/model"
)

func MovePost(move model.Move) model.Move {
    _ = model.InsertMove(move)
    return move 
}

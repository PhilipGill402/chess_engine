package piece

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Vec2 struct {
	X	int32
	Y	int32
}

type Piece interface {
	Draw(renderer *sdl.Renderer)		error
	DrawMoves(renderer *sdl.Renderer)	error
	Move(move Vec2) 					error
}

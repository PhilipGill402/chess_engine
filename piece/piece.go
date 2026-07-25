package piece

import (
	"github.com/veandco/go-sdl2/sdl"
	"errors"
)

type Vec2 struct {
	X	int32
	Y	int32
}

type Piece interface {
	GetValue()											uint8
	GetColor()											bool
	Draw(renderer *sdl.Renderer)						error
	DrawMoves(board []Piece, renderer *sdl.Renderer)	error
	Move(move Vec2) 									error
}

func FillCircle(renderer *sdl.Renderer, centerX, centerY, radius int32) error {
	renderer.SetDrawColor(128, 128, 128, 255) // black

	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if x*x+y*y <= radius*radius {
				if err := renderer.DrawPoint(centerX+x, centerY+y); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func GetPiece(board []Piece, pos Vec2) (Piece, error) {
	if (pos.X < 0 || pos.Y < 0 || pos.X > 7 || pos.Y > 7) {
		return nil, errors.New("Index out of range");
	}

	return board[pos.Y * 8 + pos.X], nil;
}

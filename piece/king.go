package piece

import (
	"github.com/veandco/go-sdl2/sdl"	
	"github.com/veandco/go-sdl2/img"
	"chess/globals"
)

type King struct {
	pos			Vec2
	color		bool // true = white; false = black
	value		uint8	
	texture		*sdl.Texture
}

func (piece *King) GetValue() uint8 {
	return piece.value;
}

func (piece *King) GetColor() bool {
	return piece.color;
}

func (piece *King) Draw(renderer *sdl.Renderer) error {
	dst := sdl.Rect{
		X: int32(piece.pos.X * globals.CellSize + 10),
		Y: int32(piece.pos.Y * globals.CellSize + 10),
		W: 80,
		H: 80,
	};

	renderer.Copy(piece.texture, nil, &dst);

	return nil;	
}

func (piece *King) DrawMoves(board []Piece, renderer *sdl.Renderer) error {
	return nil;
}

func (piece *King) Move(mov Vec2) error {
	return nil;
}

func NewKing(x, y int32, color bool, renderer *sdl.Renderer) (*King, error) {
	pos := Vec2{
		X: x,
		Y: y,
	};

	var texture *sdl.Texture;
	var err error;

	if (color) {
		texture, err = img.LoadTexture(renderer, "Assets/white_king.png");
	} else {
		texture, err = img.LoadTexture(renderer, "Assets/black_king.png");
	}

	if (err != nil) {
		return nil, err;
	}

	return &King{
		pos: pos,
		color: color,
		value: 255,	
		texture: texture,
	}, nil;
}



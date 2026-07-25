package piece

import (
	"github.com/veandco/go-sdl2/sdl"	
	"github.com/veandco/go-sdl2/img"
	"chess/globals"
)

type Queen struct {
	pos			Vec2
	color		bool // true = white; false = black
	value		uint8	
	texture		*sdl.Texture
}

func (piece *Queen) GetValue() uint8 {
	return piece.value;
}

func (piece *Queen) GetColor() bool {
	return piece.color;
}

func (piece *Queen) Draw(renderer *sdl.Renderer) error {
	dst := sdl.Rect{
		X: int32(piece.pos.X * globals.CellSize + 10),
		Y: int32(piece.pos.Y * globals.CellSize + 10),
		W: 80,
		H: 80,
	};

	renderer.Copy(piece.texture, nil, &dst);

	return nil;	
}

func (piece *Queen) DrawMoves(board []Piece, renderer *sdl.Renderer) error {
	return nil;
}

func (piece *Queen) Move(mov Vec2) error {
	return nil;
}

func NewQueen(x, y int32, color bool, renderer *sdl.Renderer) (*Queen, error) {
	pos := Vec2{
		X: x,
		Y: y,
	};

	var texture *sdl.Texture;
	var err error;

	if (color) {
		texture, err = img.LoadTexture(renderer, "Assets/white_queen.png");
	} else {
		texture, err = img.LoadTexture(renderer, "Assets/black_queen.png");
	}

	if (err != nil) {
		return nil, err;
	}

	return &Queen{
		pos: pos,
		color: color,
		value: 9,
		texture: texture,
	}, nil;
}



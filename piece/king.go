package piece

import (
	"github.com/veandco/go-sdl2/sdl"	
	"github.com/veandco/go-sdl2/img"
)

type King struct {
	pos			Vec2
	color		bool // true = white; false = black
	texture		*sdl.Texture
}

func (piece *King) Draw(renderer *sdl.Renderer) error {
	dst := sdl.Rect{
		X: int32(piece.pos.X + 10),
		Y: int32(piece.pos.Y + 10),
		W: 80,
		H: 80,
	};

	renderer.Copy(piece.texture, nil, &dst);

	return nil;	
}

func (piece *King) DrawMoves(renderer *sdl.Renderer) error {
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
		texture: texture,
	}, nil;
}



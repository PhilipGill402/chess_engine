package piece

import (
	"github.com/veandco/go-sdl2/sdl"	
	"github.com/veandco/go-sdl2/img"
)

type Pawn struct {
	pos			Vec2
	hasMoved	bool
	color		bool // true = white; false = black
	texture		*sdl.Texture
}

func (piece *Pawn) Draw(renderer *sdl.Renderer) error {
	dst := sdl.Rect{
		X: int32(piece.pos.X + 10),
		Y: int32(piece.pos.Y + 10),
		W: 80,
		H: 80,
	};

	renderer.Copy(piece.texture, nil, &dst);

	return nil;	
}

func (piece *Pawn) DrawMoves(renderer *sdl.Renderer) error {
	return nil;
}

func (piece *Pawn) Move(mov Vec2) error {
	return nil;
}

func NewPawn(x, y int32, color bool, renderer *sdl.Renderer) (*Pawn, error) {
	pos := Vec2{
		X: x,
		Y: y,
	};

	var texture *sdl.Texture;
	var err error;

	if (color) {
		texture, err = img.LoadTexture(renderer, "Assets/white_pawn.png");
	} else {
		texture, err = img.LoadTexture(renderer, "Assets/black_pawn.png");
	}

	if (err != nil) {
		return nil, err;
	}

	return &Pawn{
		pos: pos,
		hasMoved: false,
		color: color,
		texture: texture,
	}, nil;
}



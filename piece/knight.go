package piece

import (
	"github.com/veandco/go-sdl2/sdl"	
	"github.com/veandco/go-sdl2/img"
	"chess/globals"
	"errors"
)

type Knight struct {
	pos			Vec2
	color		bool // true = white; false = black
	value		uint8	
	texture		*sdl.Texture
}

func (piece *Knight) getMoves(board []Piece) []Vec2 {
	moves := make([]Vec2, 0);	

	for i := 0; i < 2; i++ {
		for diff := int32(-2); diff < 3; diff += 4 {
			for j := int32(-1); j < 2; j += 2 {
				pos := piece.pos;
				
				if (i == 0) {
					pos.X += diff;
					pos.Y += j;
				} else {
					pos.X += j;
					pos.Y += diff;
				}
	
				moves = append(moves, pos);
			}
		}
	}

	return moves;
}

func (piece *Knight) GetValue() uint8 {
	return piece.value;
}

func (piece *Knight) GetColor() bool {
	return piece.color;
}

func (piece *Knight) Draw(renderer *sdl.Renderer) error {
	dst := sdl.Rect{
		X: int32(piece.pos.X * globals.CellSize + 10),
		Y: int32(piece.pos.Y * globals.CellSize + 10),
		W: 80,
		H: 80,
	};

	renderer.Copy(piece.texture, nil, &dst);

	return nil;	
}

func (piece *Knight) DrawMoves(board []Piece, renderer *sdl.Renderer) error {
	moves := piece.getMoves(board);
	if (moves == nil) {
		return errors.New("No moves found\n");
	}

	for _, move := range moves {
		pos := Vec2 {
			X: (move.X * globals.CellSize) + (globals.CellSize / 2),
			Y: (move.Y * globals.CellSize) + (globals.CellSize / 2),
		};
		FillCircle(renderer, pos.X, pos.Y, int32(globals.CellSize / 4));
	}

	return nil;
}

func (piece *Knight) Move(mov Vec2) error {
	return nil;
}

func NewKnight(x, y int32, color bool, renderer *sdl.Renderer) (*Knight, error) {
	pos := Vec2{
		X: x,
		Y: y,
	};

	var texture *sdl.Texture;
	var err error;

	if (color) {
		texture, err = img.LoadTexture(renderer, "Assets/white_knight.png");
	} else {
		texture, err = img.LoadTexture(renderer, "Assets/black_knight.png");
	}

	if (err != nil) {
		return nil, err;
	}

	return &Knight{
		pos: pos,
		color: color,
		value: 3,
		texture: texture,
	}, nil;
}



package piece

import (
	"github.com/veandco/go-sdl2/sdl"	
	"github.com/veandco/go-sdl2/img"
	"chess/globals"
	"errors"
)

type King struct {
	pos			Vec2
	color		bool // true = white; false = black
	value		uint8	
	texture		*sdl.Texture
}

func (piece *King) getMoves(board []Piece) []Vec2 {
	moves := make([]Vec2, 0);	

	for i := int32(-1); i <= 1; i++ {
		for j := int32(-1); j <= 1; j++ {
			if (i == 0 && j == 0) {
				continue;
			}
			
			pos := Vec2 {
				X: piece.pos.X + i,
				Y: piece.pos.Y + j,
			};

			posPiece, err := GetPiece(board, pos);
			isCheck, _ := resultsInCheck(board, pos, piece.pos);
			if (isCheck) {
				continue;
			}

			if (posPiece == nil && err == nil) {
				moves = append(moves, pos);
			}
		}
	}

	return moves;
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



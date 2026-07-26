package piece

import (
	"github.com/veandco/go-sdl2/sdl"	
	"github.com/veandco/go-sdl2/img"
	"chess/globals"
	"errors"
)

type Bishop struct {
	pos			Vec2
	color		bool // true = white; false = black
	value		uint8	
	texture		*sdl.Texture
}

func (piece *Bishop) getMoves(board []Piece) []Vec2 {
	moves := make([]Vec2, 0);

	for xDiff := int32(-1); xDiff <= 1; xDiff += 2 {
		for yDiff := int32(-1); yDiff <= 1; yDiff += 2 {
			// run until we explicitly break
			currPos := piece.pos;	
			for true {
				currPos = Vec2 {
					X: currPos.X + xDiff,
					Y: currPos.Y + yDiff,
				}

				blockingPiece, err := GetPiece(board, currPos);
				isCheck, _ := resultsInCheck(board, currPos, piece.pos);
				if (isCheck) {
					continue;
				}


				if (err == nil && blockingPiece != nil && blockingPiece.GetColor() != piece.color) {
					moves = append(moves, currPos);
					break;
				} else if (err != nil || blockingPiece != nil) {
					break;
				} else {
					moves = append(moves, currPos);
				}
			}
		}
	}

	return moves;
}

func (piece *Bishop) GetValue() uint8 {
	return piece.value;
}

func (piece *Bishop) GetColor() bool {
	return piece.color;
}

func (piece *Bishop) Draw(renderer *sdl.Renderer) error {
	dst := sdl.Rect{
		X: int32(piece.pos.X * globals.CellSize + 10),
		Y: int32(piece.pos.Y * globals.CellSize + 10),
		W: 80,
		H: 80,
	};

	renderer.Copy(piece.texture, nil, &dst);

	return nil;	
}

func (piece *Bishop) DrawMoves(board []Piece, renderer *sdl.Renderer) error {
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

func (piece *Bishop) Move(mov Vec2) error {
	return nil;
}

func NewBishop(x, y int32, color bool, renderer *sdl.Renderer) (*Bishop, error) {
	pos := Vec2{
		X: x,
		Y: y,
	};

	var texture *sdl.Texture;
	var err error;

	if (color) {
		texture, err = img.LoadTexture(renderer, "Assets/white_bishop.png");
	} else {
		texture, err = img.LoadTexture(renderer, "Assets/black_bishop.png");
	}

	if (err != nil) {
		return nil, err;
	}

	return &Bishop{
		pos: pos,
		color: color,
		value: 3,
		texture: texture,
	}, nil;
}



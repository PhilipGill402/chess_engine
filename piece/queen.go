package piece

import (
	"github.com/veandco/go-sdl2/sdl"	
	"github.com/veandco/go-sdl2/img"
	"chess/globals"
	"slices"
	"errors"
)

type Queen struct {
	pos			Vec2
	color		bool // true = white; false = black
	value		uint8	
	texture		*sdl.Texture
}

func (piece *Queen) getMoves(board []Piece) []Vec2 {
	moves := make([]Vec2, 0);
	
	// get diagonal moves
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
				if (err != nil) {
					break;
				}

				isCheck, _ := resultsInCheck(board, currPos, piece.pos);
				if (isCheck) {
					continue;
				}
				
				if (err == nil && blockingPiece != nil && blockingPiece.GetColor() != piece.color) {
					moves = append(moves, currPos);
					break;
				} else if (blockingPiece != nil) {
					break;
				} else {
					moves = append(moves, currPos);
				}
			}
		}
	}
	
	// get lateral moves
	for i := 0; i < 2; i++ {
		for dir := int32(-1); dir < 2; dir += 2 {
			pos := piece.pos;
			for true {
				if (i == 0) {
					pos.X += dir;					
				} else {
					pos.Y += dir;
				}

				blockingPiece, err := GetPiece(board, pos);
				if (err != nil) {
					break;
				}

				isCheck, _ := resultsInCheck(board, pos, piece.pos);
				if (isCheck) {
					continue;
				}

				// we reached another piece or went out of bounds so break loop	
				if (err == nil && blockingPiece != nil && blockingPiece.GetColor() != piece.color) {
					moves = append(moves, pos);
					break;
				} else if (blockingPiece != nil) {
					break;
				} else {
					moves = append(moves, pos);
				}
			}	
		}
	}

	return moves
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

func (piece *Queen) Move(board []Piece, move Vec2) {
	moves := piece.getMoves(board);
	if (slices.Contains(moves, move)) {	
		board[move.Y * 8 + move.X] = piece;
		board[piece.pos.Y * 8 + piece.pos.X] = nil;
		piece.pos = move;
	}
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



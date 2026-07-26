package piece

import (
	"github.com/veandco/go-sdl2/sdl"	
	"github.com/veandco/go-sdl2/img"
	"chess/globals"
	"slices"
	"errors"
)

type Rook struct {
	pos			Vec2
	color		bool // true = white; false = black
	value		uint8	
	texture		*sdl.Texture
}

func (piece *Rook) getMoves(board []Piece) []Vec2 {
	moves := make([]Vec2, 0);

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
				isCheck, _ := resultsInCheck(board, pos, piece.pos);
				if (isCheck) {
					continue;
				}

				// we reached another piece or went out of bounds so break loop	
				if (err == nil && blockingPiece != nil && blockingPiece.GetColor() != piece.color) {
					moves = append(moves, pos);
					break;
				} else if (err != nil || blockingPiece != nil) {
					break;
				} else {
					moves = append(moves, pos);
				}
			}	
		}
	}

	return moves;
}

func (piece *Rook) GetValue() uint8 {
	return piece.value;
}

func (piece *Rook) GetColor() bool {
	return piece.color;
}

func (piece *Rook) Draw(renderer *sdl.Renderer) error {
	dst := sdl.Rect{
		X: int32(piece.pos.X * globals.CellSize + 10),
		Y: int32(piece.pos.Y * globals.CellSize + 10),
		W: 80,
		H: 80,
	};

	renderer.Copy(piece.texture, nil, &dst);

	return nil;	
}

func (piece *Rook) DrawMoves(board []Piece, renderer *sdl.Renderer) error {
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

func (piece *Rook) Move(board []Piece, move Vec2) {
	moves := piece.getMoves(board);
	if (slices.Contains(moves, move)) {	
		board[move.Y * 8 + move.X] = piece;
		board[piece.pos.Y * 8 + piece.pos.X] = nil;
		piece.pos = move;
	}
}

func NewRook(x, y int32, color bool, renderer *sdl.Renderer) (*Rook, error) {
	pos := Vec2{
		X: x,
		Y: y,
	};

	var texture *sdl.Texture;
	var err error;

	if (color) {
		texture, err = img.LoadTexture(renderer, "Assets/white_rook.png");
	} else {
		texture, err = img.LoadTexture(renderer, "Assets/black_rook.png");
	}

	if (err != nil) {
		return nil, err;
	}

	return &Rook{
		pos: pos,
		color: color,
		value: 5,	
		texture: texture,
	}, nil;
}



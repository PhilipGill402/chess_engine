package piece

import (
	"github.com/veandco/go-sdl2/sdl"	
	"github.com/veandco/go-sdl2/img"
	"chess/globals"
	"errors"
)

type Pawn struct {
	pos			Vec2
	hasMoved	bool
	color		bool // true = white; false = black
	value		uint8	
	texture		*sdl.Texture
}

func (piece *Pawn) getMoves(board []Piece) []Vec2 {
	moves := make([]Vec2, 0);	
	
	// forward movement
	for diff := int32(1); diff <= 2; diff++ {
		if (piece.hasMoved && diff == 2) {
			break;
		}
		
		adjDiff := diff;
		if (piece.color == true) {
			adjDiff *= -1;	
		}

		frontPos := Vec2 {
			X: piece.pos.X,
			Y: piece.pos.Y + adjDiff,
		};
		
		frontPiece, err := GetPiece(board, frontPos);
		isCheck, _ := resultsInCheck(board, frontPos, piece.pos);
		if (isCheck) {
			continue;
		}

		if (frontPiece == nil && err == nil) {
			moves = append(moves, frontPos);
		}
	}

	// capturing
	for xDiff := int32(-1); xDiff <= 1; xDiff += 2 {
		yDiff := int32(1);	
		if (piece.color == true) {
			yDiff = -1;	
		}

		pos := Vec2 {
			X: piece.pos.X + xDiff,
			Y: piece.pos.Y + yDiff,
		}

		capturedPiece, err := GetPiece(board, pos);
		if (capturedPiece != nil && err == nil && capturedPiece.GetColor() != piece.color) {
			moves = append(moves, pos);	
		}
	}

	return moves;
}

func (piece *Pawn) GetValue() uint8 {
	return piece.value;
}

func (piece *Pawn) GetColor() bool {
	return piece.color;
}

func (piece *Pawn) Draw(renderer *sdl.Renderer) error {
	dst := sdl.Rect{
		X: int32(piece.pos.X * globals.CellSize  + 10),
		Y: int32(piece.pos.Y * globals.CellSize + 10),
		W: 80,
		H: 80,
	};

	renderer.Copy(piece.texture, nil, &dst);

	return nil;	
}

func (piece *Pawn) DrawMoves(board []Piece, renderer *sdl.Renderer) error {
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
		value: 1,	
		texture: texture,
	}, nil;
}



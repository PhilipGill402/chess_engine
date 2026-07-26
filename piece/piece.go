package piece

import (
	"github.com/veandco/go-sdl2/sdl"
	"errors"
)

type Vec2 struct {
	X	int32
	Y	int32
}

type Piece interface {
	GetValue()											uint8
	GetColor()											bool
	Draw(renderer *sdl.Renderer)						error
	DrawMoves(board []Piece, renderer *sdl.Renderer)	error
	Move(move Vec2) 									error
}

func FillCircle(renderer *sdl.Renderer, centerX, centerY, radius int32) error {
	renderer.SetDrawColor(128, 128, 128, 255) // black

	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if x*x+y*y <= radius*radius {
				if err := renderer.DrawPoint(centerX+x, centerY+y); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func GetPiece(board []Piece, pos Vec2) (Piece, error) {
	if (pos.X < 0 || pos.Y < 0 || pos.X > 7 || pos.Y > 7) {
		return nil, errors.New("Index out of range");
	}

	return board[pos.Y * 8 + pos.X], nil;
}

func isCheck(board []Piece) (bool, error) {
	kingIdx := -1;
	for i := 0; i < 64; i++ {
		piece := board[i];
		if (piece == nil) {
			continue;
		}

		// found king
		if (piece.GetValue() == 255) {
			kingIdx = i;
			break;
		}
	}

	if (kingIdx == -1) {
		return true, errors.New("No king found");
	}

	king := board[kingIdx];
	kingPos := Vec2 {
		X: int32(kingIdx % 8),
		Y: int32(kingIdx / 8),
	};
	
	// check lateral lines
	for i := 0; i < 2; i++ {
		for dir := int32(-1); dir < 2; dir += 2 {
			pos := kingPos;
			for true {
				if (i == 0) {
					pos.X += dir;					
				} else {
					pos.Y += dir;
				}

				blockingPiece, err := GetPiece(board, pos);
				
				// we went out of bounds so break
				if (err != nil) {
					break;
				} else if (blockingPiece == nil) {
					continue;
				} else if (blockingPiece.GetColor() != king.GetColor() && (blockingPiece.GetValue() == 9 || blockingPiece.GetValue() == 5)) {
					// we hit either a queen or a rook of the opposing color so we are in check	
					return true, nil;
				} else {
					break;
				}
			}
		}
	}
	
	// check diagonals
	for xDiff := int32(-1); xDiff <= 1; xDiff += 2 {
		for yDiff := int32(-1); yDiff <= 1; yDiff += 2 {
			// run until we explicitly break
			currPos := kingPos;	
			for true {
				currPos = Vec2 {
					X: currPos.X + xDiff,
					Y: currPos.Y + yDiff,
				}

				blockingPiece, err := GetPiece(board, currPos);
				// we went out of bounds so break
				if (err != nil) {
					break;
				} else if (blockingPiece == nil) {
					continue;
				} else if (blockingPiece.GetColor() != king.GetColor() && (blockingPiece.GetValue() == 9 || blockingPiece.GetValue() == 3)) {
					// we hit either a queen or a bishop of the opposing color so we are in check	
					return true, nil;
				} else {
					break;
				}
			}
		}
	}
	
	// check for knights
	for i := 0; i < 2; i++ {
		for diff := int32(-2); diff < 3; diff += 4 {
			for j := int32(-1); j < 2; j += 2 {
				pos := kingPos;
				
				if (i == 0) {
					pos.X += diff;
					pos.Y += j;
				} else {
					pos.X += j;
					pos.Y += diff;
				}
				
				piece, _ := GetPiece(board, pos);
				if (piece != nil && piece.GetColor() != king.GetColor() && piece.GetValue() == 4) {
					return true, nil;
				}
			}
		}
	}

	return false, nil;
}

func resultsInCheck(board []Piece, dst Vec2, src Vec2) (bool, error) {
	newBoard := make([]Piece, 64);
	copy(newBoard, board);
	
	srcPiece, err := GetPiece(board, src);
	if (err != nil) {
		check, _ := isCheck(board);
		return check, err;
	}

	if (dst.Y < 0 || dst.Y > 7) {
		check, _ := isCheck(board);
		return check, errors.New("Index out of range");
	} else if (dst.X < 0 || dst.X > 7) {
		check, _ := isCheck(board);
		return check, errors.New("Index out of range");
	}

	newBoard[dst.Y * 8 + dst.X] = srcPiece;
	newBoard[src.Y * 8 + src.X] = nil;

	return isCheck(newBoard);
}



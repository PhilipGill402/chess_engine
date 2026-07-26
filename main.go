package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"chess/piece"
	"chess/globals"
)


func initGameBoard(renderer *sdl.Renderer) ([]piece.Piece, error) {
	board := make([]piece.Piece, 0);
	
	// init first row
	for i := 0; i < 8; i++ {
		var newPiece piece.Piece;
		var err error;

		if (i == 0 || i == 7) {
			newPiece, err = piece.NewRook(int32(i), 0, false, renderer);
		} else if (i == 1 || i == 6) {
			newPiece, err = piece.NewKnight(int32(i), 0, false, renderer);
		} else if (i == 2 || i == 5) {
			newPiece, err = piece.NewBishop(int32(i), 0, false, renderer);
		} else if (i == 3) {
			newPiece, err = piece.NewQueen(int32(i), 0, false, renderer);
		} else {
			newPiece, err = piece.NewKing(int32(i), 0, false, renderer);	
		}

		if (err != nil) {
			return nil, err;
		}

		board = append(board, newPiece);
	}

	// init first pawn row
	for i := 0; i < 8; i++ {
		pawn, err := piece.NewPawn(int32(i), int32(1), false, renderer);
		if (err != nil) {
			return nil, err;
		}

		board = append(board, pawn);
	}

	// init middle of board
	for i := 0; i < 32; i++ {
		board = append(board, nil);
	}
	
	// TESTING
	king, _ := piece.NewBishop(int32(2), int32(4), false, renderer);
	board[34] = king;
	//pawn, _ := piece.NewPawn(int32(3), int32(3), false, renderer);
	//board[27] = pawn;

	// init second pawn row
	for i := 0; i < 8; i++ {
		pawn, err := piece.NewPawn(int32(i), int32(6), true, renderer);
		if (err != nil) {
			return nil, err;
		}

		board = append(board, pawn);
	}

	// init first row
	for i := 0; i < 8; i++ {
		var newPiece piece.Piece;
		var err error;

		if (i == 0 || i == 7) {
			newPiece, err = piece.NewRook(int32(i), int32(7), true, renderer);
		} else if (i == 1 || i == 6) {
			newPiece, err = piece.NewKnight(int32(i), int32(7), true, renderer);
		} else if (i == 2 || i == 5) {
			newPiece, err = piece.NewBishop(int32(i), int32(7), true, renderer);
		} else if (i == 3) {
			newPiece, err = piece.NewQueen(int32(i), int32(7), true, renderer);
		} else {
			newPiece, err = piece.NewKing(int32(i), int32(7), true, renderer);	
		}

		if (err != nil) {
			return nil, err;
		}

		board = append(board, newPiece);
	}

	return board, nil;
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err);
	}
	defer sdl.Quit();

	window, err := sdl.CreateWindow("Chess", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, globals.WinWidth, globals.WinHeight, sdl.WINDOW_SHOWN);
	if (err != nil) {
		panic(err);
	}
	defer window.Destroy();

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED);
	if (err != nil) {
		panic(err);
	}
	renderer.Clear();

	board, err := initGameBoard(renderer);
	if (err != nil) {
		panic(err);
	}
	
	var selectedPiece piece.Piece;
	running := true;
	mousePos := piece.Vec2{ 0, 0 };
	for running {
		drawBoard(renderer);
		err := drawPieces(renderer, board);
		if (err != nil) {
			fmt.Println("Error:", err);
		}

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
				case *sdl.QuitEvent:
					running = false;
					break;
				case *sdl.MouseMotionEvent:
					mousePos.X = e.X / globals.CellSize;
					mousePos.Y = e.Y / globals.CellSize;
					break;
				case *sdl.MouseButtonEvent:
					if (e.Type == sdl.MOUSEBUTTONUP) {
						if (selectedPiece != nil) {
							pos := piece.Vec2 {
								X: mousePos.X,
								Y: mousePos.Y,
							};
							selectedPiece.Move(board, pos);
							selectedPiece = nil;
						}	
						selectedPiece = board[mousePos.Y * 8 + mousePos.X];
					}

					break;
			}
		}

		if (selectedPiece != nil) {
			selectedPiece.DrawMoves(board, renderer);
		}

		renderer.Present();
	}
}

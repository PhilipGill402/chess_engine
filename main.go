package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"chess/piece"
)

const (
	winWidth 	= 800
	winHeight 	= 800
	cellSize	= winWidth / 8
)

func initGameBoard(renderer *sdl.Renderer) ([]piece.Piece, error) {
	board := make([]piece.Piece, 64);
	
	// init first row
	for i := 0; i < 8; i++ {
		var newPiece piece.Piece;
		var err error;

		if (i == 0 || i == 7) {
			newPiece, err = piece.NewRook(int32(i * cellSize), 0, false, renderer);
		} else if (i == 1 || i == 6) {
			newPiece, err = piece.NewKnight(int32(i * cellSize), 0, false, renderer);
		} else if (i == 2 || i == 5) {
			newPiece, err = piece.NewBishop(int32(i * cellSize), 0, false, renderer);
		} else if (i == 3) {
			newPiece, err = piece.NewQueen(int32(i * cellSize), 0, false, renderer);
		} else {
			newPiece, err = piece.NewKing(int32(i * cellSize), 0, false, renderer);	
		}

		if (err != nil) {
			return nil, err;
		}

		board = append(board, newPiece);
	}

	// init first pawn row
	for i := 0; i < 8; i++ {
		pawn, err := piece.NewPawn(int32(i * cellSize), int32(1 * cellSize), false, renderer);
		if (err != nil) {
			return nil, err;
		}

		board = append(board, pawn);
	}

	// init middle of board
	for i := 0; i < 32; i++ {
		board = append(board, nil);
	}

	// init second pawn row
	for i := 0; i < 8; i++ {
		pawn, err := piece.NewPawn(int32(i * cellSize), int32(6 * cellSize), true, renderer);
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
			newPiece, err = piece.NewRook(int32(i * cellSize), int32(7 * cellSize), true, renderer);
		} else if (i == 1 || i == 6) {
			newPiece, err = piece.NewKnight(int32(i * cellSize), int32(7 * cellSize), true, renderer);
		} else if (i == 2 || i == 5) {
			newPiece, err = piece.NewBishop(int32(i * cellSize), int32(7 * cellSize), true, renderer);
		} else if (i == 3) {
			newPiece, err = piece.NewQueen(int32(i * cellSize), int32(7 * cellSize), true, renderer);
		} else {
			newPiece, err = piece.NewKing(int32(i * cellSize), int32(7 * cellSize), true, renderer);	
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

	window, err := sdl.CreateWindow("Chess", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN);
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
	

	running := true;
	for running {
		drawBoard(renderer);
		err := drawPieces(renderer, board);
		if (err != nil) {
			fmt.Println("Error:", err);
		}

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false;
				break;
			}
		}

		renderer.Present();
	}
}

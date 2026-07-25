package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"chess/piece"
	"chess/globals"
)

func drawBoard(renderer *sdl.Renderer) {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			rect := sdl.Rect { int32(x * globals.CellSize), int32(y * globals.CellSize), int32(globals.CellSize), int32(globals.CellSize) };
			
			if ((x + y) % 2 == 0) {
				renderer.SetDrawColor(255, 255, 255, 255);	
			} else {
				renderer.SetDrawColor(0, 0, 0, 255);
			}
			
			renderer.FillRect(&rect);
		}
	}
}

func drawPieces(renderer *sdl.Renderer, board []piece.Piece) error {
	for _, item := range board {
		if (item != nil) {
			err := item.Draw(renderer);
			if (err != nil) {
				return err;
			}
		}
	}

	return nil;
}



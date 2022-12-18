package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

// bitboard represents 64 squares of a Chess board are occupied or not,
// each bit represents a square, the most left bit is square A0, the most right
// is square H8
type bitboard uint64

func main() {
	// following code calculates rank attack ray (similar to example in
	// https://www.chessprogramming.org/Hyperbola_Quintessence)
	// example rank with the Rook and some occupied squares: "1 1 0 R 0 0 1 1",
	// result rank attack targets:                           "0 1 1 0 1 1 1 0"
	const squareD2 = 11 // square A1, A2, .. H8 is 0, 1, .. 63
	tmp, _ := strconv.ParseUint(`0000000011010011000000000000000000000000000000000000000000000000`, 2, 64)
	occupancy := bitboard(tmp)
	fmt.Println("occupancy rank 2", occupancy.Draw())
	rook := bitboard(uint64(1) << (63 - squareD2))
	fmt.Println("slider (the Rook on D2)", rook.Draw())
	fmt.Println("o-r", (occupancy - rook).Draw()) // clear the slider
	oSub2r := occupancy - 2*rook
	fmt.Println(`borrow 1 from the first blocker on the left: o-2r`, oSub2r.Draw())
	OSub2R := (occupancy.Reversed() - 2*rook.Reversed()).Reversed()
	fmt.Println(`borrow 1 from the first blocker on the right: (o'-2r')'`, OSub2R.Draw())
	attackSet := oSub2r ^ OSub2R
	fmt.Println("attackSet", attackSet.Draw())
}

func (b bitboard) Reversed() bitboard {
	return bitboard(bits.Reverse64(uint64(b)))
}

func (b bitboard) CheckOccupied(square int) bool {
	return (bits.RotateLeft64(uint64(b), square+1) & 1) == 1
}

// Draw returns visual representation of the bitboard useful for debugging.
func (b bitboard) Draw() string {
	s := "\n |A B C D E F G H\n"
	for r := 8; r >= 1; r-- {
		s += fmt.Sprintf("%v|", r)
		for f := 0; f < 8; f++ { // file A to H
			sq := (r-1)*8 + f
			if b.CheckOccupied(sq) {
				s += "1"
			} else {
				s += "0"
			}
			s += " "
		}
		s += "\n"
	}
	return s
}

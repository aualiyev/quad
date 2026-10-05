package piscine

import "github.com/01-edu/z01"

func Quad(x, y int, typeName rune) {
	switch typeName {
	case 'A':
		QuadA(x, y)
	case 'B':
		QuadB(x, y)
	case 'C':
		QuadC(x, y)
	case 'D':
		QuadD(x, y)
	case 'E':
		QuadE(x, y)
	default:
		QuadA(x, y)
	}
}

func QuadA(x, y int) {
	baseQuad(x, y, 'o', 'o', 'o', 'o', '|', '-')
}

func QuadB(x, y int) {
	baseQuad(x, y, '/', '\\', '\\', '/', '*', '*')
}

func QuadC(x, y int) {
	baseQuad(x, y, 'A', 'A', 'C', 'C', 'B', 'B')
}

func QuadD(x, y int) {
	baseQuad(x, y, 'A', 'C', 'A', 'C', 'B', 'B')
}

func QuadE(x, y int) {
	baseQuad(x, y, 'A', 'C', 'C', 'A', 'B', 'B')
}

func baseQuad(x, y int, cornerTL, cornerTR, cornerBL, cornerBR, sideSym, horizontalSym rune) {
	if x <= 0 || y <= 0 {
		return
	}

	for vertical := 0; vertical < y; vertical++ {
		for horizontal := 0; horizontal < x; horizontal++ {

			isTop := vertical == 0
			isBottom := vertical == y-1
			isLeft := horizontal == 0
			isRight := horizontal == x-1

			isCornerTopLeft := isTop && isLeft
			isCornerTopRight := isTop && isRight
			isCornerBottomLeft := isBottom && isLeft
			isConerBottomRight := isBottom && isRight
			isCorner := isCornerTopLeft || isCornerTopRight || isCornerBottomLeft || isConerBottomRight

			if isCorner {
				if isCornerTopLeft {
					z01.PrintRune(cornerTL)
				} else if isCornerTopRight {
					z01.PrintRune(cornerTR)
				} else if isCornerBottomLeft {
					z01.PrintRune(cornerBL)
				} else if isConerBottomRight {
					z01.PrintRune(cornerBR)
				}
			} else if isLeft || isRight {
				z01.PrintRune(sideSym)
			} else if isTop || isBottom && !isLeft && !isRight {
				z01.PrintRune(horizontalSym)
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}

	z01.PrintRune('\n')
}

package main

import piscine "piscine"

func main() {
	typeName := 'E'

	piscine.Quad(5, 3, typeName)
	piscine.Quad(5, 1, typeName)
	piscine.Quad(1, 1, typeName)
	piscine.Quad(1, 0, typeName)
	piscine.Quad(0, 1, typeName)
	piscine.Quad(1, 5, typeName)
	piscine.Quad(1, -5, typeName)
	piscine.Quad(-1, 5, typeName)
}

package main

import (
	//"fmt"
	"math"
)

/*func main() {
	fmt.Println(CalcSquare(10.0, SidesTriangle))
	fmt.Println(CalcSquare(10.0, SidesSquare))
	fmt.Println(CalcSquare(10.0, SidesCircle))
	fmt.Println(CalcSquare(10.0, iint(5)))
} */

type iint int

const SidesTriangle iint = 3
const SidesSquare iint = 4
const SidesCircle iint = 0

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum iint) float64 {

	switch sidesNum {
	case SidesTriangle:
		return math.Sqrt(1.5 * sideLen * (0.5 * sideLen) * 0.5 * sideLen * 0.5 * sideLen)
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	default:
		return 0
	}
}

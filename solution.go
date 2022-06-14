package square

import (
	"fmt"
	"math"
)

type kolvo int

var SidesTriangle kolvo = 3
var SidesSquare kolvo = 4
var SidesCircle kolvo = 0

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum kolvo) float64 {
	var s float64
	var hil float64
	switch {
	case SidesTriangle == 3:
		hil = math.Sqrt(3)
		s = (sideLen * sideLen * hil) / 4

	case SidesSquare == 4:
		s = sideLen * sideLen

	case SidesCircle == 0:
		s = math.Pi * (sideLen * sideLen)

	default:
		var v string = "error"
		fmt.Print(v)
	}
	return s
}

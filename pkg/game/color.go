package game

// Color : Int constant mapping
const (
	Red        = 0
	Blue       = 1
	Green      = 2
	Pink       = 3
	Orange     = 4
	Yellow     = 5
	Black      = 6
	White      = 7
	Purple     = 8
	Brown      = 9
	Cyan       = 10
	Lime       = 11
	Maroon     = 12
	Rose       = 13
	Banana     = 14
	Gray       = 15
	Tan        = 16
	Coral      = 17
	Watermelon = 18
	Chocolate  = 19
	Skyblue    = 20
	Beige      = 21
	Hotpink    = 22
	Turquoise  = 23
	Lilac      = 24
	Rainbow    = 25
	Azure      = 26
)

// ColorStrings for lowercase, possibly for translation if needed
var ColorStrings = map[string]int{
	"red":        Red,
	"blue":       Blue,
	"green":      Green,
	"pink":       Pink,
	"orange":     Orange,
	"yellow":     Yellow,
	"black":      Black,
	"white":      White,
	"purple":     Purple,
	"brown":      Brown,
	"cyan":       Cyan,
	"lime":       Lime,
	"maroon":     Maroon,
	"rose":       Rose,
	"banana":     Banana,
	"gray":       Gray,
	"tan":        Tan,
	"coral":      Coral,
	"watermelon": Watermelon,
	"chocolate":  Chocolate,
	"slyblue":    Skyblue,
	"beige":      Beige,
	"hotpink":    Hotpink,
	"turquoise":  Turquoise,
	"lilac":      Lilac,
	"rainbow":    Rainbow,
	"azure":      Azure,
}

// GetColorStringForInt does what it sounds like
func GetColorStringForInt(colorint int) string {
	for str, idx := range ColorStrings {
		if idx == colorint {
			return str
		}
	}
	return ""
}

// IsColorString determines if a string is actually one of our colors
func IsColorString(test string) bool {
	_, ok := ColorStrings[test]
	return ok
}

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
	Watermelon = 12
	Chocolate  = 13
	Skyblue    = 14
	Beige      = 15
	Hotpink    = 16
	Turquoise  = 17
	Lilac      = 18
	Lightpink  = 19
	Rainbow    = 20
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
	"watermelon": Watermelon,
	"chocolate":  Chocolate,
	"skyblue":    Skyblue,
	"beige":      Beige,
	"hotpink":    Hotpink,
	"turquoise":  Turquoise,
	"lilac":      Lilac,
	"lightpink":  Lightpink,
	"rainbow":    Rainbow,
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

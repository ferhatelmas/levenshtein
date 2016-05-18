package levenshtein

// D is the levenshtein distance calculator interface
type D interface {
	// Dist calculates levenshtein distance between two utf-8 encoded strings
	Dist(string, string) int
}

// New creates a new levenshtein distance calculator where indel is increment/deletion cost
// and sub is the substitution cost.
func New(indel, sub int) D {
	return &calculator{indel, sub}
}

type calculator struct {
	indel, sub int
}

// https://en.wikibooks.org/wiki/Algorithm_Implementation/Strings/Levenshtein_distance#C
func (c *calculator) Dist(s1, s2 string) int {
	r1, r2 := []rune(s1), []rune(s2)
	l := len(r1)
	m := make([]int, l+1)
	for y := 1; y <= l; y++ {
		m[y] = y * c.indel
	}
	var lastdiag int
	for x, rx := range r2 {
		m[0], lastdiag = (x+1)*c.indel, x*c.indel
		for y, ry := range r1 {
			m[y+1], lastdiag = min3(m[y+1]+c.indel, m[y]+c.indel, lastdiag+c.subCost(rx, ry)), m[y+1]
		}
	}
	return m[l]
}

func (c *calculator) subCost(r1, r2 rune) int {
	if r1 == r2 {
		return 0
	}
	return c.sub
}

func min3(a, b, c int) int {
	return min(a, min(b, c))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var defaultCalculator = New(1, 1)

// Dist is a convenience function for a levenshtein distance calculator with equal costs.
func Dist(s1, s2 string) int {
	return defaultCalculator.Dist(s1, s2)
}

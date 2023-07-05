package utils

const (
	fontsizeToPixel = 1.618
)

func CalculateFontSize(width, height float64, symbolsCount int) int {
	minSize := height / fontsizeToPixel
	minSize = minFloat64(minSize, (fontsizeToPixel*width)/float64(symbolsCount))
	return int(minSize)
}

func minFloat64(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}

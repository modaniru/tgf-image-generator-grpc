package utils

const (
	fontsizeToPixel = 1.618
)

// Вычисляет подходящий размер шрифта по к-ву символов и нужной площади
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

// Расчитывает к-во строк по заданному количеству элементов в строке
func CalculateRowsCount(count, countInRows int) int {
	res := count / countInRows
	if res*countInRows != count {
		res++
	}
	return res
}

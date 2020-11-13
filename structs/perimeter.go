package structs

// Perimeter returns 2*(a+b)
func Perimeter(width float64, height float64) float64 {
    return 2 * (width + height)
}

// Area returns a*b
func Area(width, height float64) float64 {
    return width * height
}
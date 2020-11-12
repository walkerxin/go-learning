package iteration

// Repeat the taked character return five
func Repeat(character string, frequency int) string {
	var repeated string
	for i:=0; i<frequency; i++ {
		repeated += character
	}
	return repeated
}
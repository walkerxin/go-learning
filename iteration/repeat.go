package iteration

// Repeat the taked character return five
func Repeat(character string) (string) {
	var repeated string
	for i:=0; i<5; i++ {
		repeated += character
	}
	return repeated
}
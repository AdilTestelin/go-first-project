package iteration

func RepeatChar(character string, iterationNumber int) string {
	var repeated string
	for i := 0; i < iterationNumber; i++ {
		repeated += character
	}
	return repeated
}

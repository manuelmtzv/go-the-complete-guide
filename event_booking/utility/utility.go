package utility

func ValidateEvery[T any](callback func(T) bool, values ...T) bool {
	valid := true

	for _, value := range values {
		if !callback(value) {
			valid = false
			break
		}
	}

	return valid
}

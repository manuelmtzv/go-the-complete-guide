package conversion

import "strconv"

func StringsToFloat(strings ...string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for index, stringValue := range strings {
		parsedFloat, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			return nil, err
		}

		floats[index] = parsedFloat
	}

	return floats, nil
}

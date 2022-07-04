package temperature

import "fmt"

func Conv(temp float64, symbol string) (float64, error) {

	if symbol == "F" || symbol == "f" {
		return tempFtoC(temp), nil

	} else if symbol == "C" || symbol == "c" {
		return tempCtoF(temp), nil
	}
	return 0, fmt.Errorf("Unrecognized symbol : %s, want: F, f, C or c\n", symbol)

}

func tempFtoC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func tempCtoF(c float64) float64 {
	return (c * 1.8) + 32
}

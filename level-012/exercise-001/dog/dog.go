// Package dog calculates the Dog's age using the human's age.
package dog

import (
	"errors"
	"math"
)

//Years calculates the Dog's age using the human's age.
//The zero value and error will be returned in the case
//of negative humanYears parameter.
func Years(humanYears int) (int, error) {
	if humanYears < 0 {
		return 0, errors.New("hunan years should be positive")
	}
	dogYears := float64(humanYears / 7)
	return int(math.Round(dogYears)), nil
}

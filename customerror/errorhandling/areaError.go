package errorhandling

import (
	"fmt"
)

type areaError struct {
	radius float64
	errMsz string
}

func New(radius float64) error {
	return &areaError{radius, "Area not less than zero"}
}

func (a *areaError) Error() string {
	return fmt.Sprintf("radius: %.2f %s", a.radius, a.errMsz)
}

package services

import (
	"errors"
)

func Convert(c Conversion) (float64, error) {
	switch c.Type {
	case "km-to-miles":
		return c.Value * 0.621371, nil
	case "kg-to-pounds":
		return c.Value * 2.20462, nil
	case "c-to-f":
		return (c.Value*9/5) + 32, nil
	default:
		return 0, errors.New("unsupported conversion type")
	}
}
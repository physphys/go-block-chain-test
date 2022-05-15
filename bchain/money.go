package bchain

import "errors"

type money float64

func (m money) validate() error {
	if m < 0 {
		return errors.New("minus money")
	}

	return nil
}

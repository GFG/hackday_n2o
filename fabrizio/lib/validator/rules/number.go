package rules

import (
	"errors"
	"strconv"
)

type (
	numberRule struct {
		originalValue string
		value float32
		mandatory bool
		min float32
		max float32
	}
)

func NewNumberRule() *numberRule  {
	return &numberRule{}
}

func (nr *numberRule) SetValue(value string)  {
	nr.originalValue = value
	floatVal, _ := strconv.ParseFloat(value, 32)
	nr.value = float32(floatVal)
}

func (nr *numberRule) IsMandatory(mandatory bool)  {
	nr.mandatory = mandatory
}

func (nr *numberRule) SetMin(min float32)  {
	nr.min = min
}

func (nr *numberRule) SetMax(max float32)  {
	nr.max = max
}

func (nr *numberRule) IsValid() (bool, error) {
	if nr.mandatory && nr.originalValue == "" {
		return false, errors.New("the field is mandatory")
	}
	return true, nil
}

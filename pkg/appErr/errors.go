package apperr

import (
	"encoding/json"
	"fmt"
)

type CustomErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
	Details string `json:"details"`
}

func (c *CustomErr) Error() string {
	return fmt.Sprintf("\nerror:\n \tType: %s,\n \tCode: %d,\n \tMessage: %s,\n \tDetails: %s,\n",
		c.Type, c.Code, c.Message, c.Details)
}

func NewAppErr(code int, message, errType, details string) *CustomErr {
	return &CustomErr{
		Code:    code,
		Message: message,
		Type:    errType,
		Details: details,
	}
}

func (c *CustomErr) Log() string {
	bytes, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return c.Error()
	}
	return string(bytes)
}

// Package abstractions provides base interface abstractions for Lattice-Go.
package abstractions

import "time"

// Envelope maps standardized response envelopes.
type Envelope struct {
	TransactionURN  string `json:"transactionUrn"`
	Status          string `json:"status"`
	ResponseMessage string `json:"responseMessage"`
	ResponseKey     string `json:"responseKey"`
	Errors          []any  `json:"errors"`
	Timestamp       string `json:"timestamp"`
	Metadata        map[string]any `json:"metadata"`
	Data            any    `json:"data"`
	ReferenceURN    string `json:"referenceUrn"`
}

type BaseController struct{}

func (c *BaseController) Envelope(data any, message string, responseKey string, statusCode int) Envelope {
	status := "SUCCESS"
	if statusCode >= 400 {
		status = "FAILED"
	}
	return Envelope{
		Status:          status,
		ResponseMessage: message,
		ResponseKey:     responseKey,
		Errors:          []any{},
		Timestamp:       time.Now().Format(time.RFC3339),
		Metadata:        map[string]any{},
		Data:            data,
	}
}

func (c *BaseController) Success(data any, message string, responseKey string) Envelope {
	return c.Envelope(data, message, responseKey, 200)
}

func (c *BaseController) Created(data any, message string, responseKey string) Envelope {
	return c.Envelope(data, message, responseKey, 201)
}

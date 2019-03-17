package error

import (
	"github.com/go-errors/errors"
)

type (
	// Error is single error instance
	Error struct {
		Status   int    `json:"status,omitempty"`
		Code     string `json:"code,omitempty"`
		Detail   string `json:"detail,omitempty"`
		Trace    string `json:"trace,omitempty"`
		Priority int    `json:"-"`
	}

	// Container is container of one or more errors that will be given to response
	Container struct {
		Errors []Error `json:"errors"`
	}
)

// Error implements golang native error
func (e Error) Error() string {
	return e.Detail
}

// Wrap golang native error into our Error
func (e Error) Wrap(err error) Error {
	return Error{
		Status: e.Status,
		Code:   e.Code,
		Detail: err.Error(),
		Trace:  errors.Wrap(err, 1).ErrorStack(),
	}
}

// ProduceStackTrace generates stack trace to Trace property
func (e Error) ProduceStackTrace() Error {
	e.Trace = errors.Wrap(e, 1).ErrorStack()
	return e
}

// ToContainer converts Error to Container
func (e Error) ToContainer() Container {
	return Container{
		Errors: []Error{e},
	}
}

// Container implements golang native error
func (c Container) Error() string {
	return c.GetFirstError().Error()
}

// MergeErrors will merge list of errors into the container
func (c Container) MergeErrors(errs []Error) {
	if c.Errors != nil {
		c.Errors = errs
	} else {
		for _, v := range errs {
			c.Errors = append(c.Errors, v)
		}
	}
}

// Wrap golang native error into our Container
func (c Container) Wrap(err error) Container {
	res := Container{
		Errors: []Error{c.GetFirstError().Wrap(err)},
	}

	if len(c.Errors) > 1 {
		res.Errors = append(res.Errors, c.Errors[1:]...)
	}

	return res
}

// GetFirstError get first error inside the Container
func (c Container) GetFirstError() Error {
	if len(c.Errors) > 0 {
		return c.Errors[0]
	}

	return Error{}
}

// RemoveTraces remove trace from all errors inside the container
func (c Container) RemoveTraces() {
	for i := 0; i < len(c.Errors); i++ {
		c.Errors[i].Trace = ""
	}
}

// Len will return size of error
func (c Container) Len() int { return len(c.Errors) }

// Less will return true if error[i] has bigger priority than j
func (c Container) Less(i, j int) bool {
	return c.Errors[i].Priority > c.Errors[j].Priority
}

// Swap ...
func (c Container) Swap(i, j int) {
	c.Errors[i], c.Errors[j] = c.Errors[j], c.Errors[i]
}

// Push ...
func (c *Container) Push(x interface{}) {
	e := x.(Error)
	c.Errors = append(c.Errors, e)
}

// Pop ...
func (c *Container) Pop() interface{} {
	n := c.Len()
	item := c.Errors[n-1]
	c.Errors = c.Errors[0 : n-1]
	return item
}

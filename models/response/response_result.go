// Copyright 2020. All rights reserved.
// Author 赵路通

package response

import "fmt"

var _ error = &ErrorResult{}

// Error implements the Error interface.
func (e *ErrorResult) Error() string {
	return fmt.Sprintf("status:%d, msg:%s", e.Status, e.Msg)
}

type ErrorResult struct {
	// The custom status
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

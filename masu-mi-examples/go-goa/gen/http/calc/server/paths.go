// Code generated by goa v3.1.2, DO NOT EDIT.
//
// HTTP request path constructors for the calc service.
//
// Command:
// $ goa gen calcsvc/design

package server

import (
	"fmt"
)

// AddCalcPath returns the URL path to the calc service add HTTP endpoint.
func AddCalcPath(a int, b int) string {
	return fmt.Sprintf("/add/%v/%v", a, b)
}

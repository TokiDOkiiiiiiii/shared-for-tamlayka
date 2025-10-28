package apperror

import "strings"

var PARSE_ERROR = "unable to parse request body"
var INVALID_INPUT = "invalid request body"
var AUTHENTHECATE_ERROR = "username or password is incorrect"

func ServiceError(serviceName string) string {
	return strings.ToLower(serviceName) + " service failed"
}

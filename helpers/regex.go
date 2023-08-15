package helpers

import "regexp"

var (
	TimeRouteRegex, _   = regexp.Compile("^/time$")
	PublicRouteRegex, _ = regexp.Compile("/public")
	UsersRouteRegex, _  = regexp.Compile("^/users$")
	UserRouteRegex, _   = regexp.Compile(`^/users/(\d+)$`)
)

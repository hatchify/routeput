package routeput

import (
	"github.com/Hatch1fy/httpserve"
	"github.com/hatchify/scribe"
)

var out = scribe.New("Routeput")

// Log will log a provided route
func Log(args ...string) (h httpserve.Handler, err error) {
	var l logger
	h = l.middleware
	return
}

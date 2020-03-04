package main

import (
	"github.com/Hatch1fy/httpserve"
	"github.com/hatchify/colors"
	"github.com/hatchify/scribe"
)

var out = scribe.New("Routeput")

var (
	cyan   = colors.New(colors.Bold, colors.FGCyan)
	green  = colors.New(colors.Bold, colors.FGGreen)
	yellow = colors.New(colors.Bold, colors.FGYellow)
	red    = colors.New(colors.Bold, colors.FGRed)

	bgCyan   = colors.New(colors.Bold, colors.BGCyan)
	bgGreen  = colors.New(colors.Bold, colors.BGGreen)
	bgYellow = colors.New(colors.Bold, colors.BGYellow)
	bgRed    = colors.New(colors.Bold, colors.BGRed)
)

// Log will log a provided route
func Log(args ...string) (h httpserve.Handler, err error) {
	var l logger
	h = l.middleware
	return
}

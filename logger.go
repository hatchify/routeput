package main

import (
	"time"

	"github.com/Hatch1fy/httpserve"
	"github.com/hatchify/stopwatch"
)

type logger struct {
	sw     stopwatch.Stopwatch
	ip     string
	method string
	route  string
}

func (l *logger) middleware(ctx *httpserve.Context) (res httpserve.Response) {
	l.sw.Start()
	l.ip = ctx.Request.RemoteAddr
	l.method = ctx.Request.Method
	l.route = ctx.Request.URL.Path
	ctx.AddHook(l.hook)
	return
}

func (l *logger) hook(statusCode int, s httpserve.Storage) {
	now := time.Now()
	dur := l.sw.Stop()
	out.Notificationf("%s | %s | %s | %s | %s", now.String(), dur.String(), l.ip, l.method, l.route)
}

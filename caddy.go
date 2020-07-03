package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"
	"github.com/caddyserver/caddy/caddyhttp/httpserver"

	_ "github.com/caddy-plugins/caddy-iplimit/iplimit"
)

func main() {
	httpserver.RegisterDevDirective("iplimit", "index")
	caddymain.EnableTelemetry = false
	caddymain.Run()
}

package security

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// SetSecurityHeaders registra los headers de seguridad para todas las peticiones.
func SetSecurityHeaders() {
	beego.InsertFilter("*", beego.BeforeExec, securityHeaders)
}

// securityHeaders agrega headers HTTP de seguridad a cada respuesta.
func securityHeaders(ctx *context.Context) {
	ctx.Output.Header("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'")
	ctx.Output.Header("Cross-Origin-Resource-Policy", "cross-origin")
	ctx.Output.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
	ctx.Output.Header("Referrer-Policy", "no-referrer")
	ctx.Output.Header("Server", "")
	ctx.Output.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	ctx.Output.Header("X-XSS-Protection", "1; mode=block")
	ctx.Output.Header("X-Content-Type-Options", "nosniff")
	ctx.Output.Header("X-Frame-Options", "DENY")
	ctx.Output.Header("X-Permitted-Cross-Domain-Policies", "none")
}

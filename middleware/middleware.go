package middleware

import (
    "strings"

    "github.com/labstack/echo/v4"
)

func RealIPMiddleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            clientIP := c.Request().RemoteAddr

            if realIP := c.Request().Header.Get("X-Real-IP"); realIP != "" {
                clientIP = realIP
            } else if forwardedFor := c.Request().Header.Get("X-Forwarded-For"); forwardedFor != "" {
                clientIP = strings.Split(forwardedFor, ",")[0]
            }

            c.Set("clientIpAddress", clientIP)

            return next(c)
        }
    }
}


package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/Aidin1998/simple-go-auth/auth"
  "account-module/config"
  "account-module/logger"
)

// NewServer constructs and returns an Echo instance with all middleware & routes.
func NewServer() *echo.Echo {
  config.Init()
  logger.Init()

  e := echo.New()
  // JWT middleware stub; uses simple-go-auth to get secret
  e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
    SigningKey: []byte(auth.GetJWTSecret()),
  }))

  // Health check
  e.GET("/healthz", func(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
  })

  return e
}

func main() {
  e := NewServer()
  e.Logger.Fatal(e.Start(":" + config.Cfg.Port))
}
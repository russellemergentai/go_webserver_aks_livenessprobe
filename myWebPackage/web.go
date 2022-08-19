package webstuff

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloWebCall() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func IDWebcall(healthy *bool) {
	e := echo.New()
	e.GET("/users/:id", func(c echo.Context) error {
		// User ID from path `users/:id`
		id := c.Param("id")
		return c.String(http.StatusOK, "Hello "+id)
	})

	// http://localhost:8080/up
	e.GET("/up", func(c echo.Context) error {
		*healthy = true
		message := fmt.Sprintf("server up %t", *healthy)
		return c.String(http.StatusOK, message)
	})

	// http://localhost:8080/down
	e.GET("/down", func(c echo.Context) error {
		*healthy = false
		message := fmt.Sprintf("server down %t", *healthy)
		panic(message)
	})

	// http://localhost:8080/status
	e.GET("/status", func(c echo.Context) error {
		message := fmt.Sprintf("server status %t", *healthy)
		if *healthy {
			return c.String(http.StatusOK, message)
		}
		// do not send StatusOK if not healthy, this causes AKS livenessprobe to restart the pod
		panic(message)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

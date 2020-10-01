package main

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func server() {
	e := echo.New()

	// Static serves
	e.Static("/", "public")

	// Webhook route
	e.POST("/webhook/:id", func(c echo.Context) error {
		err := hookEvent(c.Param("id"), c)
		if err != nil {
			var webhookError webError
			webhookError.Error = err.Error()
			return c.JSON(http.StatusInternalServerError, webhookError)
		}
		var webhookResult result
		webhookResult.Message = "OK"

		return c.JSON(http.StatusOK, webhookResult)
	})

	e.HideBanner = true

	// port := os.Getenv("PORT")
	port := "1323"
	log.Info("Launching web server on port " + port)
	e.Logger.Fatal(e.Start(":" + port))
}

type webError struct {
	Error string `json:"error"`
}

type result struct {
	Message string `json:"message"`
}

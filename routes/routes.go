package routes

import (
	"crud-mahasiswa/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/mahasiswa", controller.FetchAllMahasiswa)
	e.GET("/mahasiswa/:id", controller.FetchMahasiswaByID)
	e.POST("/mahasiswa", controller.InsertMahasiswa)
	e.PUT("/mahasiswa/:id", controller.UpdateMahasiswa)
	e.DELETE("/mahasiswa/:id", controller.DeleteMahasiswa)

	return e
}

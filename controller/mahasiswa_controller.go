package controller

import (
	"crud-mahasiswa/model"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// FetchAllMahasiswa fetch all mahasiswa data from database
func FetchAllMahasiswa(c echo.Context) error {
	res, err := model.FetchMahasiswa()
	// check if error no rows in result set
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Data not found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func FetchMahasiswaByID(c echo.Context) error {
	id := c.Param("id")

	res, err := model.FetchMahasiswaByID(id)
	// check if error no rows in result set
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Data not found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func InsertMahasiswa(c echo.Context) error {
	var mahasiswa model.Mahasiswa

	err := c.Bind(&mahasiswa)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := model.CreateMahasiswa(mahasiswa.NIM, mahasiswa.Nama, mahasiswa.Umur, mahasiswa.Prodi, mahasiswa.Alamat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateMahasiswa(c echo.Context) error {
	var mahasiswa model.Mahasiswa

	id := c.Param("id")

	//convert id into int
	idConv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = c.Bind(&mahasiswa)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := model.UpdateMahasiswa(idConv, mahasiswa.NIM, mahasiswa.Nama, mahasiswa.Umur, mahasiswa.Prodi, mahasiswa.Alamat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func DeleteMahasiswa(c echo.Context) error {
	id := c.Param("id")

	//convert id into int
	idConv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := model.DeleteMahasiswa(idConv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

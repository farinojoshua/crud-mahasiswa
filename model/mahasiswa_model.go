package model

import (
	db "crud-mahasiswa/db"
)

type Mahasiswa struct {
	ID     int    `json:"id"`
	NIM    string `json:"nim"`
	Nama   string `json:"nama"`
	Umur   int    `json:"umur"`
	Prodi  string `json:"prodi"`
	Alamat string `json:"alamat"`
}

func FetchMahasiswa() (Response, error) {
	var obj Mahasiswa
	var arrobj []Mahasiswa
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM mahasiswa"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.ID, &obj.NIM, &obj.Nama, &obj.Umur, &obj.Prodi, &obj.Alamat)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = 200
	res.Message = "Berhasil Mengambil Data"
	res.Data = arrobj

	return res, nil
}

// fecth one mahasiswa by id
func FetchMahasiswaByID(id string) (Response, error) {
	var obj Mahasiswa
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM mahasiswa WHERE id = ?"

	err := con.QueryRow(sqlStatement, id).Scan(&obj.ID, &obj.NIM, &obj.Nama, &obj.Umur, &obj.Prodi, &obj.Alamat)

	if err != nil {
		return res, err
	}

	res.Status = 200
	res.Message = "Berhasil Mengambil Data"
	res.Data = obj

	return res, nil
}

// create new mahasiswa
func CreateMahasiswa(nim string, nama string, umur int, prodi string, alamat string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO mahasiswa (nim, nama, umur, prodi, alamat) VALUES (?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nim, nama, umur, prodi, alamat)
	if err != nil {
		return res, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = 200
	res.Message = "Berhasil Menambahkan Data"
	res.Data = map[string]int64{
		"last_insert_id": lastInsertID,
	}

	return res, nil
}

// update mahasiswa by id
func UpdateMahasiswa(id int, nim string, nama string, umur int, prodi string, alamat string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE mahasiswa SET nim = ?, nama = ?, umur = ?, prodi = ?, alamat = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nim, nama, umur, prodi, alamat, id)
	if err != nil {
		return res, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = 200
	res.Message = "Berhasil Mengupdate Data"
	res.Data = map[string]int64{
		"row_affected": rowAffected,
	}

	return res, nil
}

// delete mahasiswa by id
func DeleteMahasiswa(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM mahasiswa WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = 200
	res.Message = "Berhasil Menghapus Data"
	res.Data = map[string]int64{
		"row_affected": rowAffected,
	}

	return res, nil
}

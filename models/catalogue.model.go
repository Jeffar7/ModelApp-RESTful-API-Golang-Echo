package models

import (
	"net/http"

	"myapp/db"
)

type Catalogue struct {
	Id          int    `json:id`
	Model_name  string `json:model_name`
	Image_url   string `json:image_url`
	Description string `json:description`
	Schedules Schedules `json:schedules`

}
type Schedules struct {
	Id int `json:id`
	Model_id int `json:model_id`
	Schedule_date string `json:schedule_date`
}

func FetchAllCatalogue() (Response, error) {
	var obj Catalogue
	var arrobj []Catalogue
	var res Response

	con := db.CreateCon()

	sqlStatement := 
	`SELECT 
	model_catalogues.id,model_catalogues.Model_name, model_catalogues.Image_url, model_catalogues.Description,
	model_schedules.id,model_schedules.Model_id,model_schedules.Schedule_date
	FROM model_catalogues 
	INNER JOIN 
	model_schedules ON model_schedules.Model_id = model_catalogues.id`

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id,&obj.Model_name,&obj.Image_url,&obj.Description,&obj.Schedules.Id,&obj.Schedules.Model_id,&obj.Schedules.Schedule_date)

		if err != nil {
			return res, err
		}


		arrobj = append(arrobj, obj)
	}


	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj;

	return res, nil
}

func FetchAllModelById(model_id int)(Response , error){
	var obj Catalogue
	var arrobj []Catalogue
	var res Response


	con := db.CreateCon()

	sqlStatement := 
	`SELECT a.id,a.Model_name,a.Image_url,a.Description,
	b.id,b.Model_id,b.Schedule_date 
	FROM
	model_catalogues a
	INNER JOIN
	model_schedules b ON b.Model_id = a.id
	WHERE b.Model_id = ?
	`

	rows,err := con.Query(sqlStatement,model_id)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Id,&obj.Model_name,&obj.Image_url,&obj.Description,&obj.Schedules.Id,&obj.Schedules.Model_id,&obj.Schedules.Schedule_date)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = arrobj
	
	return res, nil
	
}

func StoreModelCatalogues(Model_name string, Image_url string, Description string)(Response, error){
	var res Response

	con:= db.CreateCon()

	sqlStatement := "INSERT model_catalogues (Model_name,Image_url,Description) VALUES (?, ?, ?)"


	stmt,err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Model_name, Image_url,Description)

	if err != nil {
		return res, err
	}

	lastInsertedId,err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Model created successfully"
	res.Data = map[string]int64{
		"last_inserted_id" : lastInsertedId,
	}
	
	return res, nil
}

func FetchScheduleByModelId(id int)(Response, error){
	var obj Schedules
	var arrobj []Schedules
	var res Response

	con := db.CreateCon()
	sqlStatement := `SELECT id, Model_id AS model_id, Schedule_date AS schedule_date FROM model_schedules WHERE model_id = ?`
	rows, err := con.Query(sqlStatement,id)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Id,&obj.Model_id, &obj.Schedule_date)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = arrobj

	return res, nil
}

func StoreModelSchedules(Model_id int, Schedule_date string)(Response , error){
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT model_schedules (Model_id,Schedule_date) VALUES (?,?)"

	stmt,err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Model_id,Schedule_date)

	if err != nil {
		return res, err
	}

	lastInsertedId,err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Schedule created successfully"
	res.Data = map[string]int64{
		"last_inserted_id" : lastInsertedId,
	}

	return res,nil

}

func UpdateModelCatalogues(id int, Model_name string, Image_url string, Description string)(Response , error){
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE model_catalogues SET Model_name = ?, Image_Url = ?, Description = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Model_name, Image_url, Description, id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Model updated successfully"
	res.Data = map[string]int64{
		"rows_affected" : rowsAffected,
	}

	return res, nil
}

func UpdateScheduleByScheduleID(id int, Model_id int, Schedule_date string)(Response, error){
	var res Response
	con := db.CreateCon()

	sqlStatement := "UPDATE model_schedules SET Model_id = ?, Schedule_date = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Model_id,Schedule_date,id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Schedule updated successfully"
	res.Data = map[string]int64{
		"rows_affected" : rowsAffected,
	}

	return res, nil
}

func DeleteModelSchedule(id int)(Response, error){
	var res Response

	con:= db.CreateCon()

	sqlStatement := "DELETE FROM model_schedules WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Schedule deleted successfully"
	res.Data = map[string]int64{
		"rows_affected" : rowsAffected,
	}

	return res, nil
}

func DeleteModelCatalogues(id int) (Response, error){
	var res Response

	con:= db.CreateCon()

	sqlStatement := "DELETE FROM model_catalogues WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Model deleted successfully"
	res.Data = map[string]int64{
		"rows_affected" : rowsAffected,
	}

	return res, nil
}
package service

import (
	"consumer-importer/model"
	"consumer-importer/service/util"
)

// FileService service responsible for managing File data to the postgres database
type FileService struct {
	Props *util.DbProperties
}

//Save store a File object into File database table
func (f *FileService) Save(dto *model.File) error {
	db := util.OpenConnection(f.Props)
	defer util.CloseConnection(db)

	sqlStatement := `insert into file
		(import_date,
		name,
		successful)
	values
		($1, $2, $3)
	RETURNING id;`

	return db.QueryRow(sqlStatement, dto.ImportDate, dto.Name, dto.Successful).Scan(&dto.ID)
}

//Find find a File register into the database
func (f *FileService) Find(ID int64) (*model.File, error) {
	db := util.OpenConnection(f.Props)
	defer util.CloseConnection(db)

	sqlStatement := `select
		id,
		name,
		import_date,
		successful
	from
		file
	where
		id = $1;`

	row := db.QueryRow(sqlStatement, ID)

	file := model.File{}

	err := row.Scan(&file.ID, &file.Name, &file.ImportDate, &file.Successful)

	if err != nil {
		return nil, err
	}

	return &file, nil
}

//Delete delete a file from database
func (f *FileService) Delete(ID int64) (bool, error) {
	db := util.OpenConnection(f.Props)
	defer util.CloseConnection(db)

	sqlStatement := `DELETE FROM file
		WHERE id = $1`

	result, err := db.Exec(sqlStatement, ID)
	rowsAffected, _ := result.RowsAffected()

	return rowsAffected == 1, err
}

//UpdateFileAsSuccessful update the target tile as successful in the database
func (f *FileService) UpdateFileAsSuccessful(ID int64) (bool, error) {
	db := util.OpenConnection(f.Props)
	defer util.CloseConnection(db)

	sqlStatement := `UPDATE file
		set successful = true
		WHERE id = $1`

	result, err := db.Exec(sqlStatement, ID)
	rowsAffected, _ := result.RowsAffected()

	return rowsAffected == 1, err
}

//IsFileAlreadyRead check if the filename has already been read
func (f *FileService) IsFileAlreadyRead(filename string) (bool, error) {
	db := util.OpenConnection(f.Props)
	defer util.CloseConnection(db)

	sqlStatement := `select
		count(*)
	from
		file
	where
		name = $1;`

	row := db.QueryRow(sqlStatement, filename)

	var count int
	err := row.Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

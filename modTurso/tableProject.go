package modTurso

import (
	"fmt"
	"time"
)

type cDataProject struct {
	ID        int
	Name      string
	Datetime1 time.Time
}

type CTableProject struct {
}

const sql_create_testproject = `CREATE TABLE if not exists testproject (
	ID INTEGER PRIMARY KEY,
	Name CHAR(50) NOT NULL,
	Date1 datetime default current_timestamp
 );`

var g_singleProject *CTableProject = &CTableProject{}

func GetSingleProject() *CTableProject {
	return g_singleProject
}

func (pInst *CTableProject) Initialize() error {
	err := pInst.createTable()

	return err
}

func (pInst *CTableProject) createTable() error {
	err := GetSingleTurso().execsql(sql_create_testproject)
	if err != nil {
		return err
	}

	return nil
}

func (pInst *CTableProject) Insert(name string) error {
	sql := fmt.Sprintf("INSERT INTO testproject(Name) VALUES('%s');", name)
	err := GetSingleTurso().execsql(sql)

	return err
}
func (pInst *CTableProject) Query(projectname string) error {
	sql := "SELECT * FROM testproject"
	if projectname != "" {
		sql = sql + " WHERE Name like '%" + projectname + "%'"
	}
	rows, err := GetSingleTurso().query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var project cDataProject

		if err := rows.Scan(&project.ID, &project.Name, &project.Datetime1); err != nil {
			fmt.Println("Error scanning row:", err)
			return err
		}

		fmt.Println(project.ID, project.Name, project.Datetime1)
	}

	return nil
}

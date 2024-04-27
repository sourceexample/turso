package modTurso

import "database/sql"

type CDBTurso struct {
	sqlOpt *CSqlDbOpt
}

var g_singleTruso *CDBTurso = &CDBTurso{sqlOpt: newSqlDbOpt()}

func GetSingleTurso() *CDBTurso {
	return g_singleTruso
}

func (pInst *CDBTurso) Connect(dbUrl, token string) error {
	return pInst.sqlOpt.connect(dbUrl, token)
}

func (pInst *CDBTurso) execsql(sql string) error {
	return pInst.sqlOpt.execsql(sql)
}

func (pInst *CDBTurso) query(sql string) (*sql.Rows, error) {
	return pInst.sqlOpt.query(sql)

}

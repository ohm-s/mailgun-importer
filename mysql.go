package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type MysqlDB struct {
	Instance *sql.DB
}

func GetMysqlConnection(hostName string, portNumber string, databaseName string, userName string, passWord string, poolSize int, idleSize int) (MysqlDB, error) {
	dataSourceName := strings.Join([]string {userName, ":", passWord, "@tcp(", hostName, ":", portNumber, ")/", databaseName}, "")
	db, err := sql.Open("mysql", dataSourceName)
	db.SetMaxOpenConns(poolSize)
	db.SetMaxIdleConns(idleSize)
	mysqlDatabase := MysqlDB{Instance:db}
	return mysqlDatabase, err
}

func (db *MysqlDB)  FetchColumn(sqlQuery string, args ...interface{}) ([]string, error) {
	rows, errSql := db.Instance.Query(sqlQuery, args...)
	if errSql != nil {
		result := make([]string, 0)
		return result, errSql
	}
	defer rows.Close()
	finalResult := make([]string, 0)
	for rows.Next() {
		var dbValue string
		rows.Scan(&dbValue)
		finalResult = append(finalResult, dbValue)
	}
	return finalResult, nil
}


func (db *MysqlDB)  ExecuteNonQuery(sqlQuery string, args ...interface{}) (int64, error) {
	result, errSql := db.Instance.Exec(sqlQuery, args...)
	if errSql != nil {
		return 0, errSql
	}
	rowsNumber, err := result.RowsAffected()
	return   rowsNumber, err
}

func (db *MysqlDB)  ExecuteScalar(sqlQuery string, args ...interface{}) (string, error) {
	rows, errSql := db.Instance.Query(sqlQuery, args...)
	if errSql != nil {
		result := ""
		return result, errSql
	}
	defer rows.Close()
	finalResult := ""
	if rows.Next() {
		var dbValue string
		rows.Scan(&dbValue)
		finalResult = dbValue
	}
	return finalResult, nil
}

func (db *MysqlDB)  ExecuteTupleScalar(sqlQuery string, args ...interface{}) (string, string, error) {
	rows, errSql := db.Instance.Query(sqlQuery, args...)
	if errSql != nil {
		result := ""
		return result, result, errSql
	}
	defer rows.Close()
	finalResult := ""
	finalResult2 := ""
	if rows.Next() {
		var dbValue string
		var dbValue2 string
		rows.Scan(&dbValue, &dbValue2)
		finalResult = dbValue
		finalResult2 = dbValue2
	}
	return finalResult, finalResult2, nil
}

func (db *MysqlDB)  HashMapQuery(sqlQuery string, args ...interface{}) (map[string]string, error) {
	rows, errSql := db.Instance.Query(sqlQuery, args...)
	if errSql != nil {
		result := make(map[string]string, 0)
		return result, errSql
	}
	defer rows.Close()
	finalResult := make(map[string]string, 0)
	for rows.Next() {
		var keyVar string
		var valueVar string
		rows.Scan(&keyVar, &valueVar)
		finalResult[keyVar] = valueVar
	}
	return finalResult, nil
}

func (db *MysqlDB) Insert(insertQuery string, args ...interface{}) error {
	_, err := db.Instance.Exec(insertQuery, args...)
	return err
}

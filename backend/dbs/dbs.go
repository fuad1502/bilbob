package dbs

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"sync"
)

// SafeDB is a wrapper around sql.DB that provides a mutex to make it safe for concurrent use
type SafeDB struct {
	Lock  sync.Mutex
	DB    *sql.DB
	stmts map[string]*sql.Stmt
}

// ConnectDB connects to a Postgres database and checks if the connection is working and returns a SafeDB struct pointer if successful
func ConnectPGDB(host string, user string, password string, dbname string) (*SafeDB, error) {
	connStr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", host, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.New("Failed to connect to the database: " + err.Error())
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, errors.New("Failed to ping the database: " + err.Error())
	}
	return &SafeDB{DB: db}, nil
}

func (safeDB *SafeDB) getStmt(query string) (*sql.Stmt, error) {
	stmt, ok := safeDB.stmts[query]
	if ok {
		return stmt, nil
	}
	return safeDB.DB.Prepare(query)
}

func (safeDB *SafeDB) QueryRow(query string, row any, args ...any) error {
	stmt, err := safeDB.getStmt(query)
	if err != nil {
		panic(err)
	}

	sqlRow := stmt.QueryRow(args...)

	r := reflect.ValueOf(row)
	if r.Kind() != reflect.Pointer {
		return errors.New("dbs.QueryRow: Parameter 'row' must be of type pointer")
	}

	r = reflect.Indirect(r)
	if r.Kind() == reflect.Struct {
		param := make([]any, r.NumField())
		for i := range param {
			param[i] = r.Field(i).Addr().Interface()
		}
		return sqlRow.Scan(param...)
	} else {
		return sqlRow.Scan(row)
	}
}

func (safeDB *SafeDB) Query(query string, rows any, args ...any) (any, error) {
	t := reflect.TypeOf(rows)
	if t.Kind() != reflect.Slice {
		return rows, fmt.Errorf("dbs.QueryRow: rows should be of type slice")
	}
	rowsRefl := reflect.ValueOf(rows)

	stmt, err := safeDB.getStmt(query)
	if err != nil {
		panic(fmt.Errorf("dbs.QueryRow: %v", err))
	}

	sqlRows, err := stmt.Query(args...)
	if err != nil {
		return rows, fmt.Errorf("dbs.QueryRow: %v", err)
	}

	t = t.Elem()
	if t.Kind() == reflect.Struct {
		row := reflect.Indirect(reflect.New(t))
		param := make([]any, row.NumField())
		for i := range param {
			param[i] = row.Field(i).Addr().Interface()
		}
		count := 0
		for sqlRows.Next() {
			sqlRows.Scan(param...)
			if count >= rowsRefl.Len() {
				rowsRefl = reflect.Append(rowsRefl, row)
			} else {
				rowsRefl.Index(count).Set(row)
			}
			count += 1
		}
	} else {
		row := reflect.Indirect(reflect.New(t))
		count := 0
		for sqlRows.Next() {
			sqlRows.Scan(row.Addr().Interface())
			if count >= rowsRefl.Len() {
				rowsRefl = reflect.Append(rowsRefl, row)
			} else {
				rowsRefl.Index(count).Set(row)
			}
			count += 1
		}
	}
	return rowsRefl.Interface(), nil
}

func (safeDB *SafeDB) InsertRow(insertStmt string, row any) error {
	stmt, err := safeDB.getStmt(insertStmt)
	if err != nil {
		panic(err)
	}

	r := reflect.ValueOf(row)
	if r.Kind() != reflect.Pointer {
		return errors.New("dbs.Insert: Parameter 'row' must be of type pointer")
	}

	r = reflect.Indirect(r)
	if r.Kind() == reflect.Struct {
		param := make([]any, r.NumField())
		for i := range param {
			param[i] = r.Field(i).Addr().Interface()
		}
		_, err := stmt.Exec(param...)
		return err
	} else {
		_, err := stmt.Exec(row)
		return err
	}
}

func (safeDB *SafeDB) DeleteRow(deleteStmt string, args ...any) error {
	stmt, err := safeDB.getStmt(deleteStmt)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(args...)
	return err
}

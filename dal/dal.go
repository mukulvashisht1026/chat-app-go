package dal

import "database/sql"

type GenericDatabaseLayer interface {
	Connect() error
	Disconnect() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Insert(table string, values ...interface{}) (int64, error)
}

type MysqlDatabaseLayer struct {
	Db *sql.DB
}

func NewMysqlDatabaseLayer(dataSourceName string) (*MysqlDatabaseLayer, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &MysqlDatabaseLayer{Db: db}, nil
}

func (m *MysqlDatabaseLayer) Connect() error {
	return m.Db.Ping()
}

func (m *MysqlDatabaseLayer) Disconnect() error {
	return m.Db.Close()
}

func (m *MysqlDatabaseLayer) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return m.Db.Query(query, args...)
}

func (m *MysqlDatabaseLayer) Insert(table string, values ...interface{}) (int64, error) {
	stmt, err := m.Db.Prepare("INSERT INTO " + table + " VALUES (?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

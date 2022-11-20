package db

import (
	"database/sql"
	"testing"
)

type MockDB struct {
	query        string
	lastInsertId int64
	rowsAffected int64
}

func (md MockDB) LastInsertId() (int64, error) {
	return md.lastInsertId, nil
}

func (md MockDB) RowsAffected() (int64, error) {
	return md.rowsAffected, nil
}

func (md *MockDB) Exec(query string, args ...any) (sql.Result, error) {
	md.query = query
	return md, nil
}

func TestExecWithMock(t *testing.T) {
	mock := &MockDB{
		rowsAffected: 32,
	}

	got, _ := execQuery(mock, "SELECT * FROM table")

	if mock.query != "SELECT * FROM table" {
		t.Errorf("Expected 'SELECT * FROM table', got %s", mock.query)
	}

	if got != 32 {
		t.Errorf("Expected 32, got %d", got)
	}

}

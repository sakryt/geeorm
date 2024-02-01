package session

import (
	"database/sql"
	"geeorm/dialect"
	"os"
	"testing"
)

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("sqlite3")
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func NewSession() *Session {
	return New(TestDB, TestDial)
}

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite3", "../gee.db")
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

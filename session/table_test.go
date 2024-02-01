package session

import (
	"geeorm/dialect"
	"testing"
)

func TestSession_CreateTable(t *testing.T) {
	dialect.Init()
	s := NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}

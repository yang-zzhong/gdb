package mysql

import (
	. "gdb/query"
	. "testing"
)

func TestForQuery(t *T) {

	builder := New()
	builder.From("users")
	if "SELECT * FROM users" != builder.ForQuery() {
		t.Error("From Error")
	}
	builder.Select([]string{"name", "age"})
	if "SELECT name, age FROM users" != builder.ForQuery() {
		t.Error("Select Error")
	}

	builder = New()
	builder.From("users")
	builder.Where("name", "yang-zzhong")
	if "SELECT * FROM users WHERE name = ?" != builder.ForQuery() {
		t.Error("Where Error")
	}
	builder.Where("age", GT, "15")
	form := "SELECT * FROM users WHERE name = ? AND age > ?"
	if form != builder.ForQuery() {
		t.Error("Where And Error")
	}

	builder = New()
	builder.From("users")
	builder.Where("name", "yang-zzhong")
	builder.Quote(func(builder Builder) {
		builder.Where("age", GT, "50")
		builder.Or().Where("age", LT, "10")
	})
	form = "SELECT * FROM users WHERE name = ? AND (age > ? OR age < ?)"
	if form != builder.ForQuery() {
		t.Error("QWhere Error")
	}
}

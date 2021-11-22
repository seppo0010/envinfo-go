package envinfo

import (
	"fmt"
	"strings"
)

func GetDatabases() []*Item {
	return getItems([]func() *Item{
		GetMongoDB,
		GetPostgres,
		GetMysql,
		GetSqlite,
	})
}

func GetMongoDB() *Item {
	return GetItem("mongo", "MongoDB")
}

func GetPostgres() *Item {
	return GetItem("postgres", "PostgresQL")
}

func GetMysql() *Item {
	b := NewGetItemBuilder("mysql", "MySQL")
	return b.ParseVersion(func(unparsed string) string {
		matches := b.regex.FindAllString(unparsed, 2)
		v := matches[0]
		if strings.Contains(unparsed, "Distrib") && len(matches) >= 2 {
			v = matches[1]
		}
		if strings.Contains(unparsed, "MariaDB") {
			return fmt.Sprintf("%s (MariaDB)", v)
		}
		return v
	}).Get()
}

func GetSqlite() *Item {
	return GetItem("sqlite3", "SQLite")
}

package utils

import (
	"database/sql"
	"time"
)

func SqlString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func SqlInt64(n int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: n,
		Valid: true,
	}
}

func SqlTime() sql.NullTime {
	return SqlTimeX("")
}

func SqlTimeX(t string) sql.NullTime {
	if t == "" {
		t = time.Now().Format("2006-01-01 01:12")
	}
	return sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
}

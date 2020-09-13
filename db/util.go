package db

import "database/sql"

func StringPtrToSqlNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{
			String: *s,
			Valid:  true,
		}
	}
	return sql.NullString{
		Valid: false,
	}
}

func StringToSqlNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func StringToStringPtr(s string) *string {
	return &s
}

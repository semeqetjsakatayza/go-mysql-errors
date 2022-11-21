package mysqlerrors

import (
	mysqldriver "github.com/go-sql-driver/mysql"
)

// checkMySQLErrorNumber test if given `err` is a MySQLError with given error number `num`.
func checkMySQLErrorNumber(err error, num uint16) bool {
	if mysqlErr, ok := err.(*mysqldriver.MySQLError); ok && (mysqlErr.Number == num) {
		return true
	}
	return false
}

// IsTableNotExistError check if given `err` is MySQL table not exist error.
func IsTableNotExistError(err error) bool {
	return checkMySQLErrorNumber(err, 1146)
}

// IsDuplicatedEntryError check if given `err` is MySQL duplicated entry error.
func IsDuplicatedEntryError(err error) bool {
	return checkMySQLErrorNumber(err, 1062)
}

// IsUnknownColumnError check if given `err` is MySQL unknown column error.
func IsUnknownColumnError(err error) bool {
	return checkMySQLErrorNumber(err, 1054)
}

// IsUnknownTableError check if given `err` is MySQL unknown table error.
func IsUnknownTableError(err error) bool {
	return checkMySQLErrorNumber(err, 1051)
}

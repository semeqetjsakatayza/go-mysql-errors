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

// IsDeadlockFoundError check if given `err` is MySQL deadlock found error.
func IsDeadlockFoundError(err error) bool {
	return checkMySQLErrorNumber(err, 1213)
}

// IsLockWaitTimeoutError check if given `err` is MySQL lock wait timeout error.
func IsLockWaitTimeoutError(err error) bool {
	return checkMySQLErrorNumber(err, 1205)
}

// IsTableNotExistError check if given `err` is MySQL table not exist error.
func IsTableNotExistError(err error) bool {
	return checkMySQLErrorNumber(err, 1146)
}

// IsDuplicatedColumnError check if given `err` is MySQL duplicated column name error.
func IsDuplicatedColumnError(err error) bool {
	return checkMySQLErrorNumber(err, 1060)
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

// IsServerGoneAwayError check if given `err` is MySQL server has gone away error.
func IsServerGoneAwayError(err error) bool {
	return checkMySQLErrorNumber(err, 2006)
}

// IsServerLostConnectionError check if given `err` is lost connection to MySQL server during query error.
func IsServerLostConnectionError(err error) bool {
	return checkMySQLErrorNumber(err, 2013)
}

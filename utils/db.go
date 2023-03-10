package utils

import (
	"github.com/go-sql-driver/mysql"
)

const (
	ErrMySQLDupEntry            = 1062
	ErrMySQLDupEntryWithKeyName = 1586
)

// 唯一索引重复
func IsUniqueConstraintError(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr.Number == ErrMySQLDupEntry ||
			mysqlErr.Number == ErrMySQLDupEntryWithKeyName {
			return true
		}
	}
	return false
}

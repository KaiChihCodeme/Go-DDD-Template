package serviceProvidersDb

import "database/sql"

type IDbConnector interface {
	GetConnection() (*sql.DB, error)
	GetConnectionString() (*string, error)
}

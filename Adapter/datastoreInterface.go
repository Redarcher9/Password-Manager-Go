package Adapter

import (
	"database/sql"
	"password-manager/datastore/inmem"
)

type DatastoreInstance struct{
	DbInstance *sql.DB | Inmem
}
func DatastoreAdapter() DatastoreInstance {
	return 
}
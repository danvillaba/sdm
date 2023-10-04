package inspectorsvc

import (
	"fmt"
	pginspector "sdm/module/inspector/service/postgres"
	inspectorv1 "sdm/pb/sdm/inspector/v1"

	"github.com/jmoiron/sqlx"
)

func CreateInspectorService(
	db *sqlx.DB,
	driver string,
) (inspectorv1.InspectorServiceServer, error) {

	switch driver {
	case "postgres", "pg", "postgresql":
		return pginspector.CreatePgInspector(db), nil

	case "mssql", "sqlserver":
		return nil, fmt.Errorf("driver not implemented")

	case "mysql", "mariadb":
		return nil, fmt.Errorf("driver not implemented")

	default:
		return nil, fmt.Errorf("driver %s not supported", driver)
	}
}

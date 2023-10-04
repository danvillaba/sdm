package pginspector

import (
	inspectorv1 "sdm/pb/sdm/inspector/v1"

	"github.com/jmoiron/sqlx"
)

type PgInspector struct {
	db *sqlx.DB
	inspectorv1.UnimplementedInspectorServiceServer
}

func CreatePgInspector(db *sqlx.DB) inspectorv1.InspectorServiceServer {
	return &PgInspector{db: db}
}

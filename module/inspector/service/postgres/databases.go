package pginspector

import (
	"context"
	_ "embed"
	inspectorv1 "sdm/pb/sdm/inspector/v1"
)

func (s *PgInspector) ListDatabases(
	ctx context.Context,
	req *inspectorv1.ListDatabasesRequest,
) (*inspectorv1.ListDatabasesResponse, error) {

	dbList := []*inspectorv1.Database{}

	err := s.db.SelectContext(ctx, &dbList, databaseListQuery)
	if err != nil {
		return nil, err
	}

	return &inspectorv1.ListDatabasesResponse{
		Databases: dbList,
	}, nil
}

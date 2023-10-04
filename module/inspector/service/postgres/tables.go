package pginspector

import (
	"context"
	inspectorv1 "sdm/pb/sdm/inspector/v1"
)

func (s *PgInspector) ListTables(
	ctx context.Context,
	req *inspectorv1.ListTablesRequest,
) (*inspectorv1.ListTablesResponse, error) {

	tables := []*inspectorv1.Table{}

	err := s.db.SelectContext(ctx, &tables, tableListQuery, req.Schema)
	if err != nil {
		return nil, err
	}

	return &inspectorv1.ListTablesResponse{
		Tables: tables,
	}, nil
}

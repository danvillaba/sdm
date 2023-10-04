package pginspector

import (
	"context"
	inspectorv1 "sdm/pb/sdm/inspector/v1"
)

func (s *PgInspector) GetTableInfo(
	ctx context.Context,
	req *inspectorv1.GetTableInfoRequest,
) (*inspectorv1.GetTableInfoResponse, error) {
	db := s.db.Unsafe()

	table := &inspectorv1.Table{}
	err := db.GetContext(ctx, table, tableInfoQuery, req.Schema, req.Table)
	if err != nil {
		return nil, err
	}

	columns := []*inspectorv1.Column{}
	err = db.SelectContext(ctx, &columns, tableColumnsQuery, req.Table, req.Schema)
	if err != nil {
		return nil, err
	}

	return &inspectorv1.GetTableInfoResponse{
		Table:   table,
		Columns: columns,
	}, nil
}

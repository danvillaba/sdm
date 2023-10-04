package pginspector

import (
	"context"
	inspectorv1 "sdm/pb/sdm/inspector/v1"
)

func (s *PgInspector) Query(
	ctx context.Context,
	req *inspectorv1.QueryRequest,
) (*inspectorv1.QueryResponse, error) {

	rows, err := s.db.QueryxContext(ctx, req.Sql)
	if err != nil {
		return nil, err
	}

	columns, err := getRowsColumns(rows)
	if err != nil {
		return nil, err
	}

	listValues := parseValues(rows, len(columns))

	return &inspectorv1.QueryResponse{
		Columns: columns,
		Rows:    listValues,
	}, nil
}

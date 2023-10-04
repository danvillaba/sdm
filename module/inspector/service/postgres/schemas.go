package pginspector

import (
	"context"
	_ "embed"
	inspectorv1 "sdm/pb/sdm/inspector/v1"
)

func (s *PgInspector) ListSchemas(
	ctx context.Context,
	req *inspectorv1.ListSchemasRequest,
) (*inspectorv1.ListSchemasResponse, error) {

	schemas := []*inspectorv1.Schema{}

	err := s.db.SelectContext(ctx, &schemas, schemaListQuery)
	if err != nil {
		return nil, err
	}

	return &inspectorv1.ListSchemasResponse{
		Schemas: schemas,
	}, nil
}

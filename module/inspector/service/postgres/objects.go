package pginspector

import (
	"context"
	inspectorv1 "sdm/pb/sdm/inspector/v1"
)

func (s *PgInspector) ListObjects(
	ctx context.Context,
	req *inspectorv1.ListObjectsRequest,
) (*inspectorv1.ListObjectsResponse, error) {

	objectList := []*inspectorv1.Object{}

	err := s.db.SelectContext(ctx, &objectList, objectQuery)
	if err != nil {
		return nil, err
	}

	return &inspectorv1.ListObjectsResponse{
		Objects: objectList,
	}, nil
}

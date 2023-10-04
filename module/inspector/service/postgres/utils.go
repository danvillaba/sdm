package pginspector

import (
	"log"
	inspectorv1 "sdm/pb/sdm/inspector/v1"

	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/structpb"
)

func getRowsColumns(rows *sqlx.Rows) ([]*inspectorv1.Column, error) {
	rowTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	columns := []*inspectorv1.Column{}

	for _, coltype := range rowTypes {
		nullable, _ := coltype.Nullable()

		log.Print(coltype)
		columns = append(columns, &inspectorv1.Column{
			Name:       coltype.Name(),
			IsNullable: nullable,
		})
	}

	return columns, nil
}

func parseValues(rows *sqlx.Rows, rowlength int) []*structpb.ListValue {

	listValues := []*structpb.ListValue{}

	for rows.Next() {
		scanslice := make([]any, rowlength)
		for i := 0; i < rowlength; i++ {
			scanslice[i] = new([]byte)
		}

		rows.Scan(scanslice...)

		var parsedValues []any
		for _, value := range scanslice {
			v := value.(*[]byte)

			parsedValues = append(parsedValues, string(*v))
		}

		values, _ := structpb.NewList(parsedValues)

		listValues = append(listValues, values)
	}

	return listValues
}

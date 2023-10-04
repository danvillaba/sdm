package pginspector

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed sql/objects.sql
	objectFragment string

	//go:embed sql/table.sql
	tableFragment string

	//go:embed sql/column.sql
	columnFragment string

	//go:embed sql/databases.sql
	databaseListQuery string

	//go:embed sql/schemas.sql
	schemaListQuery string

	tableInfoQuery = fmt.Sprintf(`%s
select *
from tbl
where schema = $1
  and name = $2`, tableFragment)

	tableListQuery = fmt.Sprintf(`%s
select *
from tbl
where schema = $1`, tableFragment)

	tableColumnsQuery = fmt.Sprintf(`%s
select *
from col
where tablename = $1
  and schemaname = $2
`, columnFragment)

	objectQuery = fmt.Sprintf(`%s
select *
from obj
order by 1, 2`, objectFragment)
)

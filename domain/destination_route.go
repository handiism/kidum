package domain

import "database/sql"

type DestinationRoute struct {
	Id            int
	Order         sql.NullInt64
	RouteId       sql.NullInt64
	DestinationId sql.NullInt64
}

package models

import (
	"context"

	"github.com/lib/pq"
)

const query = `
	SELECT client_id, secret_key, grant_type, is_active, array_agg(e.enum_name) scopes
	FROM api_clients ac 
	LEFT JOIN enums e ON e.id = ANY(ARRAY[ac.scopes])
	GROUP BY client_id, secret_key, grant_type, is_active
`
type Result struct {
	ClientId string
	SecretKey string
	GrantType string
	IsActive bool
	Scopes []string
}

func (q *Queries) GetData(ctx context.Context) (Result, error) {
	row := q.db.QueryRowContext(ctx, query)

	var i Result
	err := row.Scan(
		&i.ClientId,
		&i.SecretKey,
		&i.GrantType,
		&i.IsActive,
		pq.Array(&i.Scopes),
	)

	return i, err
}
package models

import (
	"context"

	"github.com/lib/pq"
)

const queryApiClient = `
	SELECT client_id, secret_key, grant_type, is_active, uuid, array_agg(e.enum_name) scopes
	FROM api_clients ac 
	LEFT JOIN enums e ON e.id = ANY(ARRAY[ac.scopes])
	WHERE ac.client_id = $1 AND ac.secret_key = $2
	GROUP BY client_id, secret_key, grant_type, is_active, uuid
`

type ParamApiClient struct {
	ClientId string
	SecretKey string
}

type DataApiClient struct {
	ClientId string
	SecretKey string
	GrantType string
	IsActive bool
	Uuid string
	Scopes []string
}

func (q *Queries) GetDataApiClient(ctx context.Context, val ParamApiClient) (DataApiClient, error) {
	row := q.db.QueryRowContext(ctx, queryApiClient,
		val.ClientId,
		val.SecretKey,
	)

	var i DataApiClient
	err := row.Scan(
		&i.ClientId,
		&i.SecretKey,
		&i.GrantType,
		&i.IsActive,
		&i.Uuid,
		pq.Array(&i.Scopes),
	)

	return i, err
}
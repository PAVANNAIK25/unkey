// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: ratelimit_namespace_find_by_id.sql

package gen

import (
	"context"
)

const findRatelimitNamespaceByID = `-- name: FindRatelimitNamespaceByID :one
SELECT id, workspace_id, name, created_at, updated_at, deleted_at FROM ` + "`" + `ratelimit_namespaces` + "`" + `
WHERE id = ?
`

// FindRatelimitNamespaceByID
//
//	SELECT id, workspace_id, name, created_at, updated_at, deleted_at FROM `ratelimit_namespaces`
//	WHERE id = ?
func (q *Queries) FindRatelimitNamespaceByID(ctx context.Context, id string) (RatelimitNamespace, error) {
	row := q.db.QueryRowContext(ctx, findRatelimitNamespaceByID, id)
	var i RatelimitNamespace
	err := row.Scan(
		&i.ID,
		&i.WorkspaceID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

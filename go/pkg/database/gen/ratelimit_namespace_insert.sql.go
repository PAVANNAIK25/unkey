// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: ratelimit_namespace_insert.sql

package gen

import (
	"context"
)

const insertRatelimitNamespace = `-- name: InsertRatelimitNamespace :exec
INSERT INTO ` + "`" + `ratelimit_namespaces` + "`" + ` (
    id,
    workspace_id,
    name,
    created_at
)
VALUES (
    ?,
    ?,
    ?,
    NOW()
)
`

type InsertRatelimitNamespaceParams struct {
	ID          string `db:"id"`
	WorkspaceID string `db:"workspace_id"`
	Name        string `db:"name"`
}

// InsertRatelimitNamespace
//
//	INSERT INTO `ratelimit_namespaces` (
//	    id,
//	    workspace_id,
//	    name,
//	    created_at
//	)
//	VALUES (
//	    ?,
//	    ?,
//	    ?,
//	    NOW()
//	)
func (q *Queries) InsertRatelimitNamespace(ctx context.Context, arg InsertRatelimitNamespaceParams) error {
	_, err := q.db.ExecContext(ctx, insertRatelimitNamespace, arg.ID, arg.WorkspaceID, arg.Name)
	return err
}

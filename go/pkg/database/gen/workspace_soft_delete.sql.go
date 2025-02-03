// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: workspace_soft_delete.sql

package gen

import (
	"context"
	"database/sql"
)

const softDeleteWorkspace = `-- name: SoftDeleteWorkspace :execresult
UPDATE ` + "`" + `workspaces` + "`" + `
SET deleted_at = NOW()
WHERE id = ?
AND delete_protection = false
`

// SoftDeleteWorkspace
//
//	UPDATE `workspaces`
//	SET deleted_at = NOW()
//	WHERE id = ?
//	AND delete_protection = false
func (q *Queries) SoftDeleteWorkspace(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, softDeleteWorkspace, id)
}

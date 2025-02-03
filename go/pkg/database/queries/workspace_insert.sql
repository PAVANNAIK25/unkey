-- name: InsertWorkspace :exec
INSERT INTO `workspaces` (
    id,
    tenant_id,
    name,
    created_at,
    plan,
    beta_features,
    features,
    enabled,
    delete_protection
)
VALUES (
    sqlc.arg(id),
    sqlc.arg(tenant_id),
    sqlc.arg(name),
    NOW(),
    'free',
    '{}',
    '{}',
    true,
    true
);

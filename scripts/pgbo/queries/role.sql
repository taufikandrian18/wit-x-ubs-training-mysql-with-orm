-- name: InsertRole :exec
INSERT INTO role (role_name, status, created_at, created_by)
VALUES (?, ?, CURRENT_TIMESTAMP, ?);

-- name: GetLatsInsertedID :one
SELECT LAST_INSERT_ID();

-- name: UpdateRole :exec
UPDATE role
SET role_name      = ?,
    status         = ?,
    updated_at     = CURRENT_TIMESTAMP,
    updated_by     = ?
WHERE role_id = ?;

-- name: GetRole :one
SELECT r.role_id,
       r.role_name,
       r.status,
       r.created_at,
       r.created_by,
       r.updated_at,
       r.updated_by
FROM role r
WHERE r.role_id = ?
  AND r.status != 'deleted';

-- name: ListRoles :many
SELECT r.role_id,
       r.role_name,
       r.status,
       r.created_at,
       r.created_by,
       r.updated_at,
       r.updated_by
FROM role r
WHERE (? = TRUE AND r.role_id = ? OR ? = FALSE)
  AND (? = TRUE AND r.role_name = ? OR ? = FALSE)
  AND r.status != 'deleted'
ORDER BY 
    CASE WHEN ? = 'role_id ASC' THEN r.role_id END ASC,
    CASE WHEN ? = 'role_id DESC' THEN r.role_id END DESC,
    CASE WHEN ? = 'created_at ASC' THEN r.created_at END ASC,
    CASE WHEN ? = 'created_at DESC' THEN r.created_at END DESC,
    r.created_at DESC
LIMIT ? OFFSET ?;

-- name: CountRole :one
SELECT COUNT(r.role_id)
FROM role r
WHERE (? = TRUE AND r.role_id = ? OR ? = FALSE)
  AND (? = TRUE AND r.role_name = ? OR ? = FALSE)
  AND r.status != 'deleted';


-- name: UpdateRoleStatus :exec
UPDATE role
SET status     = ?,
    updated_at = CURRENT_TIMESTAMP,
    updated_by = ?
WHERE role_id = ?
  AND status != 'deleted';

-- Registry
-- name: CreatePlatformRegistry :one
INSERT INTO platform_registry (
  id, project_name, repository, username, email, created_at, deleted_at
) VALUES ($1, $2, $3, $4, $5, $6, $7)RETURNING *;

-- name: GetAllPlatformRegistry :many
SELECT id, project_name, repository, username, email, created_at, deleted_at
FROM platform_registry
ORDER BY deleted_at DESC, email ASC;

-- name: GetByFilterPlatformRegistry :many
SELECT id, project_name, repository, username, email, created_at, deleted_at
FROM platform_registry
WHERE
 (@is_project_name::bool = FALSE OR project_name LIKE '%' || COALESCE(NULLIF(@project_name::text, '') || '%', project_name))
  AND (@is_repository::bool = FALSE OR repository LIKE '%' || COALESCE(NULLIF(@repository::text, '') || '%', repository))
  AND (@is_username::bool = FALSE OR username LIKE '%' || COALESCE(NULLIF(@username::text, '') || '%', username))
  AND (@is_email::bool = FALSE OR email LIKE '%' || COALESCE(NULLIF(@email::text, '') || '%', email))
ORDER BY deleted_at DESC, email ASC;

-- name: GetByIDPlatformRegistry :one
SELECT id, project_name, repository, username, email, created_at, deleted_at
FROM platform_registry
WHERE id=$1
ORDER BY email ASC;

-- LabDestroy
-- name: CreatePlatformLabDestroy :one
INSERT INTO platform_lab_destroy (
  id, project_name, repository, username, email, created_at, target_destroy, available
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)RETURNING *;

-- name: GetAllPlatformLabDestroy :many
SELECT id, project_name, repository, username, email, created_at, target_destroy, available, error_message, updated_at
FROM platform_lab_destroy
ORDER BY target_destroy DESC, email ASC;

-- name: GetByFilterPlatformLabDestroy :many
SELECT id, project_name, repository, username, email, created_at, target_destroy, available, error_message, updated_at
FROM platform_lab_destroy
WHERE
 (@is_project_name::bool = FALSE OR project_name LIKE '%' || COALESCE(NULLIF(@project_name::text, '') || '%', project_name))
  AND (@is_repository::bool = FALSE OR repository LIKE '%' || COALESCE(NULLIF(@repository::text, '') || '%', repository))
  AND (@is_username::bool = FALSE OR username LIKE '%' || COALESCE(NULLIF(@username::text, '') || '%', username))
  AND (@is_email::bool = FALSE OR email LIKE '%' || COALESCE(NULLIF(@email::text, '') || '%', email))
  AND (@is_available::bool = FALSE OR available = @available)
ORDER BY target_destroy DESC, email ASC;

-- name: GetByIDPlatformLabDestroy :one
SELECT id, project_name, repository, username, email, created_at, target_destroy, available, error_message, updated_at
FROM platform_lab_destroy
WHERE id=$1
ORDER BY email ASC;

-- name: PatchlatformLabDestroy :one
UPDATE platform_lab_destroy
SET 
  available=$2, 
  error_message = COALESCE(NULLIF(@error_message::text, ''), error_message),
  updated_at=$3
WHERE id=$1 RETURNING *;

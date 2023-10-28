-- name: CreateActivity :one
INSERT INTO activities (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: GetActivity :one
SELECT * FROM activities
WHERE id = $1 LIMIT 1;

-- name: ListActivities :many
SELECT * FROM activities
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateActivity :one
UPDATE activities
  set name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteActivity :exec
DELETE FROM activities
WHERE id = $1;

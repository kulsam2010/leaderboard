-- name: CreateScoreAudit :one
INSERT INTO score_audit (
  user_id, activity_id, points
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetScoreAuditForUserAndActivity :many
SELECT * FROM score_audit
WHERE user_id = $1 and activity_id=$2 LIMIT 1;

-- name: GetScoreAudit :one
SELECT * FROM score_audit
WHERE id = $1 LIMIT 1;

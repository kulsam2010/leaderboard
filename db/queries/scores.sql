-- name: CreateScore :one
INSERT INTO scores (
  user_id, activity_id, score
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetScoreForUserAndActivity :one
SELECT * FROM scores
WHERE user_id = $1 and activity_id=$2 LIMIT 1;

-- name: ListScores :many
SELECT * FROM scores
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateScore :one
UPDATE scores
  set score = $3
WHERE user_id = $1 and activity_id=$2
RETURNING *;

-- name: DeleteScore :exec
DELETE FROM scores
WHERE id = $1;

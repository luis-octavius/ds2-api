-- name: CreateMiracle :one 
INSERT INTO miracles (id, name, uses, attunement, description, acquired, cost, miracle_type)
VALUES (
  gen_random_uuid(),
  $1, 
  $2, 
  $3, 
  $4, 
  $5, 
  $6, 
  $7
) 
RETURNING *;

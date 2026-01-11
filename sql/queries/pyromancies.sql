-- name: CreatePyromancy :one
INSERT INTO pyromancies(id, name, uses, attunement, description, acquired, cost, pyromancy_type)
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

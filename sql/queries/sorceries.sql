-- name: CreateSorcery :one 
INSERT INTO sorceries(id, name, uses, spellslot, intelligence, description, acquired, cost)
VALUES(
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

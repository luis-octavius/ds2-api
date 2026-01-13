-- name: CreateSorcery :one 
INSERT INTO sorceries(id, name, uses, spellslot, intelligence, description, acquired, cost, sorcery_type)
VALUES(
  gen_random_uuid(),
  $1, 
  $2, 
  $3, 
  $4,
  $5, 
  $6, 
  $7,
  $8
)
RETURNING *;

-- name: GetSorceries :many 
SELECT * FROM sorceries;

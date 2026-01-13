-- name: CreateHex :one
INSERT INTO hexes(id, name, uses, attunement, intelligence, faith, description, acquired, cost, hex_type)
VALUES (
  gen_random_uuid(),
  $1, 
  $2, 
  $3, 
  $4, 
  $5, 
  $6, 
  $7, 
  $8, 
  $9
) 
RETURNING *;

-- name: GetHexes :many 
SELECT * FROM hexes;

-- name: CreateHorseGender :one
INSERT INTO horse_gender (name, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetHorseGenderList :many
SELECT *
FROM horse_gender
ORDER BY name;

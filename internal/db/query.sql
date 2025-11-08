-- name: CreateHorseGender :one
INSERT INTO horse_gender (name, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetHorseGenderList :many
SELECT *
FROM horse_gender
ORDER BY name;

-- name: GetHorseColorList :many
SELECT *
FROM horse_color
ORDER BY name;

-- name: GetHorseBirthplaceList :many
SELECT *
FROM horse_birthplace
ORDER BY name;

-- name: GetHorseGeneticMarkerList :many
SELECT *
FROM genetic_marker
ORDER BY name;
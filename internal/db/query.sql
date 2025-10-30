-- name: CreateHorseGender :exec
INSERT INTO horse_gender (name, description)
VALUES ($1, $2);

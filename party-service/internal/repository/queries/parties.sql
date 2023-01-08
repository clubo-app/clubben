-- name: DeleteParty :exec
DELETE FROM parties
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: IncreaseFavoriteCount :exec
UPDATE parties
SET favorite_count = favorite_count + ?
WHERE id = ?;

-- name: DecreaseFavoriteCount :exec
UPDATE parties
SET favorite_count = favorite_count - ?
WHERE id = ?;


-- name: IncreaseParticipantsCount :exec
UPDATE parties
SET participants_count = participants_count + ?
WHERE id = ?;

-- name: DecreaseParticipantsCount :exec
UPDATE parties
SET participants_count = participants_count - ?
WHERE id = ?;

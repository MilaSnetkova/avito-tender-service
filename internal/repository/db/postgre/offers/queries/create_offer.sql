INSERT INTO offers (description, status, version, tender_id, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, description, status, version, created_at, updated_at, tender_id, user_id; 
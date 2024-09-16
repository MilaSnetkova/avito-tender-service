INSERT INTO tenders (title, description, status, version, organization_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, title, description, status, version, created_at, updated_at, organization_id
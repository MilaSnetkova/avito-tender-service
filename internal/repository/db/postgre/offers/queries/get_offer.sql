INSERT INTO offers (tender_id, title, description, status, version, created_by, organization_id) 
VALUES ($1, $2, $3, 'CREATED', 1, $4, $5) 
RETURNING id, title, description, status, version, created_at, updated_at, organization_id;

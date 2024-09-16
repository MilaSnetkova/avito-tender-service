UPDATE tenders 
SET 
    title = COALESCE($1, title), 
    description = COALESCE($2, description), 
    version = version + 1,
    updated_at = NOW()
WHERE id = $3 AND status = 'CREATED'
RETURNING id, title, description, status, version, updated_at, organization_id;
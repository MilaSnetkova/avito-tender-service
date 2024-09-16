UPDATE offers
SET description = COALESCE($1, description),
    status = COALESCE($2, status),
    version = version + 1,
    updated_at = NOW()
WHERE id = $3;
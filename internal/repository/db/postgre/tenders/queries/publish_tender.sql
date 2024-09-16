UPDATE tenders 
SET status = 'PUBLISHED' 
WHERE id = $1 AND status = 'CREATED';
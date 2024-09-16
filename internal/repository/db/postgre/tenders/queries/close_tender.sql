UPDATE tenders 
SET status = 'CLOSED' 
WHERE id = $1 AND status = 'PUBLISHED';
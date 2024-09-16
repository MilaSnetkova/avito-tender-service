SELECT 
id, 
name, 
description, 
type, 
created_at, 
updated_at 
FROM organization 
WHERE id = $1;
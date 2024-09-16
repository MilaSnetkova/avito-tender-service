SELECT 
id, 
title, 
description, 
status, 
version, 
created_at, 
updated_at, 
organization_id 

FROM tenders 
WHERE id = $1
	 
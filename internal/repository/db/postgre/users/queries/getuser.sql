SELECT 
id, 
username, 
first_name, 
last_name, 
created_at, 
updated_at
FROM employee 
WHERE id = $1;
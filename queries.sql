
-- GetAllAuthor:
SELECT * FROM authors ORDER BY %s LIMIT $1 OFFSET $2

-- UpdateAuthor
UPDATE authors SET 
    name=COALESCE($1, name),
    bio=COALESCE($2,bio),
    updated_at=COALESCE($3, NOW())
WHERE id=$4;


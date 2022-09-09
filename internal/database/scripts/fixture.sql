INSERT INTO categories
VALUES (1, 'Mammifère', NULL), (2, 'Chien', 1), (3, 'Chat', 1), (4, 'Singe', 1), (5, 'Gorille', 4), (6, 'Chimpanzé', 4), (7, 'Shiba', 2), (8, 'Corgi', 2), (9, 'Labrador', 2), (10, 'Poisson', NULL), (11, 'Requin', 10), (12, 'Requin blanc', 11), (13, 'Grand requin blanc', 12), (14, 'Petit requin blanc', 12), (15, 'Requin marteau', 11), (16, 'Requin tigre', 11), (17, 'Poisson rouge', 10), (18, 'Poisson chat', 10);

WITH RECURSIVE children AS (
        SELECT
            id,
            "name",
            parent_id,
            0 as level,
            '' as path
        FROM categories
        WHERE id = 1
        UNION ALL
        SELECT
            c.id,
            c."name",
            c.parent_id,
            children.level + 1,
            children.path || children.name || ' > '
        FROM
            categories c,
            children
        WHERE
            c.parent_id = children.id
    )
SELECT *
FROM children;
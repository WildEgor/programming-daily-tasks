CREATE DATABASE work_db;

CREATE SCHEMA IF NOT EXISTS job;

CREATE TABLE IF NOT EXISTS job.employee (
    id BIGSERIAL PRIMARY KEY,
    parent_id INTEGER DEFAULT NULL,
    name VARCHAR(100) NOT NULL
);

INSERT INTO job.employee (parent_id, name)
VALUES
    (NULL, 'Bob'),  -- id=1
    (1, 'John'),    -- id=2
    (1, 'Ivan'),    -- id=3
    (2, 'Petr'),    -- id=4 OK
    (3, 'Alex');    -- id=5 OK

SELECT * FROM job.employee;

SELECT name
FROM job.employee e
WHERE NOT EXISTS (
    SELECT 1
    FROM job.employee sub
    WHERE sub.parent_id = e.id
);

-- Вывести имена всех линейных сотрудников, которые не являются руководителями.
-- Иначе говоря: вывести все листья в дереве иерархии.
WITH RECURSIVE generation AS (
    SELECT id,
           name,
           parent_id,
           0 AS level
    FROM job.employee
    WHERE parent_id IS NULL

    UNION ALL

    SELECT child.id,
           child.name,
           child.parent_id,
           level + 1
    FROM job.employee child
             JOIN generation g
                  ON g.id = child.parent_id
)
SELECT  g.name AS child_name,
        g.level,
        parent.name AS parent_name
FROM generation g
         JOIN job.employee parent
              ON g.parent_id = parent.id
ORDER BY level;

SELECT name
FROM job.employee e
WHERE e.id NOT IN (SELECT parent_id FROM job.employee WHERE parent_id IS NOT NULL);
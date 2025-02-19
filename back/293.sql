-- https://coderun.yandex.ru/problem/sql-users-active-sessions?compiler=sqlite

-- @block
DROP TABLE IF EXISTS logs;
CREATE TABLE logs (
    user_id varchar,
    "time" bigint);

INSERT INTO logs VALUES
    ('u1', 0),
    ('u1', 20),
    ('u2', 30),
    ('u2', 40),
    ('u1', 40),
    ('u1', 60),
    ('u2', 65),
    ('u1', 80),
    ('u1', 130),
    ('u1', 150),
    ('u1', 200),
    ('u1', 210);

-- SELECT * from logs;

-- @block
WITH RECURSIVE rec AS
(
    (
        SELECT user_id, array_agg(time  ORDER BY time) AS a,
            min(time) AS previous,
            0::bigint AS interval,
            0::bigint AS duration,
            0 AS cnt
        FROM logs
        group by user_id
        ORDER BY user_id
    )
    UNION ALL
    (
        SELECT step.user_id, step.a, step.previous, step.interval,
            CASE WHEN step.interval <= 30 THEN step.duration ELSE 0 END,
            step.cnt
        FROM
        (
            SELECT rec.user_id, array_remove(rec.a, rec.previous) AS a,
                rec.a[2] AS previous,
                abs( rec.previous - rec.a[2] ) AS interval,
                rec.duration + abs( rec.previous - rec.a[2] ) AS duration,
                CASE WHEN abs( rec.previous - rec.a[2] )<=30 and rec.duration=0 THEN rec.cnt+1 ELSE rec.cnt END
            FROM rec WHERE array_length( rec.a, 1 ) > 1
        ) AS step
    )
)
SELECT rec.user_id, max(rec.cnt) FROM rec
GROUP BY user_id ORDER BY user_id;
-- SELECT rec.user_id, rec.a, rec.previous, rec.interval, rec.duration, rec.cnt FROM rec
-- ORDER BY user_id, array_length(rec.a, 1) DESC;

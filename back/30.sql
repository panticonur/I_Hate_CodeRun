-- https://coderun.yandex.ru/selections/backend/problems/genres
-- @block

DROP TABLE IF EXISTS genre CASCADE;
CREATE TABLE genre (
    id bigint PRIMARY KEY,
    name varchar(100) NOT NULL,
    parent_genre_id bigint,
    FOREIGN KEY (parent_genre_id) REFERENCES genre(id));

INSERT INTO genre(id, name, parent_genre_id) VALUES
    (1, 'pop', null),
    (2, 'rock', null),
    (3, 'blues', null),
    (4, 'russian pop', 1),
    (5, 'k-pop', 1),
    (6, 'euro pop', 1),
    (7, 'hard rock', 2),
    (8, 'metal', 2),
    (9, 'punk rock', 2),
    (10, 'delta blues', 3),
    (11, 'electric blues', 3),
    (12, 'heavy metal', 8),
    (13, 'trash metal', 8),
    (14, 'post punk', 9),
    (15, 'horror punk', 9);

DROP TABLE IF EXISTS track CASCADE;
CREATE TABLE track (
    id bigint PRIMARY KEY ,
    name varchar(100) NOT NULL);

INSERT INTO track(id, name) VALUES
    (1, 'Hallowed Be Thy Name'),
    (2, 'Boys Don''t Cry'),
    (3, 'Riding With The King'),
    (4, 'You Give Love A Bad Name'),
    (5, 'Since I''ve Been Loving You');

DROP TABLE IF EXISTS track_genre CASCADE;
CREATE TABLE track_genre (
    track_id bigint,
    genre_id bigint,
    PRIMARY KEY(track_id, genre_id),
    FOREIGN KEY (track_id) REFERENCES track(id),
    FOREIGN KEY (genre_id) REFERENCES genre(id));

INSERT INTO track_genre(track_id, genre_id) VALUES
    (1, 12),
    (2, 14),
    (3, 3),
    (4, 2),
    (4, 7),
    (5, 3),
    (5, 7);

-- @block

WITH RECURSIVE Rt AS (
    SELECT
            tr.id AS track_id,
            ge.id AS genre_id,\
            tr.name AS track_name,
            ge.name AS genre_name,
            ge.parent_genre_id AS ge_parent
        FROM track AS tr
        JOIN track_genre AS gt
            ON tr.id = gt.track_id
        JOIN genre AS ge
            ON ge.id = gt.genre_id
UNION 
    SELECT
            Rt.track_id AS track_id,
            pr.id AS genre_id,
            Rt.track_name AS track_name,
            pr.name AS genre_name,
            pr.parent_genre_id AS ge_parent
        FROM Rt, genre AS pr
        where Rt.ge_parent IS NOT NULL AND pr.id = Rt.ge_parent 
)
SELECT track_id, genre_id, track_name, genre_name
FROM Rt
ORDER BY track_id, genre_id;

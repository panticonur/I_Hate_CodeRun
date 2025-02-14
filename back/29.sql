-- @conn user
-- @block COMPANY
DROP TABLE IF EXISTS requests;
CREATE TABLE IF NOT EXISTS requests (
    datetime TIMESTAMP,
    request_id UUID,
    parent_request_id UUID,
    host TEXT,
    type TEXT,
    data TEXT
);

INSERT INTO requests VALUES 
('1970-01-01 00:00:00.000', '9d7abfdc-2739-478d-bd97-5e96529a46c9',                                   NULL, 'balancer.test.yandex.ru',  'RequestReceived', ''),
('1970-01-01 00:00:00.100', '9d7abfdc-2739-478d-bd97-5e96529a46c9',                                   NULL, 'balancer.test.yandex.ru',      'RequestSent', 'backend1.ru'),
('1970-01-01 00:00:00.101', '9d7abfdc-2739-478d-bd97-5e96529a46c9',                                   NULL, 'balancer.test.yandex.ru',      'RequestSent', 'backend2.ru'),
('1970-01-01 00:00:00.150', '16629e3e-9bcf-45eb-a301-d4f7b39fe15d', '9d7abfdc-2739-478d-bd97-5e96529a46c9',             'backend1.ru',  'RequestReceived', ''),
('1970-01-01 00:00:00.200', '981a28fc-89c7-42ee-b5b6-feddb6d48d94', '9d7abfdc-2739-478d-bd97-5e96529a46c9',             'backend2.ru',  'RequestReceived', ''),
('1970-01-01 00:00:00.155', '16629e3e-9bcf-45eb-a301-d4f7b39fe15d', '9d7abfdc-2739-478d-bd97-5e96529a46c9',             'backend1.ru',      'RequestSent', 'backend3.ru'),
('1970-01-01 00:00:00.210', '981a28fc-89c7-42ee-b5b6-feddb6d48d94', '9d7abfdc-2739-478d-bd97-5e96529a46c9',             'backend2.ru',     'ResponseSent', ''),
('1970-01-01 00:00:00.200', '6690a96d-13da-44c0-bbc2-77374f7440f6', '16629e3e-9bcf-45eb-a301-d4f7b39fe15d',             'backend3.ru',  'RequestReceived', ''),
('1970-01-01 00:00:00.220', '6690a96d-13da-44c0-bbc2-77374f7440f6', '16629e3e-9bcf-45eb-a301-d4f7b39fe15d',             'backend3.ru',     'ResponseSent', ''),
('1970-01-01 00:00:00.260', '16629e3e-9bcf-45eb-a301-d4f7b39fe15d', '9d7abfdc-2739-478d-bd97-5e96529a46c9',             'backend1.ru', 'ResponseReceived', 'backend3.ru OK'),
('1970-01-01 00:00:00.300', '16629e3e-9bcf-45eb-a301-d4f7b39fe15d', '9d7abfdc-2739-478d-bd97-5e96529a46c9',             'backend1.ru',     'ResponseSent', ''),
('1970-01-01 00:00:00.310', '9d7abfdc-2739-478d-bd97-5e96529a46c9',                                   NULL, 'balancer.test.yandex.ru', 'ResponseReceived', 'backend1.ru OK'),
('1970-01-01 00:00:00.250', '9d7abfdc-2739-478d-bd97-5e96529a46c9',                                   NULL, 'balancer.test.yandex.ru', 'ResponseReceived', 'backend2.ru OK'),
('1970-01-01 00:00:00.400', '9d7abfdc-2739-478d-bd97-5e96529a46c9',                                   NULL, 'balancer.test.yandex.ru',     'ResponseSent', ''),
('1970-01-01 00:00:00.500', '2c9a2700-090b-40a1-8137-2b3390594ad7',                                   NULL, 'balancer.test.yandex.ru',  'RequestReceived', ''),
('1970-01-01 00:00:00.505', '2c9a2700-090b-40a1-8137-2b3390594ad7',                                   NULL, 'balancer.test.yandex.ru',      'RequestSent', 'backend1.ru'),
('1970-01-01 00:00:00.510', 'd4369d14-4009-453a-a8f7-4559588cfe16',	'2c9a2700-090b-40a1-8137-2b3390594ad7',             'backend1.ru',  'RequestReceived', ''),
('1970-01-01 00:00:00.700', 'd4369d14-4009-453a-a8f7-4559588cfe16',	'2c9a2700-090b-40a1-8137-2b3390594ad7',             'backend1.ru',     'ResponseSent', ''),
('1970-01-01 00:00:00.710', '2c9a2700-090b-40a1-8137-2b3390594ad7',                                   NULL, 'balancer.test.yandex.ru', 'ResponseReceived', 'backend1.ru ERROR'),
('1970-01-01 00:00:00.715', '2c9a2700-090b-40a1-8137-2b3390594ad7',                                   NULL, 'balancer.test.yandex.ru',     'ResponseSent', '');

-- @block COMPANY
select * FROM requests;

-- @block COMPANY
SELECT * FROM requests WHERE host='balancer.test.yandex.ru' AND type='RequestSent';

SELECT * FROM requests WHERE host='backend1.ru' AND type='RequestReceived';

-- @block in
SELECT (r.datetime-t.datetime) as d, * -- r.datetime, r.request_id, r.data, t.datetime, t.request_id, t.data
FROM (SELECT * FROM requests WHERE type='RequestSent') t
    left join requests r
on t.request_id=r.parent_request_id and t.data=r.host where r.type='RequestReceived';

-- @block out
SELECT (recv.datetime-send.datetime) as d, * -- r.datetime, r.request_id, r.data, t.datetime, t.request_id, t.data
FROM (SELECT * FROM requests WHERE type='ResponseReceived') recv 
    left join requests send 
on recv.request_id=send.parent_request_id where send.type='ResponseSent' and (recv.data LIKE (send.host || '%'));

-- 1) host = 'balancer.test.yandex.ru' AND type = 'RequestReceived'

-- @block cast
select CAST(to_number('299', '99G2D2S') /2  as numeric);

-- @block start
-- @label start
SELECT count(*) FROM requests WHERE host='balancer.test.yandex.ru' AND type='RequestReceived';

-- @block in and out
-- @label in and out
select CAST(to_number(to_char(sum(d), 'FF3'), '99G999D9S') / (SELECT count(*) FROM requests WHERE host='balancer.test.yandex.ru' AND type='RequestReceived')  as numeric) as avg_network_time_ms from
(
    (
        SELECT (r.datetime-s.datetime) as d, * -- r.datetime, r.request_id, r.data, t.datetime, t.request_id, t.data
        FROM (SELECT * FROM requests WHERE type='RequestSent') s
            left join requests r
        on s.request_id=r.parent_request_id and s.data=r.host where r.type='RequestReceived'
    )
    union
    (
        SELECT (recv.datetime-send.datetime) as d, * -- r.datetime, r.request_id, r.data, t.datetime, t.request_id, t.data
        FROM (SELECT * FROM requests WHERE type='ResponseReceived') recv 
            left join requests send 
        on recv.request_id=send.parent_request_id where send.type='ResponseSent' and (recv.data LIKE (send.host || '%'))
    )
);

CREATE TABLE IF NOT EXISTS short_urls
(
    id SERIAL PRIMARY KEY,
    long_url CHARACTER VARYING(255) NOT NULL,
    token CHARACTER VARYING(255) NOT NULL,
    visits INTEGER DEFAULT 0 NOT NULL
)

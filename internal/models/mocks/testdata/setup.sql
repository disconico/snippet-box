CREATE TABLE snippets
(
    id      serial primary key,
    title   varchar(100)                not null,
    content text                        not null,
    created timestamp without time zone not null,
    expires timestamp without time zone not null
);

CREATE INDEX idx_snippets_created ON snippets (created);

CREATE TABLE users
(
    id              serial primary key,
    name            varchar(255)                not null,
    email           varchar(255)                not null,
    hashed_password char(60)                    not null,
    created         timestamp without time zone not null
);

INSERT INTO users(name, email, hashed_password, created)
VALUES ('Alice Jones',
        'alice@example.com',
        '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
        '2022-01-01 09:18:24');


-- +migrate Up
CREATE TABLE public."user"
(
	id uuid DEFAULT uuid_generate_v4(),
    username text not null check(length(username) >= 3 and length(username) <= 64) UNIQUE,
    email text not null check(length(email) >= 3 and length(email) <= 128 and email like '%@%.%') UNIQUE,
    password_hash bytea NOT NULL,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    active boolean default true,
    PRIMARY KEY (id),
    CONSTRAINT username UNIQUE (username),
    CONSTRAINT email UNIQUE (email)
);


ALTER TABLE public."user"
    OWNER to iridium;

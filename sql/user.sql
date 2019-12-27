CREATE TABLE public."user"
(
    user_id serial not null
    username text not null check(length(username) >= 3 and length(username) <= 64),
    email text not null check(length(email) >= 3 and length(email) <= 128),
    password_hash bytea NOT NULL,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    active boolean default true,
    PRIMARY KEY (user_id),
    CONSTRAINT username UNIQUE (username),
    CONSTRAINT email UNIQUE (email)
);

ALTER TABLE public."user"
    OWNER to iridium;
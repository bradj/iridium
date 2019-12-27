CREATE TABLE public."user"
(
    username character varying(64) NOT NULL,
    email character varying(128) NOT NULL,
    password_hash bytea NOT NULL,
    created_date timestamp default now(),
    active boolean default true,
    PRIMARY KEY (username),
    CONSTRAINT username UNIQUE (username),
    CONSTRAINT email UNIQUE (email)
);

ALTER TABLE public."user"
    OWNER to iridium;
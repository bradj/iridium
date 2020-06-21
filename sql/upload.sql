CREATE TABLE public."upload"
(
    id serial not null,
    user_id int REFERENCES public."user"(id) not null,
    type upload_type not null,
    location text not null,
    created_at timestamp default now(),
    updated_at timestamp default null,
    PRIMARY KEY (id, user_id)
);

ALTER TABLE public."user"
    OWNER to iridium;

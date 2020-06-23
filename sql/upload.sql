-- +migrate Up
CREATE TABLE uploads
(
    id serial not null,
    user_id uuid REFERENCES user(id) not null,
    type upload_type not null,
    location text not null,
    created_at timestamp default now(),
    updated_at timestamp default null,
    PRIMARY KEY (id, user_id)
);

ALTER TABLE uploads
    OWNER to iridium;

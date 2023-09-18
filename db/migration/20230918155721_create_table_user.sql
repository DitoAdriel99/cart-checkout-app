-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users 
(
    id          uuid        not null primary key,
    fullname    varchar     not null,
    email       varchar     not null,
    password    varchar     not null,
    role        varchar     not null,
    created_at  timestamp   not null,
    updated_at  timestamp,
    is_active   bool        not null default true
);
INSERT INTO users (id,fullname,email,password,role,created_at, updated_at, is_active)
VALUES ('9f68d2a2-26da-11ee-be56-0242ac120002','Admin','admin@gmail.com','$2a$14$PqbFLLY0XZT2vgj9dafsv.uCCacUeYfiCv3zXWqR4C2dyoyXLj81K', 'admin','2023-09-01 14:42:29.977','2023-09-01 14:42:29.977', true),
       ('62abe70e-293d-11ee-be56-0242ac120002','Akuntes1','akuntes1@gmail.com','$2a$14$PqbFLLY0XZT2vgj9dafsv.uCCacUeYfiCv3zXWqR4C2dyoyXLj81K','user','2023-09-01 14:42:29.977','2023-09-01 14:42:29.977', true),
       ('62abec5e-293d-11ee-be56-0242ac120002','Akuntes2','akuntes2@gmail.com','$2a$14$PqbFLLY0XZT2vgj9dafsv.uCCacUeYfiCv3zXWqR4C2dyoyXLj81K','user','2023-09-01 14:42:29.977','2023-09-01 14:42:29.977', true) ON CONFLICT DO NOTHING;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users
-- +goose StatementEnd

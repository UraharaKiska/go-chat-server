-- +goose Up
-- +goose StatementBegin
create table chat (
    id serial primary key,
    name VARCHAR(100) not NULL,
    created_at timestamp not null default now()
);

create table chat_message (
    id serial primary key,
    -- chat_id BIGINT REFERENCES chat (id) on delete cascade,
    from_user text not null,
    message text not null,
    created_at timestamp not null default now()
);

create table  chat_user (
    id serial primary key,
    chat_id BIGINT REFERENCES chat (id) on delete cascade,
    username text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chat_user;
drop table chat_message;
drop table chat;
-- +goose StatementEnd

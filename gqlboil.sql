CREATE SCHEMA IF NOT EXISTS gqlboil;

DROP TABLE IF EXISTS "gqlboil"."users";
CREATE TABLE IF NOT EXISTS "gqlboil"."users"  (
    id serial primary key,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    deleted_at timestamp with time zone,
    email varchar(255) unique not null,
    name varchar(255) DEFAULT '',
    password text not null
);

DROP TABLE IF EXISTS "gqlboil"."pet";
CREATE TABLE IF NOT EXISTS "gqlboil"."pet"  (
    id serial primary key,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    deleted_at timestamp with time zone,
    owner_id integer references "gqlboil"."users" (id),
    name varchar(255),
    breed varchar(255)
);

DROP TABLE IF EXISTS "gqlboil"."toys";
CREATE TABLE IF NOT EXISTS "gqlboil"."toys"  (
     id serial primary key,
     created_at timestamp with time zone default now(),
     updated_at timestamp with time zone default now(),
     deleted_at timestamp with time zone,
     description varchar(255) not null,
     color varchar(255),
     pet_id integer references "gqlboil"."pet" (id)
);
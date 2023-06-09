CREATE TABLE IF NOT EXISTS users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS places
(
    id serial not null unique,
    name varchar(255) not null,
    description varchar(255),
    address varchar(255),
    rating decimal
);

CREATE TABLE IF NOT EXISTS user_places
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    place_id int references places (id) on delete cascade not null
);
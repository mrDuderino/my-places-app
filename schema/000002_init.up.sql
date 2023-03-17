CREATE TABLE IF NOT EXISTS dishes
(
    id serial not null unique,
    name varchar(255) not null,
    description varchar(255),
    rating decimal
);

CREATE TABLE IF NOT EXISTS place_dishes
(
    id serial not null unique,
    place_id int references places (id) on delete cascade not null,
    dish_id int references dishes (id) on delete cascade not null
);
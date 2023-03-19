CREATE TABLE IF NOT EXISTS discount_cards
(
    id serial not null unique,
    number varchar(255) not null,
    description varchar(255),
    valid_from date,
    valid_to date
);

CREATE TABLE IF NOT EXISTS place_discount_cards
(
    id serial not null unique,
    place_id int references places (id) on delete cascade not null,
    discount_card_id int references discount_cards (id) on delete cascade not null
);
CREATE TABLE qodice (
    id serial not null,
    
--The author's info
   author_id integer not null,
   name varchar not null,
   profession varchar,
   nationality varchar,
   birth_year int,
   birth_month varchar,
   birth_date int,
   death_year int,
   death_month varchar,
   death_date int,

--The quote's info
   quote text NOT NULL unique,
   quote_id integer not null,

    date date unique not null default current_date,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz
)
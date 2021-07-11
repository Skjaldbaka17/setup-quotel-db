CREATE TABLE quotes(
   id SERIAL PRIMARY KEY,

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
   count integer default 0,
   is_icelandic boolean default false,

   created_at timestamptz default current_timestamp,
   updated_at timestamptz,
   deleted_at timestamptz,

--TSV for searching
   quote_tsv tsvector,
   tsv tsvector,

   FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE
);
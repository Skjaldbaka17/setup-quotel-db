CREATE TABLE authors(
   id SERIAL PRIMARY KEY,
   name VARCHAR NOT NULL UNIQUE,
   profession varchar,
   nationality varchar,
   birth_year int,
   birth_month varchar,
   birth_date int,
   death_year int,
   death_month varchar,
   death_date int,
   count integer default 0,
  
   nr_of_english_quotes integer default 0,
   nr_of_icelandic_quotes integer default 0,
    
   name_tsv tsvector,

    created_at timestamptz default current_timestamp,
   updated_at timestamptz,
   deleted_at timestamptz,
);
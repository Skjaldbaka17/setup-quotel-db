CREATE EXTENSION if not exists pg_trgm;

UPDATE authors SET name_tsv = setweight(to_tsvector('english', name), 'A');
UPDATE quotes SET quote_tsv = setweight(to_tsvector('english', quote), 'B');
UPDATE quotes SET tsv = setweight(to_tsvector('english', quote), 'B') || setweight(to_tsvector('english', name), 'A');

CREATE INDEX if not exists index_authors_on_name_tsv ON authors USING gin(name_tsv);
CREATE INDEX if not exists index_authors_on_name ON authors(name);
CREATE INDEX if not exists index_authors_on_count ON authors(count);
CREATE INDEX if not exists index_authors_on_profession ON authors(profession);
CREATE INDEX if not exists index_authors_on_nationality ON authors(nationality);
CREATE INDEX if not exists index_authors_on_nr_of_english_quotes ON authors(nr_of_english_quotes);
CREATE INDEX if not exists index_authors_on_nr_of_icelandic_quotes ON authors(nr_of_icelandic_quotes);
CREATE INDEX if not exists index_authors_on_birth_year ON authors(birth_year);
CREATE INDEX if not exists index_authors_on_birth_month ON authors(birth_month);
CREATE INDEX if not exists index_authors_on_birth_date ON authors(birth_date);
CREATE INDEX if not exists index_authors_on_death_year ON authors(death_year);
CREATE INDEX if not exists index_authors_on_death_month ON authors(death_month);
CREATE INDEX if not exists index_authors_on_death_date ON authors(death_date);
CREATE INDEX if not exists index_authors_on_death_day ON authors(death_day);
CREATE INDEX if not exists index_authors_on_birth_day ON authors(birth_day);

CREATE INDEX if not exists index_aods_on_date ON aods(date);
CREATE INDEX if not exists index_aodices_on_date ON aodices(date);

CREATE INDEX if not exists index_quotes_on_tsv ON quotes USING gin(tsv);
CREATE INDEX if not exists index_quotes_on_quote_tsv ON quotes USING gin(quote_tsv);
CREATE INDEX if not exists index_quotes_on_quote ON quotes(quote);
CREATE INDEX if not exists index_quotes_on_author_id ON quotes(author_id);
CREATE INDEX if not exists index_quotes_on_count ON quotes(count);

CREATE INDEX if not exists index_qods_on_date ON qods(date);
CREATE INDEX if not exists index_qodices_on_date ON qodices(date);




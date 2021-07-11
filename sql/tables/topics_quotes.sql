CREATE TABLE topics_quotes(
   id SERIAL PRIMARY KEY,
   topic_id int,
   quote_id int,
   created_at timestamptz default current_timestamp,
   updated_at timestamptz,
   deleted_at timestamptz
);
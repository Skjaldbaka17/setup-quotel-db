CREATE INDEX if not exists index_topics_view_on_name_tsv ON topicsView using gin(name_tsv);
CREATE INDEX if not exists index_topics_view_on_quote_tsv ON topicsView using gin(quote_tsv);
CREATE INDEX if not exists index_topics_view_on_tsv ON topicsView using gin(tsv);
CREATE INDEX if not exists index_topics_view_on_author_id ON topicsView(author_id);
CREATE INDEX if not exists index_topics_view_on_quote_id ON topicsView(id);
CREATE INDEX if not exists index_topics_view_on_topic_id ON topicsView(topic_id);
CREATE INDEX if not exists index_topics_view_on_topic_name ON topicsView(topic_name);
CREATE INDEX if not exists index_topics_view_on_topic_name_id ON topicsView(topic_name,id);

CREATE INDEX words_idx ON unique_lexeme USING gin(word gin_trgm_ops);
CREATE INDEX words_idx_quotes ON unique_lexeme_quotes USING gin(word gin_trgm_ops);
CREATE INDEX words_idx_authors ON unique_lexeme_authors USING gin(word gin_trgm_ops);
CREATE MATERIALIZED VIEW topicsView as 
select authors.id as author_id,
       authors.name,
       q.id as id,
       q.quote as quote,
       q.is_icelandic as is_icelandic,
       authors.name_tsv || q.tsv  as tsv,
       authors.name_tsv,
       q.tsv as quote_tsv,
       t.name as topic_name,
       t.id as topic_id,
       q.nationality,
       q.profession,
       q.birth_year,
       q.birth_month,
       q.birth_date,
       q.death_year,
       q.death_month,
       q.death_date
from authors
   inner join quotes q
      on authors.id = q.author_id
   inner join topics_quotes ttq
      on q.id = ttq.quote_id
   inner join topics t
      on t.id = ttq.topic_id;
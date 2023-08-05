CREATE INDEX IF NOT EXISTS infinitybottles_name_idx ON infinitybottles USING GIN (to_tsvector('simple', bottle_name));

CREATE INDEX IF NOT EXISTS contributions_brand_name_idx ON contributions USING GIN (to_tsvector('simple', brand_name));
CREATE INDEX IF NOT EXISTS contributions_tags_idx ON contributions (tags);

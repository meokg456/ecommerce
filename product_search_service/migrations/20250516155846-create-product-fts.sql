-- +migrate Up
ALTER TABLE products
ADD COLUMN tsv tsvector;

UPDATE products
SET tsv = setweight(to_tsvector(coalesce(title,'')), 'A') ||
    setweight(to_tsvector(coalesce(descriptions,'')), 'C');

CREATE INDEX products_tsv_idx ON products USING gin(tsv);

-- +migrate StatementBegin
CREATE FUNCTION products_tsv_trigger() 
RETURNS trigger AS $$
BEGIN
  new.tsv :=
     setweight(to_tsvector(coalesce(new.title,'')), 'A') ||
     setweight(to_tsvector(coalesce(new.descriptions,'')), 'C');
  return new;
END
$$ LANGUAGE plpgsql;
-- +migrate StatementEnd

CREATE TRIGGER tsvectorupdate BEFORE INSERT OR UPDATE
    ON products FOR EACH ROW EXECUTE FUNCTION products_tsv_trigger();

-- +migrate Down
ALTER TABLE products
DROP COLUMN tsv;
DROP TRIGGER tsvectorupdate ON products;
DROP FUNCTION products_tsv_trigger();
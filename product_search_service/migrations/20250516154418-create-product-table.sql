
-- +migrate Up
CREATE TABLE "products" (
  "id" serial PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "descriptions" text,
  "category" varchar(255),
  "images" text[],
  "additionInfo" json,
  "merchantId" integer,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "products";
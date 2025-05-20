
-- +migrate Up
CREATE TABLE "products" (
  "id" varchar(255) PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "descriptions" text,
  "category" varchar(255),
  "images" text[],
  "addition_info" json,
  "merchant_id" integer,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "products";
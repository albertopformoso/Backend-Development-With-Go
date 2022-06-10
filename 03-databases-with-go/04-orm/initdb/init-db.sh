#!/bin/bash
set -ex

psql -v ON_ERROR_STOP=1 -U "$POSTGRES_USER" -d "$POSTGRES_DB" <<-EOSQL
    CREATE TABLE IF NOT EXISTS "products" (
        "id" SERIAL NOT NULL,
        "name" VARCHAR(25) NOT NULL,
        "observations" VARCHAR(100),
        "price" NUMERIC(5, 2) NOT NULL,
        "created_at" TIMESTAMP NOT NULL DEFAULT now(),
        "updated_at" TIMESTAMP,
        CONSTRAINT products_id_pk PRIMARY KEY (id)
    );

    CREATE TABLE IF NOT EXISTS "invoice_headers" (
        "id" SERIAL NOT NULL,
        "client" VARCHAR NOT NULL,
        "created_at" TIMESTAMP NOT NULL DEFAULT now(),
        "updated_at" TIMESTAMP,
        CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id)
        
    );

    CREATE TABLE IF NOT EXISTS "invoice_items" (
        "id" SERIAL NOT NULL,
        "invoice_header_id" SERIAL,
        "product_id" SERIAL,
        "created_at" TIMESTAMP NOT NULL DEFAULT now(),
        "updated_at" TIMESTAMP,
        CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),
        CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers(id) ON UPDATE RESTRICT ON DELETE RESTRICT,
        CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE RESTRICT ON DELETE RESTRICT
    );

EOSQL
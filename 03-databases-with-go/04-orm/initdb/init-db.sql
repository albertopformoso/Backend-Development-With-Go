CREATE DATABASE IF NOT EXISTS mysql_db;
USE mysql_db;

CREATE TABLE IF NOT EXISTS `products` (
    `id` INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name` VARCHAR(25) NOT NULL,
    `observations` VARCHAR(100),
    `price` DECIMAL(5, 2) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT now(),
    `updated_at` TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `invoice_headers` (
    `id` INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `client` VARCHAR(25) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT now(),
    `updated_at` TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `invoice_items`(
    `id` INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `invoice_header_id` INT NOT NULL,
    `product_id` INT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT now(),
    `updated_at` TIMESTAMP,
    CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);
-- +goose Up
-- +goose StatementBegin
CREATE TABLE `orders` (
	id int(10) unsigned NOT NULL AUTO_INCREMENT,
    transaction_number varchar(50) NOT NULL,
    media varchar(50) NOT NULL,
    customer_id int(10) unsigned NOT NULL,
    is_member int(1) NOT NULL,
    customer_name varchar(150) NULL,
    customer_email varchar(150) NULL,
    customer_phone varchar(50) NULL,
    customer_address varchar(255) NULL,
    created_at timestamp NULL DEFAULT NULL,
    updated_at timestamp NULL DEFAULT NULL,
    deleted_at timestamp NULL DEFAULT NULL,
	PRIMARY KEY (id),
	KEY orders_id_index (id),
    KEY `orders_customer_id_foreign` (`customer_id`),
    CONSTRAINT orders_customer_id_foreign FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE CASCADE    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `orders`;
-- +goose StatementEnd

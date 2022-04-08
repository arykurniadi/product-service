-- +goose Up
-- +goose StatementBegin
CREATE TABLE `order_items` (
	id int(10) unsigned NOT NULL AUTO_INCREMENT,
    order_id int(10) unsigned NOT NULL,
    name varchar(150) NOT NULL,
    price double(12,2) NOT NULL DEFAULT 0,
    created_at timestamp NULL DEFAULT NULL,
    updated_at timestamp NULL DEFAULT NULL,
	PRIMARY KEY (id),
	KEY order_items_id_index (id),
    KEY `order_items_order_id_foreign` (`order_id`),
    CONSTRAINT order_items_id_foreign FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `order_items`;
-- +goose StatementEnd

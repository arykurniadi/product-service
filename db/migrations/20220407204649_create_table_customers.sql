-- +goose Up
-- +goose StatementBegin
CREATE TABLE `customers` (
	id int(10) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(150) NOT NULL,
    email varchar(150) NOT NULL,
    phone varchar(30) NOT NULL,
    address varchar(255) NULL,
	PRIMARY KEY (id),
	KEY customers_id_index (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `customers`;
-- +goose StatementEnd

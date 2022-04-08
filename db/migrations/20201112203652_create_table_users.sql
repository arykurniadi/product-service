-- +goose Up
-- +goose StatementBegin
CREATE TABLE `users` (
	id int(10) unsigned NOT NULL AUTO_INCREMENT,
	username varchar(200) NOT NULL,
    password varchar(200) NOT NULL,
	email varchar(200) NOT NULL,
	role varchar(200) NOT NULL,
	PRIMARY KEY (id),
	KEY users_id_index (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `users`;
-- +goose StatementEnd

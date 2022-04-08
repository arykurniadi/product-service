-- +goose Up
-- +goose StatementBegin
CREATE TABLE `roles` (
	id varchar(50) NOT NULL,
    name varchar(200) NOT NULL,
	PRIMARY KEY (id),
	KEY roles_id_index (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `roles`;
-- +goose StatementEnd

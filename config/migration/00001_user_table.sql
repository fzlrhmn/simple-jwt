-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE tbl_user (
  id varchar(50) NOT NULL,
  username varchar(250),
  password varchar(250),
  created_at timestamp with time zone DEFAULT NOW(),
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  PRIMARY KEY (id)
);

CREATE UNIQUE INDEX tbl_user_id ON tbl_user (id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE tbl_user;
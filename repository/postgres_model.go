package repository

// User holds user table information
type User struct {
	tableName struct{} `sql:"tbl_user"`
	ID        string   `sql:",pk"`
	Username  string
	Password  string
}

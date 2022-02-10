package database

const (
	Getuser = "SELECT * FROM users WHERE id = $1;"
	EnterUser = "INSERT INTO users(\"F_name\", \"L_name\", \"M_name\") VALUES($1,$2,$3) returning id;"
)
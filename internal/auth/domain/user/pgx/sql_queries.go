package pgx

const (
	userCreateQuery     = "insert into users (id, firstname, lastname, login, hash) values (:id, :firstname, :lastname, :login, :hash) returning id"
	userUpdateQuery     = "update users set  firstname = :firstname,lastname= :lastname  where id = :id"
	userGetByLoginQuery = "SELECT id,firstname,lastname, login, hash from users where login = $1 "
	userGetByIDQuery    = "SELECT id,firstname,lastname, login, hash from users where id = $1 "
)

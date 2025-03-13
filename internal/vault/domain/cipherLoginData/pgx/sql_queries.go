package pgx

const (
	createQuery = "insert into cipher_login_data (id, uri, login, password, user_id) values (:id, :uri, :login, :password, :user_id) returning id"
	updateQuery = "update cipher_login_data set  login = :login  where id = :id"
)

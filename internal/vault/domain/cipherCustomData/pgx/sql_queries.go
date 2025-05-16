package pgx

const (
	createQuery = "insert into cipher_custom_data (id, key, value, meta_data, user_id) values (:id, :key, :value, :meta_data, :user_id) returning id"
	updateQuery = "update cipher_custom_data set  key = :key  where id = :id"
)

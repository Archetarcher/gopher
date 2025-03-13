package pgx

const (
	createQuery = "insert into cipher_card_data (id, card_holder_name, brand,number, exp_month, exp_year, code, meta_data, user_id) values (:id, :card_holder_name, :brand,number, :exp_month, :exp_year, :code, :meta_data, :user_id) returning id"
	updateQuery = "update cipher_card_data set  card_holder_name = :card_holder_name  where id = :id"
)

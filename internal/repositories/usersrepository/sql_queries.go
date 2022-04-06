package usersrepository

const (
	createUserQuery = "INSERT INTO users_store (public_id, payload) VALUES ($1, $2) RETURNING id"
	getUserQuery    = "SELECT id, public_id, payload FROM users_store WHERE public_id=$1"
)

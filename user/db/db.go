package db

import "github.com/A-Victory/user-mig/user/models"

func (db *DBconn) CreateUser(user models.User) error {

	_, err := db.DB.Exec("INSERT INTO users (email, password, first_name, last_name) VALUES ($1, $2, $3, $4)",
		user.Email, user.Password, user.FirstName, user.LastName)
	return err

}

func (db *DBconn) GetUser(email string) (*models.User, error) {
	user := &models.User{}
	row := db.DB.QueryRow("SELECT id, email, password, first_name, last_name FROM users WHERE email=$1", email)
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName); err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DBconn) ListUsers() ([]models.User, error) {
	rows, err := db.DB.Query("SELECT id, email, first_name, last_name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

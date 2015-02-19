package manager

import(
	UserDAO "RX/services/model/UserDAO"
)

func Login(username string, password string) (map[string][]string, error) {
	return UserDAO.Login(username, password)
}

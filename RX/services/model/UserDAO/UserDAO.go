package UserDAO

import(
    DBConn "RX/services/model/DBConn"
)

func Login(username string, password string) (map[string][]string, error) {
	var db,_ = DBConn.DBConn()
    stmt := `SELECT * FROM public."Users" WHERE "UserName" = '`+username+`' AND "Password" ='`+password+`'`
    res, err := db.Query(stmt)
    defer db.Close()
    return DBConn.ReadResultSet(res),err
}
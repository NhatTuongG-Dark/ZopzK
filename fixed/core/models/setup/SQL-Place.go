package SetupBuild

import (
	"Yami/core/models/Json"
	YamiDB "Yami/core/db"
	"log"
	"math/rand"
	"os"
	"time"
)

func InsertTables() bool {
	Rows, error := YamiDB.SQL.Query("CREATE TABLE `attacks` (`ID` int(10) unsigned NOT NULL AUTO_INCREMENT, `Username` varchar(20) NOT NULL, `Target` varchar(255) NOT NULL, `Method` varchar(20) NOT NULL, `Type` varchar(20) NOT NULL, `Port` int(11) NOT NULL, `Duration` int(11) NOT NULL, `End` bigint(20) NOT NULL, `Created` bigint(20) NOT NULL, PRIMARY KEY (`ID`),  KEY `username` (`Username`));")
	if error != nil || Rows.Err() != nil {
		return false
	}


	Row, err := YamiDB.SQL.Query("CREATE TABLE `users` (`ID` int(10) unsigned NOT NULL AUTO_INCREMENT, `Username` varchar(64), `Password` varchar(128), `MFA` varchar(200) NOT NULL, `NewUser` tinyint(1), `Admin` tinyint(1), `Reseller` tinyint(1), `Banned` tinyint(1), `Vip` tinyint(1), `MaxTime` int(10) UNSIGNED DEFAULT NULL, `Cooldown` int(10) UNSIGNED DEFAULT NULL, `Concurrents` int(10) UNSIGNED DEFAULT NULL, `MaxSessions` int(10) UNSIGNED DEFAULT NULL, `PowerSavingExempt` tinyint(1), `BypassBlacklist` tinyint(1), `PlanExpiry` BigInt(20), PRIMARY KEY (`ID`), KEY `username` (`Username`));"); if err != nil || Row.Err() != nil {
		log.Println("Failed to Place User Table")
		return false
	}

	Password := RandStringBytes(10)

	Pass := YamiDB.HashPassword(Password)

	PlanExpiry := time.Now().Add(time.Hour * 8760)




	Row, err = YamiDB.SQL.Query("INSERT INTO `users` (`ID`, `Username`, `Password`, `MFA`, `Admin`, `Reseller`, `Banned`, `Vip`, `NewUser`, `MaxTime`, `Cooldown`, `Concurrents`, `MaxSessions`, `PowerSavingExempt`, `BypassBlacklist`, `PlanExpiry`) VALUES (NULL, ?, ?, 0, 1, 0, 0, 1, 1, 1200, 30, 3, 1, 0, 0, ?);", JsonParse.ConfigSyncs.SQL.SQLAudit.Username, Pass, PlanExpiry.Unix()); if err != nil || Row.Err() != nil {
		log.Println("Failed to Place User Table", err)
		return false
	}

	File, err := os.Create(JsonParse.ConfigSyncs.SQL.SQLAudit.LogDetails); if err != nil {
		return false
	}

	File.WriteString(JsonParse.ConfigSyncs.SQL.SQLAudit.Username+"\r\n"+string(Password))

	File.Close()

	log.Printf("	- Username: "+JsonParse.ConfigSyncs.SQL.SQLAudit.Username)
	log.Printf("	- Password: "+Password)
	log.Printf("	- Expiry: "+PlanExpiry.Format("Mon Jan _2 15:04:05 2006"))

	return true
}



// *from stackoverflow.

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}
package SetupBuild

import (
	"Yami/core/db"
)

func CheckTableExist() bool {

	_, Check := YamiDB.SQL.Query("SELECT * FROM `users`, `attacks`;")

	if Check == nil {
		return true
	} else {
		return false
	}
}

/*
CREATE TABLE `attacks` (
     `ID` int(10) unsigned NOT NULL AUTO_INCREMENT, 
     `Username` varchar(20) NOT NULL, 
     `Target` varchar(255) NOT NULL, 
     `Method` varchar(20) NOT NULL, 
     `Port` int(11) NOT NULL, 
     `Duration` int(11) NOT NULL, 
     `End` bigint(20) NOT NULL, 
     `Created` bigint(20) NOT NULL, 
     PRIMARY KEY (`ID`),  
     KEY `username` (`Username`)
);*/
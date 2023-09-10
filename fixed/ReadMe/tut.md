# Tutorial for Yami


`Your server must be Ubuntu 21.04 Or 18`, `Support for other versions is coming soon!`


# SQL Server Installation
1. Installing Mysql or any other SQL Server
          Ubuntu: *sudo apt update*, *sudo apt install mysql-server*, *sudo mysql_secure_installation*, *sudo mysql*

2. Create a new database and give it a name. *CREATE DATABASE `Yami`;*, *c*
3. Create a new user which has access to your SQL Database
          - *CREATE USER 'LicenseV2'@'localhost' IDENTIFIED BY '439083409843340984098freokgreoekrgK$D';*
          - *GRANT ALL PRIVILEGES ON * . * TO 'LicenseV2'@'localhost';*
          - *FLUSH PRIVILEGES;*
4. Now sync your database infomation with your ./build/config.json
     + Current Info
     - Username: `YamiC2`
     - Password: `439083409843340984098freokgreoekrgK$D`
     - Database Name: `Yami`
     - Database Host: `localhost:3306`

# Starting Yami
1. Installing screen:
     `sudo apt install screen`
2. Screening the cnc:
     `sudo setcap 'cap_net_bind_service=+ep' ./Yami`
     `chmod 777 *`
     `screen ./Yami`

# Please note that FB reserves the right to revoke your license at any time, you agree to this by screening the cnc
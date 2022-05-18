# Trademarkia_GoLang_Task_DatingApp

# Stumble Dating app Backend in GoLang

## _By Patel Jesal Manoj, 19BCE1259_ 


### Features And Technologies

- GoLang and ThunderClient for GET/PUT/DELETE/POST Requests off of API
- Postgre SQL RDBMS with GoLang ORM (GORM)
- Deadlock protection
- Pre populated Schema in database access
- 100% Completed

### Libraries
- Gorilla MUX ("github.com/gorilla/mux")
- GORM Databse IO ("gorm.io/gorm")
- GORM Postgre Driver ("gorm.io/driver/postgres")
- Miscellaneous basic libraries of GoLang

## Working
- After importing the data from json files in the Postgre SQL database, the Database access url is set to ' dsn = "host=localhost user=postgres password=jes123 dbname=goLangjesalMPtask port=5432 sslmode=disable TimeZone=Asia/Shanghai" '
- Database name : goLangjesalMPtask
	Database port : 5432
	Database Username : postgre
	Database host : localhost
	Database Password : jes123
- GoLang LocalHost : 8081
	Endpoints : 
		localhost:8081/availablematches/
		localhost:8081/usersindistance/
		localhost:8081/usernamesearch/
- After connecting to data base and running the initialMigration() function, 2 tables are generated according to database attributes


### GetAvailableMatches endpoint (localhost:8081/availablematches/)
	GET Method
	Use of Hashmap has been made to form pairs of userIDs who like each other and then data of those respective ids are extracted from users database.
	JSON array of 2 users is returned.

### GetUsersInDistance endpoint (localhost:8081/usersindistance/)
	GET Method
	endpoint format : /usersindistance/{userid}/{distance}
	linear search in users database is made to find users within distance.
	JSON array of all users within givin {distance} of user with ID = {userid}

### GetUsersWithGivenStringInName (localhost:8081/usernamesearch/)
	GET Method
	endpoint format : /usernamesearch/{nametag}
	linear search in users database is made to find users within {nametag} in their name as a substring.
	names are compared after converting to lowercase for maximum efficient search.
	substring method of strings library in GoLang is used.
	JSON array of all users with {nametag} as a substring in their name is returned.


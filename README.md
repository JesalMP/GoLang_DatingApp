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

- 
![image](https://user-images.githubusercontent.com/84318539/169162056-80007db3-50ba-41b9-933c-ce103375fd68.png)
![image](https://user-images.githubusercontent.com/84318539/169162125-10d841b8-2523-4bb6-b317-45e9899764e5.png)
![image](https://user-images.githubusercontent.com/84318539/169162165-eee5c0aa-328b-4baf-a0a4-104c8a94da7b.png)






### <u> GetAvailableMatches endpoint (localhost:8081/availablematches/)
- GET Method
- Use of Hashmap has been made to form pairs of userIDs who like each other and then data of those respective ids are extracted from users database.
- JSON array of 2 users is returned.
![image](https://user-images.githubusercontent.com/84318539/169162387-d57cff1a-1a83-4499-b926-e98e311c5840.png)
![image](https://user-images.githubusercontent.com/84318539/169162469-b79d5e6b-02a9-4bf8-9dfb-cc61bc8ffc8f.png)
![image](https://user-images.githubusercontent.com/84318539/169162482-239f6644-f89a-40ce-8730-90118832936a.png)
 
- Total 3 pairs of available matches found.
- - userID 8 and 34
- - userID 36 and 66
- - userID 11 and 81
	
	
	
	
	
### <u> GetUsersInDistance endpoint (localhost:8081/usersindistance/)
- GET Method
- endpoint format : /usersindistance/{userid}/{distance}
- linear search in users database is made to find users within distance.
- JSON array of all users within givin {distance} of user with ID = {userid}
![image](https://user-images.githubusercontent.com/84318539/169162807-812f00fb-87f4-4bc7-8d19-c7c88781ee82.png)
- 6 users are near userID 16 with distance 1
![image](https://user-images.githubusercontent.com/84318539/169162957-e29d63b0-a6c3-4759-a7bb-98749d8c32bd.png)
- 15 users are near userID 47 with distance 8
	
	
	
	
	

### <u> GetUsersWithGivenStringInName (localhost:8081/usernamesearch/)
- GET Method
- endpoint format : /usernamesearch/{nametag}
- linear search in users database is made to find users within {nametag} in their name as a substring.
- names are compared after converting to lowercase for maximum efficient search.
- substring method of strings library in GoLang is used.
- JSON array of all users with {nametag} as a substring in their name is returned.
![image](https://user-images.githubusercontent.com/84318539/169163116-f5951f79-a1ee-4678-8d51-d40e207739f0.png)
- 7 users with "ll" in their name
![image](https://user-images.githubusercontent.com/84318539/169163188-c5139e1a-e77e-445a-9612-c8cdcbd5827e.png)
- Complete name of "Auguste"



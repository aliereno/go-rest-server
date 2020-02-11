# Rest API Server with Go

* go
* gin
* gorm
* postgresql

### Installation & Run:
 
* Before running, you should set the database informations on  ***internal/orm/env.go***. 

Then on terminal;
```
scripts/run.sh
```
For build;
```
scripts/build.sh
```
## API
#### /files
* `GET` : Static folder.
### Without Token
#### /login
* `POST` : `{email:"a@b.com",password:"pass"}`
#### /register
* `POST` : `{name:"username",email:"a@b.com",password:"pass"}`
#### /books
* `GET` : Returns all books.
#### /books/:id
* `GET` : Returns book detail.

### With User Token
#### /user/settings
* `GET` : Returns user settings.
#### /user/settings
* `POST` : Update user settings.
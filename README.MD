# golang web application

to launch the application type the following command

```
go run ./snippetbox/cmd/web
```

### routes and allowed methods to them

- /snippet/view - GET 
- /snippet/delete - DELETE
- /snippet/create - POST

## SQL commands

in order to interact with SQL, it is mandatory to install the following package through the terminal 

```
go get github.com/go-sql-driver/mysql@v1.0.3
```

and go.sum file will appear

```
mysql -D snippetbox -u web -p
```
this one connects you to the table with the name of "snippetbox"

to gain package

- `-D snippetbox`: This option specifies the name of the database you want to use, in this case, "snippetbox". It selects the database with the name "snippetbox" for the current session.
- `-u web` This option specifies the MySQL user to log in as, in this case, "web". The -u flag is followed by the username.
- `-p`: This option prompts you for the MySQL user's password. When you run the command, MySQL will ask you to enter the password for the specified user ("web" in this case).

 - if you wish to change your current database just use the command 

```
USE YOUR_DATABASE_NAME;
```

- or if you wish to show the available tables in the current database

```
SHOW TABLES;
```

- if you wish to quit from the terminal just type
```
mysql> exit
```

- to generate a TLS certificate you need to type the following command
```
go run /opt/homebrew/Cellar/go/1.21.6/libexec/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

if you're receiving 400 http error, and enable to proceed to the website then, visit the following link

https://localhost:4000/

instead of using 

http://localhost:4000/


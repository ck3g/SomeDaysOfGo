# Golang testing

Hands-on project from Udemy course [Unit, integration and functional Testing in Golang](https://www.udemy.com/share/1020jMA0QaeF1aTHg=/).


## Database

```
mysql> CREATE DATABASE go_sql_client_example;
mysql> USE go_sql_client_example;
mysql> CREATE TABLE users (id BIGINT NOT NULL AUTO_INCREMENT, first_name VARCHAR(20), last_name VARCHAR(20), email VARCHAR(50), date_created DATETIME, PRIMARY KEY(id));
mysql> INSERT INTO users (first_name, last_name, email, date_created) VALUES ('John', 'Doe', 'john.doe@example.com', NOW());
mysql> SELECT * FROM users;
```
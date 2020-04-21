# Working with Postgres

## Postgress commands

#### Login

```
psql postgres
```

#### List databases

```
\l
```

#### Connect to a database


```
\c <database name>
```

#### Create database


```
CREATE DATABASE go_company;
```


#### Create table

```
CREATE TABLE employees (
   ID INT PRIMARY KEY     NOT NULL,
   NAME           TEXT    NOT NULL,
   RANK           INT     NOT NULL,
   ADDRESS        CHAR(50),
   SALARY         REAL DEFAULT 25500.00,
   BDAY			      DATE DEFAULT '1900-01-01'
);
```

#### Show tables in a database


```
\d
```


#### Show table details


```
\d employees
```
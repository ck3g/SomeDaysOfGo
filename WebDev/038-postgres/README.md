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


#### Create tables

```
CREATE TABLE employees (
   ID  SERIAL PRIMARY KEY NOT NULL,
   NAME           TEXT    NOT NULL,
   SCORE            INT     DEFAULT 10 NOT NULL,
   SALARY         REAL
);

CREATE TABLE phonenumbers (
   ID  SERIAL PRIMARY KEY NOT NULL,
   PHONE           CHAR(50) NOT NULL,
   EMP_ID         INT      references employees(ID)
);
```

#### Insert some records

```
INSERT INTO employees (NAME,SCORE,SALARY)
VALUES  ('Daniel', 23, 55000.00),
        ('Arin', 25, 65000.00),
        ('Juan', 24, 72000.00),
        ('Shen', 26, 64000.00),
        ('Myke', 27, 58000.00),
        ('McLeod', 26, 72000.00),
        ('James', 32, 35000.00);


INSERT INTO phonenumbers (PHONE,EMP_ID)
VALUES ('555-777-8888', 4), ('555-222-3345', 4), ('777-543-3451', 1), ('544-756-2334', 2);
```

#### Show tables in a database


```
\d
```


#### Show table details


```
\d employees
```

#### View data

```
SELECT * FROM employees;
SELECT * FROM phonenumbers;
```
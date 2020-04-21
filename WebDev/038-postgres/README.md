# Working with Postgres

## Dependencies

```
$ go get github.com/lib/pq
```

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
CREATE DATABASE go_bookstore;
```


#### Create tables

```
CREATE TABLE books (
  isbn    char(14)     PRIMARY KEY NOT NULL,
  title   varchar(255) NOT NULL,
  author  varchar(255) NOT NULL,
  price   decimal(5,2) NOT NULL
);
```

#### Insert some records

```
INSERT INTO books (isbn, title, author, price) VALUES
('0134494164', 'Clean Architecture', 'Robert C. Martin', 24.99),
('B06XPJML5D', 'Designing Data-Intensive Applications', 'Martin Kleppmann', 40.99),
('B07LCM8RG2', 'Refactoring', 'Martin Fowler', 38.99);
```

#### Show tables in a database


```
\d
```


#### Show table details


```
\d books
```

#### View data

```
SELECT * FROM books;
```
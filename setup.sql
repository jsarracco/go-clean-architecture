CREATE TABLE greetings (
    id integer not null primary key, 
    date datetime, 
    author varchar(50), 
    content varchar(50)
);

-- run to create sqlite schema
-- sqlite3 /var/tmp/guestbook.sqlite < setup.sql
-- http://localhost:8080/
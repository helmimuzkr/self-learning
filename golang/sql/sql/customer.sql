ALTER TABLE customer
    DROP COLUMN rating;

DESC customer;

INSERT INTO customer (name, email, balance, birth_date,married)
    VALUES 
        ('Rein','reinsftr@gmail.com', 100000,'1999-10-05',false),
        ('Kanza', 'Kanza@gmail.com', 100000, '2016-10-05', false),
        ('Hafiz', null, null, null, true);

SELECT * FROM customer;

CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(25) NOT NULL,
    username VARCHAR(25) NOT NULL,
    password VARCHAR(25) NOT NULL
);

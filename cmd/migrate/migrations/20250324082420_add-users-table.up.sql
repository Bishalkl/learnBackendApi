CREATE TABLE users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	firstName VARCHAR(255) NOT NULL,
	lastName VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	createdAt TIMESTAMP NOT NULL
);

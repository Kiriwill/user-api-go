BEGIN;
	CREATE TABLE customer (
	  id int PRIMARY KEY AUTO_INCREMENT,
	  name varchar(255) NOT NULL,
	  birthdate timestamp NOT NULL,
	  email varchar(150) NOT NULL,
	  password varchar(255) NOT NULL,
	  
	  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

	  CONSTRAINT un_email UNIQUE(email)
	);

	CREATE TABLE address (
	  id int PRIMARY KEY AUTO_INCREMENT,
	  user_id int NOT NULL,
	  street_address varchar(255),
	  street_number varchar(100),
	  city varchar(255),
	  region varchar(255),
	  postal_code varchar(300) ,
	  country varchar(255) NOT NULL,
	  
	  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

	  FOREIGN KEY (user_id) REFERENCES customer(id) ON DELETE CASCADE
	);

COMMIT;



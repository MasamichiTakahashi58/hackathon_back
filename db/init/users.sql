DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  email VARCHAR(100) UNIQUE NOT NULL,
  username VARCHAR(50) UNIQUE NOT NULL,
  display_name VARCHAR(100),
  profile_image VARCHAR(255),
  header_image VARCHAR(255),
  bio VARCHAR(160),
  location VARCHAR(100),
  website VARCHAR(255),
  birthdate DATE,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);


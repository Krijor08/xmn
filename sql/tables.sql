CREATE TABLE roles (
	ID INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_time TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
	role NVARCHAR(50) NOT NULL
) AUTO_INCREMENT = 1000;

CREATE  TABLE users (
	ID INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_time TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    username NVARCHAR(100) NOT NULL,
    password CHAR(64) NOT NULL,
    email NVARCHAR(150),
    phone INT,
    role_ID INT NOT NULL DEFAULT 1000,
    FOREIGN KEY (role_ID) REFERENCES roles(ID)
) AUTO_INCREMENT = 1000;

CREATE OR REPLACE VIEW user_view AS
SELECT 
	u.ID,
    u.created_time AS user_created,
    u.updated_time AS user_updated,
    r.created_time AS role_created,
    r.updated_time AS role_updated,
    u.username,
    u.password,
    u.email,
    u.phone,
    r.role
FROM users AS u
JOIN roles AS r ON
	u.role_ID = r.ID;
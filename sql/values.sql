INSERT INTO roles (role)
VALUES ('guest'), ('admin'), ('user');

INSERT INTO users (username, password, email, phone, role_ID)
VALUES 
('user', '1234', 'user@test.com', 12345678, 1002),
('test', 'test', NULL, NULL, DEFAULT);
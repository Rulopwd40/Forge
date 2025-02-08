CREATE TABLE User (
    username VARCHAR(20) PRIMARY KEY auto_increment,
    password VARCHAR(128),
    name VARCHAR(40),
    level INTEGER
);

CREATE TABLE Trainee (
    username VARCHAR(20) PRIMARY KEY,
    weight DECIMAL(5,2),
    height INTEGER,
    FOREIGN KEY (username) REFERENCES User(username)
);
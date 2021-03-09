USE homestead;

CREATE TABLE users(
    user_id INT(11) UNSIGNED AUTO_INCREMENT NOT NULL,
    user_name VARCHAR(30) NOT NULL,
    token VARCHAR(30) NOT NULL UNIQUE,
    PRIMARY KEY (user_id)
);

CREATE TABLE characters(
    character_id INT(11) UNSIGNED AUTO_INCREMENT NOT NULL,
    character_name VARCHAR(30) NOT NULL,
    times FLOAT(4,4) NOT NULL,
    PRIMARY KEY (character_id)
);

CREATE TABLE user_characters(
    user_character_id INT(11) UNSIGNED AUTO_INCREMENT NOT NULL,
    user_id INT(11) UNSIGNED NOT NULL,
    character_id INT(11) UNSIGNED NOT NULL,
    PRIMARY KEY (user_character_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (character_id) REFERENCES characters(character_id)
);

INSERT INTO users(user_name,token) VALUES ("test","5555");
INSERT INTO characters(character_name,times) VALUES("testCharacter",0.4444);
INSERT INTO user_characters(user_id,character_id) VALUES(1,1);
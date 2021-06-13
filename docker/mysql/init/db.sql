/* 本番用のDBに流し込むSQL */
USE game_db;

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
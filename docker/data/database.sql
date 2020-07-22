/*
    This file is used by the docker-compose build command to build the mysql db
    
    Tables:
    * article : stores articles (id, title, header, authors, created_on, updated_on)
    * section : stores sections (id, title, paragraph, position, media, created_on, updated_on,
    parent_id, article_id)
    * comment : stores comments (id, title, content, author, created_on, updated_on)
*/

CREATE TABLE article (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT, 
    title VARCHAR(255),
    header TEXT NOT NULL,
    authors TEXT NOT NULL,
    created_on DATETIME,
    updated_on DATETIME,
    PRIMARY KEY (id)
);

CREATE TABLE section (
     id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
     title VARCHAR(255),
     paragraph TEXT,
     position INT UNSIGNED NOT NULL,
     media TEXT,
     created_on DATETIME,
     updated_on DATETIME,
     parent_id BIGINT UNSIGNED DEFAULT NULL, 
     article_id BIGINT UNSIGNED NOT NULL,
     PRIMARY KEY (id),
     UNIQUE KEY position (position, article_id),
     FOREIGN KEY (article_id) REFERENCES article(id) ON DELETE CASCADE, 
     FOREIGN KEY (parent_id) REFERENCES section(id) ON DELETE CASCADE
);

CREATE TABLE comment (
     id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
     title VARCHAR(255),
     content TEXT,
     author INT UNSIGNED NOT NULL,
     created_on DATETIME,
     updated_on DATETIME,
     PRIMARY KEY (id)
);

/* 
    Default values
*/
INSERT INTO `article` (`id`, `title`, `header`, `authors`, `created_on`, `updated_on`) VALUES
('1', 'Bonjour', 'helloworld', 'AT', '2020-07-22 15:40:14', '2020-07-22 15:40:14'),
('2', 'Bonsoir', 'byeworld', 'TA', '2020-07-22 15:40:37', '2020-07-22 15:40:37'),
('3', 'Bonne nuit', 'good night world', 'Marchand de Sable', '2020-07-22 15:41:11', '2020-07-22 15:41:11');

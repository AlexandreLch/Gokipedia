/*
    This file is used by the docker-compose build command to build the mysql db
    
    Tables:
    * category : stores categories (id, name, desc, creation, update)
    * image : stores images (id, name, desc, type, creation, update, category ID)
    * tag : stores tags (id, name, creation date)
    * image_tag : links images to tags by ids (Many to Many relation)
*/

CREATE TABLE article (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT, 
    title VARCHAR(255),
    header TEXT,
    authors TEXT,
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
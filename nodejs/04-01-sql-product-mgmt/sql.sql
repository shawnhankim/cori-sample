/*
 * Create database and table of goods 
 *
 *   $ mysql -u root 
 *
 *   MariaDB [(none)]> CREATE DATABASE sample_db;
 *
 *   $ mysql -u root sample_db < sql.sql
 *
 *   MariaDB [(none)]> use sample_db;
 *
 *   MariaDB [(none)]> show tables;
 *     +---------------------+
 *     | Tables_in_sample_db |
 *     +---------------------+
 *     | goods               |
 *     +---------------------+
 *
 *   MariaDB [sample_db]> describe goods;
 *     +-------------+--------------+------+-----+---------+----------------+
 *     | Field       | Type         | Null | Key | Default | Extra          |
 *     +-------------+--------------+------+-----+---------+----------------+
 *     | id          | int(11)      | NO   | PRI | NULL    | auto_increment |
 *     | name        | varchar(128) | NO   |     | NULL    |                |
 *     | category    | varchar(128) | NO   |     | NULL    |                |
 *     | price       | int(11)      | NO   |     | NULL    |                |
 *     | description | text         | NO   |     | NULL    |                |
 *     +-------------+--------------+------+-----+---------+----------------+
 */
CREATE TABLE IF NOT EXISTS `goods` (
  `id`          int          NOT NULL AUTO_INCREMENT,
  `name`        varchar(128) NOT NULL,
  `category`    varchar(128) NOT NULL,
  `price`       int          NOT NULL,
  `description` text         NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
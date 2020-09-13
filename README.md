A utility script to import Mailgun bounces data to a Mysql database 

Create .env file
````
db.host: xxx
db.port: xxx
db.user: xxx
db.pass: xxxx
db.schema: xx
db.poolsize: 2
db.idlesize: 1
mailgun.domain: xxx
mailgun.key: xx
fetcher.complaints: 1
fetcher.bounces: 1
fetcher.unsubscribes: 1
````

Ensure that your RDS schema has the tables 

````
CREATE TABLE `MailgunBounces` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain` varchar(30) DEFAULT NULL,
  `code` varchar(30) DEFAULT NULL,
  `email` varchar(90) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `error` text,
  `updatedon` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `email` (`email`)
) ENGINE=InnoDB;

CREATE TABLE `MailgunComplaints` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain` varchar(30) DEFAULT NULL,
  `email` varchar(90) DEFAULT NULL,
  `complaints_count` int(32) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updatedon` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `email` (`email`)
) ENGINE=InnoDB;
 
CREATE TABLE `MailgunUnsubscribes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain` varchar(30) DEFAULT NULL,
  `email` varchar(90) DEFAULT NULL,
  `tags` text,
  `created_at` datetime DEFAULT NULL,
  `updatedon` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `email` (`email`)
) ENGINE=InnoDB;
 
````


go run .
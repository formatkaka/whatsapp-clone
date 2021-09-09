-- Db operations done

create table users(
   id INT NOT NULL AUTO_INCREMENT,
   userName VARCHAR(100) NOT NULL,
   userEmail VARCHAR(40),
   submission_date DATE,
   PRIMARY KEY ( id )
);
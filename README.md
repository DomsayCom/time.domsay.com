# time.domsay.com
time.domsay.com

DB

CREATE SCHEMA `time` DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci ;

- users
	- id
	- username
	- password
	- active
	- created
	- created_id
	- modified
	- modified_id

CREATE TABLE `time`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(45) NULL,
  `password` VARCHAR(45) NULL,
  `active` INT NULL,
  `created` DATETIME NULL,
  `created_id` INT NULL,
  `modified` DATETIME NULL,
  `modified_id` INT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `username_UNIQUE` (`username` ASC));



- users_permissions
	- id
	- user_id
	- controller
	- action
	- active
	- created
	- created_id
	- modified
	- modified_id

CREATE TABLE `time`.`users_permissions` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NULL,
  `controller` VARCHAR(45) NULL,
  `action` VARCHAR(45) NULL,
  `active` INT NULL,
  `created` DATETIME NULL,
  `created_id` INT NULL,
  `modified` DATETIME NULL,
  `modified_id` INT NULL,
  PRIMARY KEY (`id`));


- subjects
	- id
	- vat
	- name
	- address
	- created
	- created_id
	- modified
	- modified_id

CREATE TABLE `time`.`subjects` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `vat` VARCHAR(45) NULL,
  `name` VARCHAR(45) NULL,
  `address` VARCHAR(45) NULL,
  `created` DATETIME NULL,
  `created_id` INT NULL,
  `modified` DATETIME NULL,
  `modified_id` INT NULL,
  PRIMARY KEY (`id`));


- projects
	- id
	- subject_id
	- name
	- created
	- created_id
	- modified
	- modified_id

CREATE TABLE `time`.`projects` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `subject_id` INT NULL,
  `name` VARCHAR(45) NULL,
  `created` DATETIME NULL,
  `created_id` INT NULL,
  `modified` DATETIME NULL,
  `modified_id` INT NULL,
  PRIMARY KEY (`id`));


- tasks
	- id
	- project_id
	- name
	- desc
	- start
	- end
	- created
	- created_id
	- modified
	- modified_id

CREATE TABLE `time`.`tasks` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `project_id` INT NULL,
  `name` VARCHAR(150) NULL,
  `desc` MEDIUMTEXT NULL,
  `start` DATETIME NULL,
  `end` DATETIME NULL,
  `created` DATETIME NULL,
  `created_id` INT NULL,
  `modified` DATETIME NULL,
  `modified_id` INT NULL,
  PRIMARY KEY (`id`));



VIEWS

/
/login
/logout
/dashboard

/reports

/users
/users/add
/users/edit/{id}

/subjects
/subjects/add
/subjects/edit/{id}

/projects
/projects/add
/projects/edit/{id}

/tasks
/tasks/add
/tasks/edit/{id}

DROP table IF EXISTS class CASCADE;
DROP table IF EXISTS sex CASCADE;
DROP table IF EXISTS students CASCADE;

CREATE table class (
	id			SERIAL PRIMARY KEY,
	name		VARCHAR(255) NOT NULL UNIQUE,
	deleted		BOOLEAN DEFAULT false,
	timestamp 	TIMESTAMP DEFAULT current_timestamp
);

CREATE table sex (
	id			SERIAL PRIMARY KEY,
	sex			VARCHAR(255) NOT NULL UNIQUE,
	timestamp 	TIMESTAMP DEFAULT current_timestamp
);

CREATE table students (
	id						SERIAL PRIMARY KEY,
	chinese_name			VARCHAR(255) NOT NULL,
	pinyin					VARCHAR(255) NOT NULL,
	english_name			VARCHAR(255) NOT NULL,
	student_id				VARCHAR(255) NOT NULL,
	class_id				INT NOT NULL,
	sex_id					INT NOT NULL,
	deleted					BOOLEAN DEFAULT false,
	timestamp 				TIMESTAMP DEFAULT current_timestamp,
	FOREIGN KEY(class_id) 	REFERENCES class(id),
	FOREIGN KEY(sex_id) 	REFERENCES sex(id)	
);

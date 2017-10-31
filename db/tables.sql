DROP table IF EXISTS class;
CREATE table class (
	id		SERIAL PRIMARY KEY,
	name	VARCHAR(255) NOT NULL UNIQUE
);

DROP table IF EXISTS sex;
CREATE table sex (
	id		SERIAL PRIMARY KEY,
	sex		VARCHAR(255) NOT NULL UNIQUE
);

DROP table IF EXISTS students;
CREATE table students (
	id						SERIAL PRIMARY KEY,
	chinese_name			VARCHAR(255) NOT NULL,
	pinyin					VARCHAR(255) NOT NULL,
	english_name			VARCHAR(255) NOT NULL,
	student_id				VARCHAR(255) NOT NULL,
	class_id				INT NOT NULL,
	sex_id					INT NOT NULL,
	FOREIGN KEY(class_id) 	REFERENCES class(id),
	FOREIGN KEY(sex_id) 	REFERENCES sex(id)	
);

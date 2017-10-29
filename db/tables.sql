PRAGMA foreign_keys = ON;

DROP table IF EXISTS class;
CREATE table class (
	id		INTEGER PRIMARY KEY AUTOINCREMENT,
	name	TEXT NOT NULL UNIQUE
);

DROP table IF EXISTS sex;
CREATE table sex (
	id		INTEGER PRIMARY KEY AUTOINCREMENT,
	sex		TEXT NOT NULL UNIQUE
);

DROP table IF EXISTS students;
CREATE table students (
	id						INTEGER PRIMARY KEY AUTOINCREMENT,
	chinese_name			TEXT NOT NULL,
	pinyin					TEXT NOT NULL,
	english_name			TEXT NOT NULL,
	student_id				TEXT NOT NULL,
	FOREIGN KEY(id) 	REFERENCES class(id),
	FOREIGN KEY(id) 	REFERENCES sex(id)	
);

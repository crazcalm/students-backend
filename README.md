# No official name yet
This project is an attempt to port the command line apps I have been building for class to a web app.

## Example of User curl Post
	curl -H "Content-Type: application/json" -X POST -d '{"name":"xyz","password":"xyz"}' http://localhost:8825/hello

## Example of Student Curl Post
	curl -H "Content-Type: application/json" -X POST -d '{"chinese_name":"李慧珍","pinyin":"li3hui4zhen1", "class":"tv show", "sex":"female", "english_name":"Jane", "student_id":"001"}' http://localhost:8825/student

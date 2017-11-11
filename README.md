# No official name yet
This project is an attempt to port the command line apps I have been building for class to a web app.

## Example of Student Curl Post
	curl -H "Content-Type: application/json" -X POST -d '{"chinese_name":"李慧珍","pinyin":"li3hui4zhen1", "class_id":"1", "sex_id":"2", "english_name":"Jane", "student_id":"001"}' http://localhost:8825/student

## Example of Student Curl PUT
	curl -h "Content-Type: application/json" -X PUT -d '{"id": "1", "chinese_name": "new_name", "pinyin":"testing", "class_id":"1", "sex_id"1", "english_name": "testing", "student_id": "007"}' http://localhost:8825/student

## Postgres
	sudo -i -u postgres

## Config file
rename sample_config.json to config.json and fill in the values for your database configuration.
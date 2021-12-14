package site

var ExperienceItemsQuery = `SELECT item_type, title, description, start_date, end_date
	 				   		FROM experience_items;`

var NoteListQuery = `SELECT title, topic, content
	 				 FROM notes;`

var NoteQuery = `SELECT title, topic, content
	 			 FROM notes
				 WHERE title=$1
				 LIMIT 1;`

var ProjectListQuery = `SELECT title, description, github_link, status, start_date, end_date
	 					FROM projects;`

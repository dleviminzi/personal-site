INSERT INTO notes (title, topic, content) VALUES ('TestNote1', 'tests', 'This is a test'),
                                                 ('TestNote2', 'tests', 'This is a test'),
                                                 ('TestNote3', 'tests', 'This is a test'),
                                                 ('TestNote4', 'tests', 'This is a test'),
                                                 ('TestNote5', 'tests', 'This is a test'),
                                                 ('TestNote6', 'tests', 'This is a test');

INSERT INTO experience_items (item_type, title, description, start_date, end_date) VALUES ('tests', 'Test EI 1', 'This is another test', '2019-01-05 00:00:00.000', '2021-03-06 00:00:00.000'),
                                                                                          ('tests', 'Test EI 2', 'This is another test', '2019-01-06 00:00:00.000', '2021-01-06 00:00:00.000'),
                                                                                          ('tests', 'Test EI 3', 'This is another test', '2019-01-07 00:00:00.000', '2021-02-06 00:00:00.000'),
                                                                                          ('tests', 'Test EI 4', 'This is another test', '2019-01-08 00:00:00.000', '2021-04-06 00:00:00.000'),
                                                                                          ('tests', 'Test EI 5', 'This is another test', '2019-01-09 00:00:00.000', '2021-06-06 00:00:00.000'),
                                                                                          ('tests', 'Test EI 6', 'This is another test', '2019-01-15 00:00:00.000', '2021-08-06 00:00:00.000'),
                                                                                          ('tests', 'Test EI 7', 'This is another test', '2019-01-25 00:00:00.000', '2021-07-06 00:00:00.000');

INSERT INTO projects (title, description, github_link, status, start_date, end_date) VALUES ('Test P 1', 'This is a test', 'githooob', 'finished', '2019-01-05 00:00:00.000', '2021-03-06 00:00:00.000'),
                                                                                            ('Test P 2', 'This is a test', 'githooob', 'finished', '2019-01-06 00:00:00.000', '2021-03-16 00:00:00.000'),
                                                                                            ('Test P 3', 'This is a test', 'githooob', 'finished', '2019-01-08 00:00:00.000', '2021-06-06 00:00:00.000'),
                                                                                            ('Test P 4', 'This is a test', 'githooob', 'ongoing', '2019-01-15 00:00:00.000', '2021-09-06 00:00:00.000');

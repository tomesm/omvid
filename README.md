# Virtual Degree

An experimental project in which everyone can create its own path to a “degree” from public available MOOC courses.
You can assemble your own curriculum of courses which would cover the topic of a chosen subject in a University.

This project has just started (April 2019) and its under development :)

## API implemented

| Method  | Pattern  | Handler  | Action  |
|---|---|---|---|
| ANY | /  | home  | Display the home page  |
| ANY | /course  | showCourse  | Display a specific course  |
| POST | /course/create  | createCourse  | Create a new course  |

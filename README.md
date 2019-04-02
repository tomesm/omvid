# Virtual Degree

An experimental project (web app) in which everyone can create its own path/roadmap to a “degree” from public available MOOC courses.
You can assemble your own curriculum of courses which would cover as much topics as possible of a chosen subject.

The idea is to gain knowledge about a subject which is close to knowledge which you would normally receive in a University.

This project has just started (April 2019) and it is still under development :) It is based on great [Let's Go professional](https://lets-go.alexedwards.net) tutorial book which I am following replacing the original *snippet* app with my own.

## API implemented

| Method  | Pattern  | Handler  | Action  |
|---|---|---|---|
| ANY | /  | home  | Display the home page  |
| ANY | /course  | showCourse  | Display a specific course  |
| POST | /course/create  | createCourse  | Create a new course  |

## Project structure

### cmd
The *cmd* directory conatins the application-specific code for web and cli application.


### pkg
The *pkg* contains the ancillary non-application-specific and reusable code like database models, validation helpers etc.

### ui
The *ui* folder conatins user-interface assets like HTML templates, static (CSS) files etc.

Templates are named based on the following pattern: `<name>.<role>.tmpl`  where `<role>` is either *page*, *partial* or *master* (the layout).

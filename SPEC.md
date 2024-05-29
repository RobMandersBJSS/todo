# Part 1

[X] Use a variadic function to print a list of 10 todos.
[X] Output a list of 10 todos in JSON format.
[X] Use a variadic function to output 10 todos in JSON file.
[X] Read todos from a JSON file.
[X] Create console program to read 10 todos from a JSON file and display them on screen.
[X] Simulate a race condition... 
    [X] When one Goroutine updates a variable with odd numbers.
    [X] And another Goroutine updates the variable with even numbers.
    [X] Then after each update, attempt to display the variable.
[X] Refactor program to use channels and mutexes.
[ ] Print a list of todos and their current status, using two goroutines which alternate between items and statuses.

# Part 2

[X] Create CLI app to manage todo list stored in memory. Create, Read, Update, Delete.
[X] Create web page app to manage todo list stores in memory.
[X] Create server that can concurrently receive a list of pre-defined commands. The following should be avaiable via specific commands:
    [X] Server status.
    [X] Task status.
[X] Create a web API to receive web page actions (remote commands) applied to todo list.

# Stretch Goals

[ ] Extend web API to receive actions for todo list stored in a database.
[ ] Extend web API to receive actions to be applied to todos from multiple users. Each user should have their own table.
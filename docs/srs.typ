= personal dashboard software requirement specification 

Personal dashboard to keep up with personal work & todos ( tracking time, overview of task deadlines, eisenhower method, brag document ), social connections & events ( friend database, reconnection reminders ), health & fitness ( habits such as taking multivitamins & working out ), daily journaling, maybe even study related stuff


== work & todos

- add work category / job
- list work categories / jobs


=== todo management / planning

- add todo ( name, description, deadline / priority )
- list todos
- start tracking time for job
- stop tracking time for job
- list time spent on a job in the last week


brag documents

right now the todos can be displayed in a table
[ work, todo name, deadline if any , priority, description ],

in the future we could draw a dependency diagram

priority could be a number 1 through 3.
- 1 is the lowest priority ("nice to have")
- 2 is the medium priority ("should have")
- 3 is the highest priority ("as soon as possible")

UI should come first

h2 for work & todos

start tracking time section
- input for selecting a work category
- start button
- stop button ( if started )
- spent time ( if started )
- history of time "sprints" or some other name (table)

input with add button (form) for creating a new work category

table for listing work categories and time spent each day on them
(let's see about 7 last days)

another table for listing work categories and time spent each week on them
(let's see about 4 last weeks)

form for creating a new todo (job, name, description, deadline, priority)
priority should be a select from 1 through 3 ( 1 is default )

table displaying unfinished todos
[ job, name, deadline, priority, description, action ]
unfinished todos are sorted first by priority, then by deadline.
if two todos have the same priority, the one with the earlier deadline comes first.
no deadline todos come last.

action for unfinished todos is a button to mark the todo as finished
( for now no edit button is displayed, but is planned )


table displaying finished todos
( job, name, deadline, priority, description )

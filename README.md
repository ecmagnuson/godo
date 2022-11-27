# godo

A command line todo application heavily inspired by [todo.txt-cli](https://github.com/todotxt/todo.txt-cli) using [cobra](https://github.com/spf13/cobra), [GORM](https://github.com/go-gorm/gorm), and [go-sqlite3](https://github.com/mattn/go-sqlite3).

### Design
A task is split into 8 fields, with each field corresponding to one column in a sqlite3 database.

```
ID        int       // unique ID
Priority  string    // urgent Task prefixed with "!"
Task      string    // "sweep kitchen floor"
Location  string    // "@home"
Project   string    // "+cleaning"
Created   time.Time // time.Now() called when Task created
Completed time.Time // time.Now() called when Task is done
Todo      bool      // true when created, false when done
```

## Philosophy

    I like to format my tasks as actionable events that have a location
    "clean the fridge @home"

    tasks can have priority associated with them with `!`
    "! send important email @work"

    Optionally, tasks can have a project associated with it
    "clean the fridge @home +cleaning"

this way I can split various projects into each location, for example

    "clean fridge @home +cleaning"
    "pay credit card bill @home finance"

both of these tasks are `@home` but they have different `+project` tags associated with them.
this underlying though process forms the basis of `godo`

the following is a list of the functionaility of `godo`

- add one task at a time
- add multiple tasks at a time
- list every tasks with a unique `@location`
- list every task with a unique `+project`
- list every high priority task prefixed with `!`
- list every unique `@location` currently used in todo.db
- list every unique `+project` currently used in todo.db
- list every unique `@location +project` currently used in todo.db
- do one task at a time
- do multiple tasks at a time

notes:

- When a task is completed, none of the `ls` methods will list it anymore
- task do not need a location. If one is not given, it will be assigned `@unknown`
- if a task does not have an `@location` the current implemtation does not allow an `+project`


## Usage
`godo setup` will create a `.todo` directory in the users home directory.

## Add
adding a task for the first time will create `todo.db` inside of `~/.todo/`

### one at a time

arguments after `add` command will be added to the database

`godo add complete assignment @school`

arguments can be prefixed with `!` to mark it as high priority

`godo add ! send important email to boss @work`

arguments can be given a `+project` as well

`godo add clean fridge @home +cleaning`

### multiple tasks at once
when `godo add` is called with no arguments you can to add multiple tasks at once

```
> ! call bank @home +finance
> grade papers @school +grading
> finish my part of presentation @school
>
```
To break out of this and add each line - corresponding to each task - just hit `Enter ↵` on a blank line and it will add each task to the todo.db

## List
high priority tasks will always print first

list everything in `todo.db`

```
godo ls
2 ! send important email to boss @work
4 ! call bank @home +finance
1   complete assignment @school
3   clean fridge @home +cleaning
5   grade papers @school +grading
6   finish my part of presentation @school
```

list everything `@home`

```
godo ls @home
4 ! call bank @home +finance
3   clean fridge @home +cleaning
```

list every taks with `+finance` tag

```
godo ls +finance
4 ! call bank @home +finance
```

list every task with a high priority

```
godo ls !
2 ! send important email to boss @work
4 ! call bank @home +finance
```

Sometimes I forget the exact `@location` or `+project` I have used in the `todo.db`, so I have implemented a subcommand `ls all`

list all `@location` tags being used in the `todo.db`

```
godo ls all @
@home
@school
@work
```

list all `+project` tags being used in the `todo.db`

```
godo ls all +
+cleaning
+finance
+grading
```

list every `@home +project` unique tag used in the `todo.db`

```
@home +cleaning
@home +finance
@school +grading
```

## Do
do a task one at a time with its id

```
godo do 3
```

do multiple tasks at once

```
godo do 1 2 4
```

### Requirements
[Go](https://go.dev/)

### Installation
- `git clone https://github.com/ecmagnuson/godo.git`
- `cd godo`
- `go build godo`

### TODO
- implement listing of done tasks


### Future?
- Reimpliment with raw SQL instead of GORM

















Build instructions for Windows  
https://stackoverflow.com/questions/41566495/golang-how-to-cross-compile-on-linux-for-windows
`GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build`   


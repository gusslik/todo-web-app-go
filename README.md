# Todo List Backend

This is the backend of a todo list web app written in Go

This app implements all essential CRUD operations, such as: Fetching Todos, Creating Todos, Updating Todos, Deleting Todos 

This app utilizes a Postgres database, all the credentials of which are stored in a seperate .env file in the /cmd directory
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DBHOST` = host of your database

`DBPORT` = port of your database

`DBUSER` = user of your database (postgres in the provided schema)

`DBPASSWORD` = password of your database

`DBNAME` = name of your database


## API Reference

#### Get all todos

```http
  GET /api/tasks
```

#### Create todo

```http
  POST /api/tasks
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `task_name`      | `string` | **Required**. Name of task to create |

#### Update todo
```http
  PUT /api/tasks
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `task_id`      | `int` | **Required**. Id of task to update |
| `task_name`      | `string` |  Name of task to create |

#### Delete todo
```http
  DELETE /api/tasks
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `task_id`      | `int` | **Required**. Id of task to delete |

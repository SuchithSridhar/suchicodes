# Configurations Using Environment Variables

| ConfigKey        | Purpose                              |
| ---------------- | ------------------------------------ |
| `ENVIRONMENT`    | The environment program launching in |
| `DATABASE_URI`   | The database location                |
| `ROOT_DIRECTORY` | The root directory of the web app    |
| `SECRET_KEY`     | The secret key for login encryption  |
| `LISTEN_PORT`    | The port to listen on                |


- `ENVIRONMENT`: Expects one of `development`, `production`, `debug`, `testing`.
- `DATABASE_URI`: Expects a valid database URI format such as `sqlite:/home/suchi/database.sqlite` 

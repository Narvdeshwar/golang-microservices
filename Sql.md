# RUN db in interactive mode
```bash
docker exec -it user-db psql -U user -d userdb
```
| Part          | Meaning                                      | Description                                                                           |
| ------------- | -------------------------------------------- | ------------------------------------------------------------------------------------- |
| `docker exec` | Run a command **inside a running container** | Lets you interact with a container that’s already up.                                 |
| `-it`         | **Interactive + TTY mode**                   | `-i` keeps STDIN open, `-t` gives you a terminal — so you can type commands.          |
| `user-db`     | The **container name**                       | Must match the name in your `docker-compose.yml` (here it’s your Postgres container). |
| `psql`        | PostgreSQL **command-line client**           | This is the program you use to run SQL commands manually.                             |
| `-U user`     | **Username** to connect with                 | Must match `POSTGRES_USER` in your Docker environment.                                |
| `-d userdb`   | **Database name**                            | Must match `POSTGRES_DB` in your environment.                                         |
---

| Command                | Description                  |
| ---------------------- | ---------------------------- |
| `\l`                   | List all databases           |
| `\c userdb`            | Connect to `userdb` database |
| `\dt`                  | List all tables              |
| `SELECT * FROM users;` | Query a table                |
| `\q`                   | Exit PostgreSQL shell        |

---
### Run SQL directly
```bash
docker exec -it user-db psql -U user -d userdb -c "SELECT * FROM users;"

```
---
### Export Database (Backup)
```bash
docker exec -t user-db pg_dump -U user userdb > backup.sql
```
Explanation:
- Dumps the contents of the userdb database into a backup.sql file on your local system.

### Restore Database (Import Backup)
```bash
cat backup.sql | docker exec -i user-db psql -U user -d userdb
```

Explanation:
- Restores your backup SQL file into the userdb database.

### Query for creation of database
```sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    password varchar(100)
);

```
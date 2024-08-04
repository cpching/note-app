# How to build and execute?
## Preparation
### Database
1. Start a MySQL server instance using docker
    ``` bash
    docker run --name mysql-note-app -e MYSQL_ROOT_PASSWORD=<dbPassword> -p <dbIp>:<dbPort>:3306 -d mysql:latest
    ```
2. Connect the MySQL container
    ``` bash
    docker exec -it mysql-note-app mysql -uroot -p
    ```
3. Create a database `note_app` in mysql server, and use it
    ``` sql
    CREATE DATABASE note_app;
    USE note_app;
    ```
4. Create a table `notes`
    ``` sql
    CREATE TABLE notes (
        id int auto_increament primary key,
        title varchar(255) not null,
        content text,
        created_at timestamp default current_timestamp,
        modified_at timestamp default current_timestamp on update current_timestamp
    );
    ```
    
### Environment Variables
- Create a `.env` file under `backend` directory
    ```
    PORT=<noteAppPort>
    MYSQLDB_URI=root:<dbPassword>@tcp(<dbIp>:<dbPort>)/note_app
    ```


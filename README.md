# Password-Manager-Go
 Password manager backend API built using Golang and POSTGRESQL.
 
 # Features
 - Lets users sign up with username and password combination
 - Lets users store and retrieve their account credentials 
 
 # Requirements
 - Golang 1.20
 - PostGreSQL 14
 
 # Installation
 
 1. Make sure you install all the requirements mentioned above.
 
 2. Create a `.env` file in the project root directory. Add the below line:
 
  ```
  CONNSTR = "<database_name>://<postgres_username>:<postgres_password>@localhost:<port_number>/password_manager?sslmode=disable"
  ```
 
 
   Replace the placeholders with database configuration. This will set up the environment variables to run the server
    
 3. Run the below commands in PSQL terminal to create the database tables:
 
 ```
     CREATE DATABASE password_manager;
 ```
 
 ```
     CREATE TABLE users(
        userid int,
        username varchar,
        password varchar
        PRIMARY KEY( userid )
     );
     
     CREATE TABLE credentials(
        userid int,
        cred_desc varchar,
        cred_id varchar,
        cred_pass varchar
     )
 ```
 
 4. Run the below command in project root directory to start the server
 
 ```
      go run ./cmd/api     
 ```
 
 
 
 
  

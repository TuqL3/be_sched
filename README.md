# Schedule 
## Be PTIT practical laboratory center
### 1. Clone
```
    $ git clone https://github.com/TuqL3/be_sched.git
```
### 2. Config env
```
    Create env file and change parameters
    
    PORT = your port
    APP_ENV = local

    DB_HOST = your db host
    DB_PORT = your db port
    DB_DATABASE = your db database
    DB_USERNAME = your db username
    DB_PASSWORD = your db password
    DB_SCHEMA = your db schema
    JWT_KEY = your jwt key


    PORT=8080
    APP_ENV=local
    
    DB_HOST=localhost
    DB_PORT=5432
    DB_DATABASE=graduate
    DB_USERNAME=postgres
    DB_PASSWORD=123456
    DB_SCHEMA=public
    JWT_KEY=secretKey

    .air.toml windown
    bin = "tmp\\main.exe"
    cmd = "go build -o ./tmp/main.exe ."

    .air.toml linux
    bin = "tmp/main"                 
    cmd = "go build -o ./tmp/main ."  
    
```
### 3. Run
```shell
  air
```

```shel
  username: admin
  password: 12345678
```

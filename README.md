# account-book
An "Account Book" backend implementation in Golang + Gin framework.  
  
# Import
```
1. http/https server framework
go get -u github.com/gin-gonic/gin

2. Database framework
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

3. Set database connection string as environment variable
go get github.com/joho/godotenv

4. Database version control
go get -u github.com/golang-migrate/migrate/v4

5. Data Validation
go get github.com/go-playground/validator/v10

6. Test
go get github.com/stretchr/testify
```

# Usage
Connect using RESTful API  

1. Create User  
  ```
  POST 127.0.0.1/acctbook/user
  JSON:
  {
      "name": "Yami",
      "uid": "yami0001",
      "pwd": "12345678",
      "currency": "TWD"
  }
  ```  
2. Get User Data
  ```
  GET 127.0.0.1/acctbook/user
  JSON:
  {
      "uid": "yami0001",
      "pwd": "12345678"
  }
  ```  

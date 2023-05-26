# account-book
An "Account Book" backend implementation in Golang + Gin framework.  
  
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

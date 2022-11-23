# case-go

This project is an example of CRUD API.

### Stacks used:
- Golang
- Echo Framework
- Gorm
- Mysql

### Routes
```
- GET    /products         Get all products
- GET    /products/id      Get prodcust by id
- POST   /products         Create a new product
- PUT    /products/:id     Update a product
- DELETE /products/:id     Delete a product
```
### Struct Product

```
Product {
  Name  string
  Brand string
  Price float64
}

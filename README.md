# Software Engineer Challange

## Features
- User Authentication: Create an account or log in securely.
- Products Checkout: Browse, select, and redeem available gifts.
- Role-Based Access Control: Admins manage the gift catalog.
- Token-Based Authentication: Secure API endpoints.
## Tech Specifications

### Libraries/Frameworks Used

- Go (version 1.19)
- Gorilla Mux
- PostgreSQL

### Architecture/Modularity

The project follows a clean architecture pattern, separating concerns into different layers:

- **Presentation Layer**: Contains the API handlers and controllers.
- **Service Layer**: Contains the business logic.
- **Repository Layer**: Deals with data storage and retrieval.
- **Database Layer**: Connects to the database.

## Quick Start
### API Documentation
#### Login
```bash
curl --location --request POST 'localhost:3030/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "admin@gmail.com",
    "password": "admin"
}'
```
#### Register
```bash
curl --location --request POST 'localhost:3030/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "fullname": "abdur",
    "email": "ragil@gmail.com",
    "password": "admin"
}'
```
#### Create Gift
```bash
curl --location --request POST 'localhost:3030/product' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFkbWluIiwiZW1haWwiOiJhZG1pbkBnbWFpbC5jb20iLCJleHAiOjE2OTUwNjAzOTN9.pF1LWtqmM_KoUnePBVib-h0nLyvpr2qKQRlnEjV35cg' \
--header 'Content-Type: application/json' \
--data-raw '{
  "title": "test 6",
  "description": "This is a sample product description.Teasasas",
  "price": 232333,
  "quantity": 1111,
  "image": "product_image.jpg",
  "type": "HP",
  "banner": "product_banner.jpg",
  "info": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry'\''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
}'
```
#### Get Product
```bash
curl --location --request GET 'localhost:3030/product?order_by=created_at&order_type=asc' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFkbWluIiwiZW1haWwiOiJhZG1pbkBnbWFpbC5jb20iLCJleHAiOjE2OTUwNjAzOTN9.pF1LWtqmM_KoUnePBVib-h0nLyvpr2qKQRlnEjV35cg'
```
#### Get Detail Product
```bash
curl --location --request GET 'localhost:3030/product/58a90ea8-563d-11ee-a07c-00ffb41493a0' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFkbWluIiwiZW1haWwiOiJhZG1pbkBnbWFpbC5jb20iLCJleHAiOjE2OTUwNjAzOTN9.pF1LWtqmM_KoUnePBVib-h0nLyvpr2qKQRlnEjV35cg'
```
#### Update Product
```bash
curl --location --request PUT 'localhost:3030/product/ecb9a293-5609-11ee-a0aa-00ffb41493a0' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFiZHVyIiwiZW1haWwiOiJkaXRvYWRyaWVsQGdtYWlsLmNvbSIsImV4cCI6MTY5NTAzNjAwM30.aEKqb50s8thxaAr94vVgrwnYwjliGCsrdOInMAOfwOQ' \
--header 'Content-Type: application/json' \
--data-raw '{
    "price": 1321351
}'
```
#### Delete Product
```bash
curl --location --request DELETE 'localhost:3030/product/ecb9a293-5609-11ee-a0aa-00ffb41493a0' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFiZHVyIiwiZW1haWwiOiJkaXRvYWRyaWVsQGdtYWlsLmNvbSIsImV4cCI6MTY5NTAzNjAwM30.aEKqb50s8thxaAr94vVgrwnYwjliGCsrdOInMAOfwOQ'
```
#### Add To Cart
```bash
curl --location --request POST 'localhost:3030/cart/items' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFkbWluIiwiZW1haWwiOiJhZG1pbkBnbWFpbC5jb20iLCJleHAiOjE2OTUwNjAzOTN9.pF1LWtqmM_KoUnePBVib-h0nLyvpr2qKQRlnEjV35cg' \
--header 'Content-Type: application/json' \
--data-raw '[
    {
        "product_id": "5b83ae19-563d-11ee-a07c-00ffb41493a0",
        "quantity": 1
    }
]'
```
#### Get Cart
```bash
curl --location --request GET 'localhost:3030/cart/items' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFkbWluIiwiZW1haWwiOiJhZG1pbkBnbWFpbC5jb20iLCJleHAiOjE2OTUwNjAzOTN9.pF1LWtqmM_KoUnePBVib-h0nLyvpr2qKQRlnEjV35cg'
```
#### Delete Cart
```bash
curl --location --request DELETE 'localhost:3030/cart/items' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFkbWluIiwiZW1haWwiOiJhZG1pbkBnbWFpbC5jb20iLCJleHAiOjE2OTUwNjAzOTN9.pF1LWtqmM_KoUnePBVib-h0nLyvpr2qKQRlnEjV35cg' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cart_id": ["c98d1f3d-563f-11ee-86ca-00ffb41493a0"]
}'
```
#### Checkout Cart
```bash
curl --location --request POST 'localhost:3030/cart/items/checkout' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFkbWluIiwiZW1haWwiOiJhZG1pbkBnbWFpbC5jb20iLCJleHAiOjE2OTUwNjAzOTN9.pF1LWtqmM_KoUnePBVib-h0nLyvpr2qKQRlnEjV35cg' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cart_id": ["c98d1f3d-563f-11ee-86ca-00ffb41493a0"]
}
'
```
### Installation guide
#### 1. install go version 1.19
```bash
# Please Read This Link Installation Guide of GO

# Link Download -> https://go.dev/dl/
# Link Install -> https://go.dev/doc/install

```

#### 2. Run the application
```bash
# run command :
git clone https://Dito_Adriel@bitbucket.org/Bhenedicto_Adriel/dito-rgb-golang-test.git

# install dependency
go mod tidy

# setup env
DB_DRIVER=postgres
DB_USERNAME=        #change to your db username
DB_PASSWORD=        #change to your db password
DB_HOST=            #change to your db host
DB_PORT=            #change to your db port 
DB_DATABASE=        #change to your db name 
DB_URL=             #postgres://{DB_USERNAME}:{DB_PASSWORD}@{DB_HOST}:{DB_PORT}/{DB_DATABASE}?sslmode=disable

KEY=                #change to your key
EXPIRED=            #change to your expiration time

LOGS_SERVICE=       #change to your logs service url
# Run App
make start

# Migrate db
make migrate-up //this for up migrations
make migrate-down //this for down migrations
```
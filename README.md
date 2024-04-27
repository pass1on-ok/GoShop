# Online Store
```
Authors and his IDs:
Oskenbay Kumar, 22B030495
Kenesbek Asylmurat, 22B030376
Murat Dias, 22B030564
Mirpulatov Rozimurat, 22B031185
```

Application where users can view products, add them to cart, place orders and view order history.

## API structure:
```
GET /users
GET /users/{id}
POST /users
PUT /users/{id}
DELETE /users/{id}

GET /categories
GET /categories/{id}
POST /categories
PUT /categories/{id}
DELETE /categories/{id}

GET /products
GET /products/{id}
POST /products
PUT /products/{id}
DELETE /products/{id}

GET /products/{id}/reviews
POST /products/{id}/reviews

GET /cart/{product_id}
POST /cart/{product_id}
DELETE /cart/{product_id}

GET /orders/{user_id}
POST /orders/{user_id}

GET /shipping
POST /shipping
PUT /shipping/{id}
DELETE /shipping/{id}

GET /payment-methods
POST /payment-methods
PUT /payment-methods/{id}
DELETE /payment-methods/{id}

GET /order-status
POST /order-status
PUT /order-status/{id}
DELETE /order-status/{id}
```

## DB Structure:
```
Table user {
ID bigserial primary key
Name varchar
Surname varchar
E-mail varchar
Password_hash varchar
Delivery_address text
History_of_orders jsonb
User_role varchar
}

Table category_product {
ID bigserial primary key
Name varchar
Description text
Link varchar
}

Table product {
ID bigserial primary key
Name varchar
Price numeric
Description text
Quantity_in_stock integer
ImagePath text
}
Table product_reviews {
ID bigserial primary key
Product_ID integer ref product(id)
ID_user integer
Review_text text
Product_rating integer
}

//Many to many
Table cart{
ID bigserial primary key 
User_ID integer ref user(id)
Product_ID integer ref product(id)
}

Table order {
ID bigserial primary key
User_ID integer ref user(id)
Products_and_their_quantity text
Order_price numeric
Order_date timestamp
}

Table shipping_information{
ID bigserial primary key
delivery_method varchar
Cost_of_delivery numeric
Delivery_terms text
}

Table payment_methods {
ID bigserial primary key
Name varchar
Fees numeric
}

Table order_status{
ID bigserial primary key
Status varchar
}
```

# Online Store
# Project description:
Application to create an online store where users can view products, add them to cart, place orders and view order history.

## API structure:
```
GET /products: Get a list of all products
GET /products/{id}: Get information about a product by its id
POST /cart/add: Add product to cart
GET /cart: Get cart contents
POST /order: Place an order
GET /orders: Get a list of all user orders
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

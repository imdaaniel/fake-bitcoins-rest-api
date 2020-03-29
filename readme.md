# 📈 A RESTful API to buy/sell bitcoins 💸🌎
# Get Started

## Description
This is an **representational** application to buy and sell bitcoins built with [Go](https://golang.org/) (it means that you will not buy or sell real bitcoins here!).

## Running
To run this application, you will need:

 1. [Install Go](https://golang.org/dl/)
 2. Crete a database
 3. Set the name of created database in _.env_ > *DB_NAME* (the default is _bitcoins-exchange_)
 4. Execute the main.go file

# Routes

 - **Remember that all parameters should be passed in JSON format.**
 - You can get the postman collection of this API [here](https://www.getpostman.com/collections/cac75a963200d813b55e)

## GET / (index)

A simple welcome message. You can use this route to test if the application is correctly running.

_Success response example_

_HTTP 200 OK_

![](https://imgur.com/7t0Eg5J.png)

## POST /users

Creation of an user. The expected parameters are:

 - name (_String_) - Ex.: Elon Musk
 - email (_String_) - Ex.: elonmusk@spacex.com
 - password (_String_) - Ex.: Rocket2moon@123
 - dateofbirth (_String_) - Ex.: 1971-06-28

_Success response example_

_HTTP 201 Created_

![](https://imgur.com/ws3nUWx.png)

## GET /users

List existing users.

_Success response example_

_HTTP 200 OK_

![](https://imgur.com/Dhuw5Fe.png)

## POST /login

There are the login. The expected parameters are:

 - email (_String_) - Ex.: elonmusk@spacex.com
 - password (_String_) - Ex.: Rocket2moon@123

In success case, you will get a JWT that will be used in the next route.

_Success response example_

_HTTP 200 OK_

![](https://imgur.com/5BCmTLi.png)

## POST /orders

Crete an order. The expected parameters are:

 - Bearer token (JWT obtained at Login. This parameter should be at HTTP Header (key _Authorization_)
 - author_id (_int_) - Ex.: 3
 - amount (_int_) - Ex.: 0.06
 - action (_String_) - "buy" or "sell"
 - date (_String_) - Ex.: "2020-03-29"

_Success response example_

_HTTP 201 Created_

![](https://imgur.com/XLaWElY.png)

## GET /orders/user/:id

List all orders of the user. The expected parameter are:

 - id - Ex.: 3

_Success response example_

_HTTP 200 OK_

![](https://imgur.com/gpA0kux.png)


## GET /orders/date/:date

List all orders realized in the date. The expected parameter are:

 - date - Ex.: 2020-03-29

_Success response example_

_HTTP 200 OK_

![](https://imgur.com/3N1wDOb.png)

All routes have been documented, so it looks like

![](https://media.giphy.com/media/upg0i1m4DLe5q/giphy.gif)

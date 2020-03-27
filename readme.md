# 📈 A RESTful API to buy/sell bitcoins 💸🌎
## [Get started]

### Description
This is an **representational** application to buy and sell bitcoins built with [Go](https://golang.org/) (it means that you will not buy or sell real bitcoins here!).

### Running
To run this application, you will need:

 1. [Install Go](https://golang.org/dl/)
 2. Crete a database
 3. Set the name of created database in _.env_ > *DB_NAME* (the default is _bitcoins-exchange_)
 4. Execute the main.go file

# Routes

**Remember that all parameters should be passed in JSON format.**

## GET / (index)

A simple welcome message. You can use this route to test if the application is correctly running.

## POST /users

Creation of an user. The expected parameters are:

 - name (_String_) - Ex.: Elon Musk
 - email (_String_) - Ex.: elonmusk@spacex.com
 - password (_String_) - Ex.: Rocket2moon@123
 - dateofbirth (_String_) - Ex.: 1971-06-28

_Success response example_

_HTTP 201 Created_

![](https://i.imgur.com/TrUqaCd.png)

## GET /users

List existing users.

_Success response example_

_HTTP 200 OK_

![](https://i.imgur.com/QeqBmcV.png)

## POST /login

There are the login. The expected parameters are:

 - email (_String_) - Ex.: elonmusk@spacex.com
 - password (_String_) - Ex.: Rocket2moon@123


_Success response example_

_HTTP 200 OK_

![](https://imgur.com/Mwpr3mc.png)

## POST /orders

Crete an order. The expected parameters are:

 - Bearer token (Header _Authorization_)
 - author_id (_int_) - Ex.: 3
 - amount (_int_) - Ex.: 0.05
 - action (_String_) - "buy" or "sell"

_Success response example_

_HTTP 201 Created_

![](https://imgur.com/YvyM5tm.png)

## GET /orders/user/:id

List all orders of the user. The expected parameter are:

 - id - Ex.: 3

_Success response example_

_HTTP 200 OK_

![](https://imgur.com/8f78PkQ.png)


## GET /orders/date/:date

List all orders realized in the date. The expected parameter are:

 - date - Ex.: 2020-03-27

_Success response example_

_HTTP 200 OK_

![](https://imgur.com/HfCunOY.png)

# Disc Golf Store

This website was created by Luke and August as a final project for CSC-391 at Wisconsin Lutheran College. Our disc golf store was created using Go, SQLite, and Vue. Go was used to create and run all endpoints. Vue was used for all frontend developement and the consumption of endpoints. SQLite was used as our DBMS which allowed for a simple implementation of our data which was used in our endpoints. 

## Running the Website
In order to run the website, you must type the following commands into your terminal from the `Disc-Golf-Store>` directory.
```
cd frontend/disc-golf
```
```
npm install
```
```
npm run dev
```

## Running the API
In order for your website to be data driven, you must run the API that accesses our database `discgolf.db`. To do this, type the following commands into your terminal from the `Disc-Golf-Store>` directory.
```
cd backend/main-api
```
```
go run .
```
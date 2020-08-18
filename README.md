# cache
--------------------------------------------
Cache is the microservice deployed on kubernetes. It stores data in memory and has a persistent backup from data base.
- The microservice provides functionality of get and post to cache memory. The memory make sure that the same is persisted in mysql database.
- On redeployment of service the rabbitmq fetches all data from db and inserts it into cache memory.

API exposed -
- GET - /cache
- POST - /cache

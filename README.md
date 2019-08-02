# ArkLog
 
Ark log is centralized log stored in mongoDB


```
docker run -d --name local-mongo -p 27017:27017 -v /my/mongo:/etc/mongo \
    -e MONGO_INITDB_ROOT_USERNAME=root \
    -e MONGO_INITDB_ROOT_PASSWORD=root \
    mongo
```
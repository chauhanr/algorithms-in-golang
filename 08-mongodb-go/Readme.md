# Mongo DB 

## Installing the mongo db using docker 

1. pull the docker image `docker pull mongo:4.0.4`
2. To run the mongo image run the command  `docker run -d -p 27017-27019:27017-27019 --name mongodb
   mongo:4.0.4` 

## Interacting with the mongo db - command line 

We need to get into the docker image command line: 
* `docker exec -it mongodb bash` - this will allow us into the container. 
* next we need to run the mongo command - `mongo` this will open an interactive shell. 
* Some basic commands to get going 
	* `show dbs` - shows all the databases that mongo instance has. 
	* `use developers` - this does not create the database instead just defines it. It will
	  create the database when we add data to it in the next step 
	* `db.people.save({ firstname: "Nic", lastname: "Ryan"})`

From the example above we not have a developers database with people table and one document in the
people collection. 

## Setting the code up to interact with Mongo DB 

First we need to use a mongo session / driver to interact with the NoSQL database. The user
controller in the code base which has been defined as a struct will be the holder for that session. 

```
type UserController struct{
    session *mgo.Session
}
```

On the main program we need to get the mongo session 

```
func getSession() *mgo.Session{
   s, err := mgo.Dial("mongo://localhost") 
   if err != nil {
      panic(err) 
   }
   return s
}

```


 

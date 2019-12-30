# Mongo Databases 

## Database commands 

**Create or Use a database** 
when you issues the `use` command the  database is either created or you can switch to the database
that is already present.

```
$ use <db-name>    # switches you to the database 
```

**Show command** 
show command can be used to either view the dbs that we have or the collections in a db. 

* `show dbs` - lists all the databases in the mongo instance 
* `show collections` - lists all the collections in the current db. 

**db command** 
The `db` command is the command that contains all the operations that can be performed on the
mongodb instance. 

* `db`  - shows the current database that you are in or using 
* `db.<collection-name>.insert({"name": "Chauhan"})`  - this is used to insert a document to the
  collection with the name given. 
* `db.<collection>.find()` - view all the documents in the colection at hand. We can add filter
  criteria in find() to select only certian documents in the collection. 
* `db.dropDatabase()`  


## Collections 

Collections in mongo can be created implicity or explicity. 

**Implicit creation**

`db.cats.insert({"firstname": "coco"})`  - here we have explicity created cats if it did not exist
earlier. 

**Explicit creation** 

`db.createCollection(<name>, {<optinal options>})`

the options that can be set are: 

* **capped** of type bool and mentions that the collection is capped in size. 
* **size** - sets the cap in number of bytes it is an integer
* **max** - this determines the maximum  number of documents allowed in the capped collection. 

example -  `db.createCollection("cps" , { capped: true, size: 65536, max: 1000000})`

We can also drop the collection by the using the following command: 

`db.<collection-name>.drop()` 


## Documents 

Documents in MongoDB is a plain json inorder to insert these documents we need to use the insert
command. 

`db.<collection-name>.insert([{document}, {document}, {document}])`

```
db.crayons.insert([
                    {
                        "hex": "#EFDECD", 
                        "name": "Almond", 
                        "rgb": "(239, 222, 205)"
                    }, 
                    {
                        "hex": "#CD9575", 
                        "name": "Antique Brass", 
                        "rgb": "(205, 149, 117)"
                    }, 
                    {
                        "hex": "#FDD9B5", 
                        "name": "Apricot", 
                        "rgb": "(253, 217, 181)"
                    }, 
                    {
                        "hex": "#78DBE2", 
                        "name": "Aquamarine", 
                        "rgb": "(120, 219, 226)"
                    }, 
                    {
                        "hex": "#87A96B", 
                        "name": "Asparagus", 
                        "rgb": "(135, 169, 107)"
                    }, 
                    {
                        "hex": "#FFA474", 
                        "name": "Atomic Tangerine", 
                        "rgb": "(255, 164, 116)"
                    }, 
                    {
                        "hex": "#FAE7B5", 
                        "name": "Banana Mania", 
                        "rgb": "(250, 231, 181)"
                    }
                ])

```



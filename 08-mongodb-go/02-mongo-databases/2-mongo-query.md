# MongoDB query 

## find() 

setup the following stores 
```
use store
db
show dbs
db.customers.insert([{"role":"double-zero","name": "Bond","age": 32},{"role":"citizen","name": "Moneypenny","age":32},{"role":"citizen","name": "Q","age":67},{"role":"citizen","name": "M","age":57},{"role":"citizen","name": "Dr. No","age":52}])

```

to get documents under store collection use the following: 
* `db.customers.find()` - gets all document as we have not specified any filter criteria. 
* `db.customers.findOne()` - this returns the first document it can find in this collection. 

**Adding some filters**

* `db.customers.find({"name": "Bond"})` - this will filter result to get the customers with name
  bond. 

**And operation** 

* `db.customers.find({$and: [{role: "citizen"}, {age: {$gt: 20}}]})`
* `db.customers.find({$and: [{role: "citizen"}, {age: {$lt: 30}}]})` 

**Or operation** 

* `db.customers.find({$or: [{role: "citizen"}, {age: {$gt: 20}}]})`


**and and or operations together** 

```
db.customers.find({$or:[
{ $and : [ { role : "citizen" }, { age : 32 } ] },
{ $and : [ { role : "citizen" }, { age : 67 } ] }
]})
```

**Regex** 

The following will look for any name which starts with M. 

```
db.customers.find({name: {regex: ^M}}) 
```
[Regex cheat sheet](regex.pdf) 

**Operators**  

* eqaulity -  `db.customers.find({name: "Bond"})`
* not equality - `db.customers.find({"age": {$ne: 20}})`
* greater than - `db.customers.find({age: {$gt: 32}})`
* greater than equal to - `db.customers.find({age: {$gte: 32}})`
* lesser than - `db.customers.find({age: {$lt: 32}})`
* lesser than equals - `db.customers.find({age: {$lte: 32}})`


## Update and Save 

* The update operation will update a record that already exists in the collection 
* The save operation overwrites the existing record in the collection. 

The syntax for update - `db.<collection>.update(<selection-criteria>,<data to update>, <optional
options>)`

Run the insert command in a database called store. 

```
db.customers.insert([{  "role" : "double-zero", "name" : "Bond", "age" : 32 }, {  "role" : "double-zero", "name" : "Bond", "age" : 32 }, { "role" : "citizen", "name" : "Q", "age" : 67 }, { "role" : "citizen", "name" : "M", "age" : 57 }, { "role" : "citizen", "name" : "Dr. No", "age" : 52 }])
```

Now we can run the following queries to update the records. 

* `db.customers.update({name:"Moneypenny"},{$set:{role:"double-zero"}})`
* `db.customers.update({name:"Moneypenny"},{$set:{role:"citizen", name: "Miss Moneypenny"}})`
* `db.customers.update({age:{$gt:35}},{$set:{role:"double-zero"}})` -- this will not work as we have
  to mention that this query will update all records that are selected by the criteria
* `db.customers.update({age:{$gt:35}},{$set:{role:"double-zero"}}, {multi:true})` -- the
  `{multi:true}` option allows for multiple records to be updated. 
* `db.customers.update({},{$set:{role:"citizen"}}, {multi:true})` -- as you can see a no selection
  criteria means match all records. Therefore this statement will update all records. 

**Save** 

* `db.customers.save({"role":"villain","name":"Jaws","age":43})`  -- as a record does not exist a
  new record will be created
* `db.customers.save({"_id": ObjectId("5e0aacbf25db187e28cad44e"), "name": "Dr. No", "role": "villan", "age": 77})`  


## Remove 

This will delete records in the collection. 

`db.<collection-name>.remove(<selection-criteria>)` 
examples: 
* `db.customers.remove({"role": "villain"})`
* `db.customers.remove({"role": "double-zero"})`

Here we do not have to mention multiple match true as the remove operation removes all that it
matches. However if you want to limit the records deleted we can mention a number 

* `db.customers.remove({role: "citizen"}, 1)`
* `db.customers.remove({})` -- removes all the records. 


## Projection 

This about only selecting only certian fields in the collection. 

The projections syntax is - `db.<collection>.find(<selection-criteria>, <list of fields with
toggle 0, 1>)`

* `db.customers.find({},{_id:0, name:1,})`
* `db.customers.find({},{_id:0, name:1,age: 1, })`
* `db.customers.find({age: {"$gt:35"}},{_id:0, name:1,age: 1, })`




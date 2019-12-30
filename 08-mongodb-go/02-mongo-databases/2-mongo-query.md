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

## Operators 

* eqaulity -  `db.customers.find({name: "Bond"})`
* not equality - `db.customers.find({"age": {$ne: 20}})`
* greater than - `db.customers.find({age: {$gt: 32}})`
* greater than equal to - `db.customers.find({age: {$gte: 32}})`
* lesser than - `db.customers.find({age: {$lt: 32}})`
* lesser than equals - `db.customers.find({age: {$lte: 32}})`





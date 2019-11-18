# Golang Data strucutres and Algorithms

## Strucutral Design Patterns
These are patterns GoF patterns that are used often in the implementation of software.

### Adapter
There are 4 palyers in the adapter pattern
* Client - calling the adapter in order to invoke a functionality
* Adapter - is the class called to convert the call to the incomaptible functionality
* Adaptee - is the incompatible functionity that is converted using the adapter
* Target - is the interface that the client calls and invokes the method on the adapter.

### Bridge Pattern
This is a pattern that is built to promote composition over inheritence. The scenarios where bridge
pattern is relevant are:

1. Runtime binding the application
2. Mapping  orthogonal class hierarchies
3. platform independence implementation are scenarios

```
type IConture interface{
 drawContour(x[5] float32, y[5] float32)
 resizeByFactor(factor int)
}

type DrawContour struct{
   x[5] float32
   y[5] float32
   shape DrawShape
   factor int
}

func (contour DrawContour) drawContour(x[5] float32, y[5] float32) {
    fmt.Printf("..")
    counter.shape.drawShape(contour.x, contour.y)  
}

func (contour DrawContour) resizeByFactor(factor int){
   contour.factor = factor  
}
```

### Composite Pattern
Composite is a group of similar objects in a single object. The objects are stored in tree format.

The examples where the composite pattern is used is UI tree layout, directories and files etc.

```
type IComposite interface {
   perform()
}

type Leaflet struct{
  name string
}

func (leaf *Leaflet) perform(){
  ...
}

type Branch struct {
  leafs    []Leaflet
  name     string
  branches []Branch
}

func (branch *Branch) perform(){

   for _, leaf := range branch.leafs{
     leaf.perform()
   }
   for _, branch := range branch.branches{
     branch.perform()
   }
}

func (branch *Branch) add (leaf Leaflet) {
    branch.leafs = append(branch.leafs, leaf)
}
```

[Next](2-algorithm-reps.md)

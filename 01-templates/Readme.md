# Template 

Templates section of this project helps understand how the templates for text and html are present. 

## Understand the text/template 
golang template package allow for data driven templates to generate text based output. 
Templates are executed against data structures and the annotations in the template refer to elements
in data structures (fields in a structure or keys in a map). 

The structure of a template is very simple. It has two elements: 
1. Actions - data evaluations or control structures are denoted by `{{ and }}
2. Text - anything that is not within an Action is copied to the output as is. 

There are some formating allowed in the action section: 
```
{{- 23}} < {{45 -}}
``` 
will lead to the output  `23<45` as the minus sign after/before the action delimiters indicate that
the traling or leading spaces must be trimmed. However `{{-3}}` would result is -3 being displayed. 


## Actions 

**Data Evaluations** 
In order to refer to a field is a struct we can use `{{ .FieldName}}` where FieldName is a name of
field in a struct. 

**Conditions** 

We can use a `if else` construct in conditions. 

```
# if the pipeline is empty then T1 is executed otherwise T0 is executed. 
{{ if pipeline }} 
    T0
{{ else }} 
    T1 
{{ end }}
``` 

There is a support for loops as well 

```
# for this to work the value of pipeline has to be an array, map, slice.
# if the pipeline is empty the nothing is done else T1 is executed. 
{{range pipeline }} 
    T1 
{{ end }} 
```

**Templates** 

```
{{template "name"}} 
# teomplate with specified name is executed. 

{{template "name" pipeline}} 

# template with name specified is executed with dot set to value of the pipeline   

{{with pipeline}} T1 {{else}} T2 {{end}} 

```





## Arguments 
* Any of go's untyped constants, bool, int, float string etc. 
* variables are denoted by the $ sign  e.g. `$value` etc. 
* name of field in a struct can be denoted with `.` operator e.g. `.FieldName` or
  `$value.Field1.Field2` 
* Method which is run as `.Method` this will run on the structure passed as `.Method()` this too can
  be chanined. But the mehtod must return not more than one return type or almost 2 second of which
is a error. Other wise the execution of hte template stops. 


## Pipeline 
Pipeline is a possible chained sequence of commands. A `command` is simple value i.e. an argument,
or function method call with multiple arguments. 

* Argument - the result after the evaulation of argument is used as piepline output.
* .Method[Argument] - method that are at the end of the chain, unlike the ones in the middle of the
  chain can take arguments. 
* functionName[Argument] 


## Examples to understand 

A constant string 
```
{{ "\"Output\"" }} 
{{ `"output"` }} 

``` 

Funtion calls 
```
{{printf "%q" "output"}}

{{"output" | printf "%q"}} - function call the input of which comes from previous command.  

```

With examples 
```
{{with "output"}}{{printf "%q" .}}{{end}}

{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
```

## Nested Templates 

```
{{define "T1"}}ONE{{end}} 
{{define "T2"}}TWO{{end}} 
{{define "T3"}}{{template "T1"}}  {{template "T2"}}{{end}} 
{{template "T3"}} 

```

**Blocks** 
blocks are short hand for defining a template and then executing it in its place. 

```
{{block "name" pipeline}} T1 {{end}} 

# this is equal to 
{{define "name"}} T1 {{end}} 
# and executing it in place 
{{template "name" pipeline}}

This is generally used to define base templates that are then customized and redefined later.   
 
```



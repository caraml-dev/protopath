# Protopath
Implementation of JsonPath syntax that enable value extraction from protobuf message. In the future it can be used to extract any value from any type. This library use Antlr4 as parser.

## Getting Started

- [Go 1.18 or above](https://go.dev/doc/install)
- [Antlr 4.10.1](https://www.antlr.org/download.html). Download antlr4 binary and put the jar into this location /usr/local/lib/antlr-4.10.1-complete.jar


## Capability Matrix
This protopath only support dot notation, we don't have any plan to support bracket notation
Below are the supported expression in this library

| Expression                                | Description                                                   |
|-------------------------------------------|---------------------------------------------------------------|
| $.property                                | Selects the specified property in a parent object.            |
| $.properties[n]                           | Selects the n-th element from an array. Indexes are 0-based.  |
| $.properties[index1, index2...]           | Selects array elements with the specified indexes, return list|
| $.properties[start:end] or $.properties[start:] or $.properties[:end] or $.properties[:]                  | Selects array elements from the start index and up to, but not including, end index. If end is omitted, selects all elements from start until the end of the array. Returns a list|
| $.properties[*]                           | Select all objects in properties the result will be flatenned list  |
| $.properties[@.length-n]                  | Select the length-n th elements from array                    |
| $.properties[?(expression)]               | Filter expression. Select all elements in array that match the specified filter|


Below are the supported filter expression
| Expression          |   Description                                           |
|---------------------|---------------------------------------------------------|
| ==                  | Equal                                                   |
| !=                  | Not Equal                                               |
| >                   | Greater than                                            |
| >=                  | Greater equal than                                      |
| <                   | Less than                                               |
| <=                  | Less equal than                                         |
| &&                  | Logical and                                             |
| ||                  | Logical or                                              |

Currently this library haven't supported regex expression filtering

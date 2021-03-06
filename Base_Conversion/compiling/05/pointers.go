package main

import "fmt"

// STRUCTS0 OMIT

type person struct {
  name string
  age int
  address string
}

type node struct {
  info person
  back, next *node
}

// STRUCTS0 OMIT

// ADDNODE1 OMIT

func add(pointer *node, data person) *node {
  var cell node
  cell.info = data
  cell.back, cell.next = pointer, nil
  newer := &cell
  if pointer == nil {
    pointer = newer
  } else {
    pointer.next = newer
    pointer = newer
  }
  return pointer
}

// ADDNODE1 OMIT

// DUMP OMIT

func dump(pointer *node) {
  if pointer != nil {
    fmt.Printf("Content: %p \n", *pointer)
    fmt.Println("--------------------")
    dump(pointer.back)
  }
}

// DUMP OMIT

// APPLICATION OMIT

func main() {
    var info person
    info.name, info.age, info.address = "Germán Alberto xxxxxxx xxxxx", 30, "xxxxxxxxxxxxxxx xxx" 
    last := add(nil, info)
    info.name, info.age, info.address = "Juan Francisco Giménez Silva", 20, "xxx xxxxx xxx" 
    last = add(last, info)
    info.name, info.age, info.address = "Ernestina yyyyyyy yyyyyy", 1, "rrrrrrr nnnnnnn y9y"
    last = add(last, info)
    info.name, info.age, info.address = "María Jimena gggggg", 22, "fffffff vvvvvvv 355"
    last = add(last, info)

    dump(last)

}

// FINISH OMIT

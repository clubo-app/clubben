package id

// Used to create, manage, and parse unique IDs. This code creates stripe-like IDs,
// (e.g., xx_000000000000000000000000000). It allows for a prefix and a 27
// character ksuid separated by an underscore. The prefix makes it easy
// to visibly identify what the ID is for.
//
// Original Author: Kirk Morales, https://gist.github.com/knation/3541b4da1c5274eaf03ceafa6985bd0a

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

// Number of characters before the underscore in an ID
const prefixLength = 2

// Define enum for ID types
type TypeId int32

const (
	Undefined TypeId = iota
	User
	Company
)

// `idData` contains data for each TypeId
var idData = map[TypeId]([]string){
	Undefined: {"", "undefined"},
	User:      {"us", "user"},
	Company:   {"co", "company"},
}

// Used as a reverse lookup of TypeId by prefix
var (
	prefixLookup     map[string]TypeId
	prefixLookupInit = false
)

// Populates `prefixLookup`
func createPrefixLookup() {
	prefixLookup = make(map[string]TypeId)
	for key, val := range idData {
		prefixLookup[val[0]] = key
	}
}

// `String()` method for `TypeId`
func (typeId TypeId) String() string {
	return idData[typeId][1]
}

// Creates an ID string for the given object type
func New(typeId TypeId) string {
	if typeId == Undefined {
		return ""
	}

	id := ksuid.New().String()
	prefix := idData[typeId][0]

	return fmt.Sprintf("%s_%s", prefix, id)
}

// Gets the type for the given ID string
func GetType(id string) TypeId {
	if len(id) != (prefixLength+28) || id[prefixLength:prefixLength+1] != "_" {
		return Undefined
	}

	if !prefixLookupInit {
		createPrefixLookup()
		prefixLookupInit = true
	}

	typeId := prefixLookup[id[0:2]]
	if typeId != 0 {
		return typeId
	} else {
		return Undefined
	}
}

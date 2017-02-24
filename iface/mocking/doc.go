// This demo show how to mock external package's structs
// in case of lack of interfaces.
//
// Basic concept is - write an interface by hand, completely or
// partially (depending on use case) declaring same methods as
// target struct. Then use this interface in code anywhere where needed.
// Target struct implicitly satisfies our interface and then
// we could generate or write mocks for this interface and use it in our tests.
//
// Basically we could wrap target struct in our custom struct and use this in our code,
// and then use this struct in our test, depends on use case.
package main
golang-euler
============

A collection of solutions to [Project Euler Problems](https://projecteuler.net/problems) written in the Go language.

Use them to see Go in action, or get a gist of programming fundamentals in general.

Feel more than free to add your own solutions! 

*** Conventions

Since this is meant to be a collection of solutions, I want finding or running a specific problem
to be as simple as possible. 

So:

1. Filenames will consistently be in the format XXX.go, where XXX is a three digit string of the problem number (with leading zeros 
as neccessary). This is so Github will organize the solutions numerically
2. Each file has a function with the name PX(), where X is the number of the solution. X should NOT have leading zeros 
(ex. P2() and P34() are the valid solution functions for problems 2 and 34)
3. All tests go under solutions_test.go
4. Test functions have the name TestX, with X being the problem number (in the style of #2). Each function should have
a test case for the solution function, but is encouraged to also have tests for any other utility functions used when solving
the problem.

Feel free to discuss these or suggest your own guidelines.

*** Running

The solutions are meant to be run through the command "go test". 
Go will automatically find the tester and run each test method.

To run an individual solution yourself, I reccomend going to [Go Playground](https://play.golang.org/),
adding the code of the solution to there, as well as a main function where you can Println whatever you're looking to see.

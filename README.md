# Coding Interview Patterns - ByteByteGo - Golang
Hello there! Did you also find yourself studying off the book ["Coding Interview Patterns"](https://github.com/ByteByteGoHq/coding-interview-patterns) Alex Xu \& Shaun Gunawardane? Did you also find yourself wondering why those problems are not presented in Golang? Did you also look for any Repo implementing that? Well.. stop looking. This is it!

You'll find here the setup to run your own solutions for each problem presented in the book. Of course, check out [their repo](https://github.com/ByteByteGoHq/coding-interview-patterns), where you can find some of those solutions in Go. The structure of this repo follows the chapter structure of the book:

1. Two Pointers
2. Hash Maps and Sets
3. Linked List
4. Fast and Slow Pointers
5. Sliding Windows
6. Binary Search
7. Stacks
8. Heaps
9. Intervals
10. Prefix Sums
11. Trees
12. Tries
13. Graphs
14. Backtracking
15. Dynamic Programming
16. Greedy
17. Sort and Search
18. Bit Manipulation
19. Math and Geometry

## How does this work?
It's very simple: just clone the repo locally, implement your solution inside the `YourSolution()` func, then try by to run it (e.g. with the Pair-Sum Sorted exercise):
```bash
# navigate to folder
cd 01-TwoPointers/PairSumSorted

# run test
go run pair-sum-sorted.go
```
Each file run will check your solution against the tests mentioned in the **Test Cases** paragraph of the book (the table that includes Input, Expected Output and Description).

## Contributions
Feel free to open a PR if you want to contribute. In case you might want to post your solution, this is not the best repo to do so: please do it on the [Official Repo - Go section](https://github.com/ByteByteGoHq/coding-interview-patterns/tree/main/go)



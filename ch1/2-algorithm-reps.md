# Algorithms

## Complexity and performance

Alogrithms are compared based on their processing time and resource consumption. The complxity of an algorithm is represented by the Big O notation.

There are several Big O notations that an algorithm can be in terms of time:

1. Linear time O(n) - this is the complexity of algorithms like linear search, traversing and finding min and max values in an array. Data structures likes Array List and queues are ones that generally have this algorithmic complexity.
2. Logrithmic complexity O(log n) can be seem for binary search, in tree data structures.
3. Square complexity O(n^2) is seen with bubble sort, insertion sort.
4. Big Omega and Big Theta are notations for lower and upper bounds of particular algorithm.

## Brute Force Algorithms

These algorithms are based on searching or working the entire set of possible solutions without any elimination based on different criterias. Therefore these algorithms are less performant.

## Divide and Conquer Algorithms

The divide and conquer algorithms break the main problem into smaller sub problems. The smaller sub problems are broken down till it is a known trivial problem. At this point the sub problems are sovled and the results are combines recursively to get to the final solution.Examples are:
* binary search
* quick sort
* fast Fourier transform.
* merge sort

Generally these algorithms are memory efficient and performance sometimes due to recurrsion are not that good because of recurrsion however the use of multi processor can allow for individual processes to be dealt with by different processors.

## Backtracking Algorithms

This algorithm attempts to solve a problem by constructing the solution incrementally. Multiple options are evaluated and algorithm chooses to go to the next component of solution through recursion.

Another feature of backtracking is that it finds candidates for solution and rejects a candidate based the basis of feasibiluty and validity. It is better than brute force as it rejects a lot of potential non solutions in an iteration. Examples where backtracking is used:

1. parsing
2. rules engine
3. kanpsack problem
4. combinatorial optimizations.

[Next](../ch3/1-intro.md)

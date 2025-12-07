# Big O Notation

## 1. O(1) - Constant Time
**Explanation:** Execution time does not depend on input size. It always takes a fixed amount of time.

**Real-world examples:**
- Accessing array elements by index
- Getting values from HashMap
- Push/Pop operations on Stack
- Adding/removing elements from the beginning/end of LinkedList

## 2. O(log n) - Logarithmic Time
**Explanation:** Each step reduces the problem size by half (or a fraction). Very efficient for large datasets.

**Real-world examples:**
- Binary Search
- Searching in Binary Search Tree
- Heap operations
- Divide and Conquer algorithms

**Code sample:**
```go
// Binary Search - O(log n)
func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1

    for left <= right {
        mid := left + (right-left)/2

        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1
}
```

## 3. O(n) - Linear Time
**Explanation:** Time increases linearly with input size. Traverses each element once.

**Real-world examples:**
- Traversing arrays/slices
- Linear search
- Calculating sum/average
- Copying arrays

## 4. O(n log n) - Linearithmic Time
**Explanation:** Combines linear and logarithmic complexity. Often appears in efficient divide and conquer algorithms.

**Real-world examples:**
- Merge Sort
- Quick Sort (average case)
- Heap Sort
- Standard library sorting algorithms

## 5. O(n²) - Quadratic Time
**Explanation:** Time increases quadratically with input size. Involves nested loops.

**Real-world examples:**
- Bubble Sort
- Selection Sort
- Insertion Sort
- Checking all pairs in an array

## 6. O(2ⁿ) - Exponential Time
**Explanation:** Time doubles with each addition to the input size. Very inefficient for large inputs.

**Real-world examples:**
- Recursive Fibonacci (without memoization)
- Generating all subsets of a set
- Tower of Hanoi problem

## 7. O(n!) - Factorial Time
**Explanation:** Time grows factorially with input size. Extremely inefficient.

**Real-world examples:**
- Generating all permutations
- Solving the Traveling Salesman Problem (brute force)
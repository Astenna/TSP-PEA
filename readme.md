# Travelling Salesman Problem

This project was developed during my fifth term of Computer Science studies on Wroclaw University of Science and Technology. During this course we implemented many different algorithms solving the travelling salesman problem.This was my first attempt to use Golang instead of C++. It was hard at the beginning but certainly worth it! :)

# Implemented algorithms

- **Precise algorithms:**
    Located at [exact](https://github.com/Astenna/TSP-PEA/tree/master/src/exact) package
    - Brute Force
    - Held-Karp (dynamic programming)
    - Branch and bound
- **Heuristic algorithms:**
    Located at [local](https://github.com/Astenna/TSP-PEA/tree/master/src/local) package
    - Simulated annealing (basic version)
    - List-based Simulated annealing (based on [publication](https://www.hindawi.com/journals/cin/2016/1712630/#B25))
- **Genetic algorithms:**
    Located at [genetic](https://github.com/Astenna/TSP-PEA/tree/master/src/genetic) package
    - Genetic algorithm

# Tests
    The test results of my implementations can be found in raports folder in the root project repository. 

# About implementations
    Helper methods used in each of the algorithms were separated to the [sliceExtensions](https://github.com/Astenna/TSP-PEA/tree/master/src/sliceExtensions) package.
    At the start I tried to implement base-struct that would generalize flow and structure of each algorithm but finally I only used that struct in the implementation of precise algorithms (the base class travelling_salesman_problem.go can be found in the exact package).

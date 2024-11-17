# README

This is a simple program that demonstrates the usage of single process, multi process, and multi thread execution modes to perform CPU intensive tasks. The program allows the user to specify the mode of execution, complexity level, and the number of iterations to run.

## Usage

To run the program, use the following command:

```
go run main.go -m <mode> -c <complexity> -i <iterations>
```

- `<mode>`: Select the mode of execution. Options are `singleprocess` (or `s`), `multiprocess` (or `p`), or `multithread` (or `t`).
- `<complexity>`: Set the complexity level of the task.
- `<iterations>`: Set the number of iterations to run.

### Example

To run the program with multi process mode, complexity level 2, and 5 iterations:

```
go run main.go -m p -c 2 -i 5
```

## Implementation

The program spawns multiple processes or threads based on the selected mode and performs CPU intensive tasks according to the complexity level specified. The execution time for each mode is calculated and displayed at the end.



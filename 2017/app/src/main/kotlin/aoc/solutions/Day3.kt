package aoc.solutions

import aoc.utils.Direction

fun solveDayThree() {
    val input = 361527
    println("Day 3 part 1: ${partOne(input)}")
    println("Day 3 part 2: ${partTwo(input)}")
}

private fun nextDirection(direction: Direction): Direction =
    when (direction) {
        Direction.RIGHT -> Direction.UP
        Direction.UP -> Direction.LEFT
        Direction.LEFT -> Direction.DOWN
        Direction.DOWN -> Direction.RIGHT
    }

private fun partTwo(input: Int): Int {
    val grid = mutableMapOf<Pair<Int, Int>, Int>()
    var x = 0
    var y = 0
    var direction = Direction.RIGHT
    var steps = 1
    var stepCount = 0
    grid[Pair(x, y)] = 1

    while (true) {
        when (direction) {
            Direction.RIGHT -> x++
            Direction.UP -> y++
            Direction.LEFT -> x--
            Direction.DOWN -> y--
        }
        stepCount++

        if (stepCount == steps) {
            stepCount = 0
            direction = nextDirection(direction)
            if (direction == Direction.RIGHT || direction == Direction.LEFT) {
                steps++
            }
        }

        grid[Pair(x, y)] =
            (x - 1..x + 1).sumOf { i ->
                (y - 1..y + 1).sumOf { j ->
                    grid.getOrDefault(Pair(i, j), 0)
                }
            }

        if (grid[Pair(x, y)]!! > input) {
            return grid[Pair(x, y)]!!
        }
    }
}

private fun partOne(input: Int): Int {
    val square = findClosestOddSquare(input)
    val axisDistance = (square - 1) / 2
    val cornerDistance = Math.abs(input - square * square)
    val distanceToAxis = cornerDistance % (square - 1)
    return axisDistance + Math.abs(distanceToAxis - axisDistance)
}

private fun findClosestOddSquare(input: Int): Int {
    var i = 1
    while (i * i < input) {
        i += 2
    }
    return i
}

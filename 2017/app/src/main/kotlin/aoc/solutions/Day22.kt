package aoc.solutions

import aoc.utils.Direction
import aoc.utils.FileLoader
import aoc.utils.Point

fun solveDayTwentyTwo() {
    val input = FileLoader("day22.txt").readLines()
    println("Day 22 part 1: ${partOne(input)}")
    println("Day 22 part 2: ${partTwo(input)}")
}

private fun partTwo(input: List<String>): Int {
    val points = mutableMapOf<Point, Type>()
    input.forEachIndexed { i, row ->
        for (j in row.indices) {
            points[Point(i, j)] = if (row[j] == '#') Type.INFECTED else Type.CLEAN
        }
    }

    // How many bursts cause a node to become infected?
    var currNode = Point(input.size / 2, input[input.size / 2].length / 2)
    var direction = Direction.UP
    var infected = 0
    repeat(10_000_000) {
        when (points[currNode]!!) {
            Type.CLEAN -> {
                direction = turnLeft(direction)
                points[currNode] = Type.WEAKENED
            }

            Type.WEAKENED -> {
                points[currNode] = Type.INFECTED
                infected += 1
            }

            Type.INFECTED -> {
                direction = turnRight(direction)
                points[currNode] = Type.FLAGGED
            }

            Type.FLAGGED -> {
                direction = turnRight(turnRight(direction))
                points[currNode] = Type.CLEAN
            }
        }
        // Move carrier
        currNode = move(currNode, direction)
        // if the point is not in the points, insert it as clean
        points.putIfAbsent(currNode, Type.CLEAN)
    }
    return infected
}

private fun partOne(input: List<String>): Int {
    val points = mutableMapOf<Point, Boolean>()
    input.forEachIndexed { i, row ->
        for (j in row.indices) {
            points[Point(i, j)] = row[j] == '#'
        }
    }

    // How many bursts cause a node to become infected?
    var currNode = Point(input.size / 2, input[input.size / 2].length / 2)
    var direction = Direction.UP
    var infected = 0
    repeat(10_000) {
        if (points[currNode]!!) {
            // If the current point is infected, turn right
            direction = turnRight(direction)
            points[currNode] = false
        } else {
            direction = turnLeft(direction)
            points[currNode] = true
            infected += 1
        }
        // Move carrier
        currNode = move(currNode, direction)
        // if the point is not in the points, insert it as clean
        points.putIfAbsent(currNode, false)
    }

    return infected
}

private enum class Type {
    CLEAN,
    WEAKENED,
    INFECTED,
    FLAGGED,
}

private fun move(
    node: Point,
    direction: Direction,
): Point {
    return when (direction) {
        Direction.LEFT -> Point(node.x, node.y - 1)
        Direction.UP -> Point(node.x - 1, node.y)
        Direction.DOWN -> Point(node.x + 1, node.y)
        Direction.RIGHT -> Point(node.x, node.y + 1)
    }
}

private fun turnLeft(currDir: Direction): Direction {
    return when (currDir) {
        Direction.LEFT -> Direction.DOWN
        Direction.DOWN -> Direction.RIGHT
        Direction.RIGHT -> Direction.UP
        Direction.UP -> Direction.LEFT
    }
}

private fun turnRight(currDir: Direction): Direction {
    return when (currDir) {
        Direction.RIGHT -> Direction.DOWN
        Direction.UP -> Direction.RIGHT
        Direction.LEFT -> Direction.UP
        Direction.DOWN -> Direction.LEFT
    }
}

package aoc.solutions

import aoc.utils.Direction
import aoc.utils.FileLoader
import aoc.utils.Point

fun solveDayTwentyTwo() {
    val input = FileLoader("day22.txt").readLines()
    println("Day 22 part 1: ${partOne(input)}")
}

private fun partOne(input: List<String>): Int {
    val points = mutableMapOf<Point, Boolean>()
    // Add begnning points
    input.forEachIndexed { i, row ->
        row.split("").forEachIndexed { j, col ->
            points[Point(i, j)] = col == "#"
        }
    }

    // How many bursts cause a node to become infected?
    var currNode = Point(input.size / 2, input[input.size / 2].length / 2)
    println(currNode)
    var direction = Direction.UP

    return 0
}

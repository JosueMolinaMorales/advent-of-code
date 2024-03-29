package aoc.solutions

import aoc.utils.FileLoader
import aoc.utils.Point
import kotlin.math.abs

fun solveDayEleven() {
    val input = FileLoader("day11.txt").readFile()
    println("Day 11 part 1: ${partOne(input)}")
    println("Day 11 part 2: ${partTwo(input)}")
}

private fun partOne(input: String): Int {
    var point = Point(0, 0)
    input.split(",").forEach {
        point = movePoint(it, point)
    }

    return calcDist(point)
}

private fun partTwo(input: String): Int {
    var point = Point(0, 0)
    return input.split(",").maxOfOrNull {
        point = movePoint(it, point)
        calcDist(point)
    } ?: 0
}

private fun calcDist(point: Point): Int {
    return (abs(point.y) + abs(point.x) + maxOf(0, abs(point.y) - abs(point.x))) / 2
}

private fun movePoint(
    move: String,
    point: Point,
): Point {
    val newPoint = Point(point.x, point.y)
    when (move) {
        "n" -> newPoint.x -= 2
        "s" -> newPoint.x += 2
        "e" -> newPoint.y += 2
        "w" -> newPoint.y -= 2
        "ne" -> {
            newPoint.x -= 1
            newPoint.y += 1
        }

        "nw" -> {
            newPoint.x -= 1
            newPoint.y -= 1
        }

        "se" -> {
            newPoint.x += 1
            newPoint.y += 1
        }

        "sw" -> {
            newPoint.x += 1
            newPoint.y -= 1
        }

        else -> throw IllegalArgumentException("Not know direction")
    }
    return newPoint
}

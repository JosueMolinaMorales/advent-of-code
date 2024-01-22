package aoc.solutions

import aoc.utils.Direction
import aoc.utils.FileLoader
import aoc.utils.Point

fun solveDayNineteen() {
    val input = FileLoader("day19.txt").readLines()
    val res = solve(input)
    println("Day 19 part 1: ${res.first}")
    println("Day 19 part 2: ${res.second}")
}

private fun solve(input: List<String>): Pair<String, Int> {
    val grid =
        input.map {
            it.split("")
        }

    val curr = Point(0, grid[0].indexOf("|"))
    var direction = Direction.DOWN
    var res = ""
    var steps = 0
    while (true) {
        // Move in the direction
        when (direction) {
            Direction.DOWN -> curr.x += 1
            Direction.UP -> curr.x -= 1
            Direction.LEFT -> curr.y -= 1
            Direction.RIGHT -> curr.y += 1
        }
        steps += 1
        if (grid[curr.x][curr.y].trim().isEmpty()) {
            break
        }

        // Check if the position is a '+'
        if (grid[curr.x][curr.y] == "+") {
            direction =
                if (direction != Direction.RIGHT && curr.y - 1 >= 0 && grid[curr.x][curr.y - 1] == "-") {
                    Direction.LEFT
                } else if (direction != Direction.LEFT && curr.y + 1 < grid[curr.x].size && grid[curr.x][curr.y + 1] == "-") {
                    Direction.RIGHT
                } else if (direction != Direction.DOWN && curr.x - 1 >= 0 && curr.y < grid[curr.x - 1].size && grid[curr.x - 1][curr.y] == "|") {
                    Direction.UP
                } else {
                    Direction.DOWN
                }
        }
        if (grid[curr.x][curr.y].matches("[a-zA-Z]".toRegex())) {
            res += (grid[curr.x][curr.y])
        }
    }
    return Pair(res, steps)
}

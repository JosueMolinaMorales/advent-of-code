package aoc.solutions

import aoc.utils.Point
import java.util.*

fun solveDayFourteen() {
    val input = "xlqgujun"
    println("Day 14 part 1: ${partOne(input)}")
    println("Day 14 part 2: ${partTwo(input)}")
}

private fun partOne(input: String): Int {
    return createGrid(input).sumOf { it.count { it == "#" } }
}

private fun partTwo(input: String): Int {
    val grid = createGrid(input)
    val seen = HashSet<Point>()
    var groups = 0
    for (row in grid.indices) {
        for (col in grid[row].indices) {
            if (grid[row][col] == "." || Point(row, col) in seen) {
                continue
            }
            val group = bfs(grid, Point(row, col))
            seen.addAll(group)
            groups += 1
        }
    }

    return groups
}

private fun bfs(
    grid: List<List<String>>,
    start: Point,
): HashSet<Point> {
    val queue = LinkedList<Point>()
    val visited = HashSet<Point>()

    queue.add(start)
    visited.add(start)

    while (queue.isNotEmpty()) {
        val curr = queue.poll()
        val neighbors =
            listOf(Point(1, 0), Point(-1, 0), Point(0, 1), Point(0, -1)).map { dir ->
                Point(curr.x + dir.x, curr.y + dir.y)
            }.filter {
                it.x >= 0 && it.x < grid.size && it.y >= 0 && it.y < grid[it.x].size && grid[it.x][it.y] == "#"
            }
        for (neighbor in neighbors) {
            if (neighbor in visited) {
                continue
            }
            visited.add(neighbor)
            queue.add(neighbor)
        }
    }

    return visited
}

private fun createGrid(input: String): List<List<String>> {
    val grid = mutableListOf<List<String>>()
    for (i in 0..127) {
        val hash = hexToBinary(knotHash("$input-$i"))
        grid.add(hash.map { if (it == '1') "#" else "." })
    }

    return grid
}

private fun hexToBinary(hex: String): String {
    return hex.uppercase().map { it.toString().toInt(16).toString(2).padStart(4, '0') }.joinToString("")
}

package aoc.solutions

import aoc.utils.FileLoader

fun solveDayFive() {
    val input = FileLoader("day5.txt").readLines()
    println("Day 5 part 1: ${partOne(input)}")
    println("Day 5 part 2: ${partTwo(input)}")
}

private fun partOne(input: List<String>): Int {
    return solve(input) {
        1
    }
}

private fun partTwo(input: List<String>): Int {
    return solve(input) {
        if (it >= 3) {
            -1
        } else {
            1
        }
    }
}

private fun solve(
    input: List<String>,
    offsetFunc: (Int) -> Int,
): Int {
    var i = 0
    var steps = 0
    val nums = input.map { it.toInt() }.toMutableList()

    while (i in nums.indices) {
        val offset = nums[i]
        nums[i] += offsetFunc(offset)
        i += offset
        steps++
    }
    return steps
}

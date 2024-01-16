package aoc.solutions

import aoc.utils.FileLoader

fun solveDayTwo() {
    val input = FileLoader("day2.txt").readFile()

    println("Day 2 part 1: ${partOne(input)}")
    println("Day 2 part 2: ${partTwo(input)}")
}

private fun partOne(input: String): Int {
    var checkSum = 0

    for (line in input.split("\n")) {
        val nums = line.split("\\s+".toRegex()).map { it.toInt() }
        checkSum += (nums.max() - nums.min())
    }
    return checkSum
}

private fun partTwo(input: String): Int {
    var checkSum = 0
    for (line in input.split("\n")) {
        val nums = line.split("\\s+".toRegex()).map { it.toInt() }
        for (i in nums.indices) {
            for (j in nums.indices) {
                if (i == j) {
                    continue
                }
                if (nums[i] % nums[j] == 0) {
                    checkSum += nums[i] / nums[j]
                }
            }
        }
    }

    return checkSum
}

package aoc.solutions

import aoc.utils.FileLoader

fun solveDayTen() {
    val input = FileLoader("day10.txt").readFile()
    println("Day 10 part 1: ${partOne(input)}")
    println("Day 10 part 2: ${partTwo(input)}")
}

private fun partOne(input: String): Int {
    // Create list of numbers from
    val ranges = input.split(",").map { it.toInt() }
    val numbers = (0..255).toMutableList()
    var current = 0
    var skipSize = 0
    for (range in ranges) {
        // Reverse sub array with Length l
        reverseSubArray(current, range, numbers)
        current += range + skipSize
        skipSize += 1
    }

    return numbers[0] * numbers[1]
}

private fun partTwo(input: String): String {
    val ranges = input.map { it.code }.toMutableList()
    ranges.addAll(listOf(17, 31, 73, 47, 23))
    val numbers = (0..255).toMutableList()
    var skipSize = 0
    var current = 0
    repeat(64) {
        ranges.forEach { range ->
            reverseSubArray(current, range, numbers)
            current += range + skipSize
            skipSize++
        }
    }

    val denseHash = numbers.chunked(16) { it.reduce { acc, value -> acc xor value } }
    return denseHash.joinToString("") { it.toString(16).padStart(2, '0') }
}

private fun reverseSubArray(
    start: Int,
    length: Int,
    numbers: MutableList<Int>,
) {
    val subList = (start until start + length).map { numbers[it % numbers.size] }.reversed()
    subList.forEachIndexed { index, value -> numbers[(start + index) % numbers.size] = value }
}

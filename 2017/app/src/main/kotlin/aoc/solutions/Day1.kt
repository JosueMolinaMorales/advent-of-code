package aoc.solutions

import aoc.utils.FileLoader

fun solveDayOne() {
    val input = FileLoader("day1.txt").readFile()
    println("Day 1 part 1: ${partOne(input)}")
    println("Day 1 part 2: ${partTwo(input)}")
}

private fun partOne(input: String): Int {
    var sum = 0
    for (i in input.indices) {
        if (input[i] == input[(i + 1) % input.length]) {
            sum += input[i].digitToInt()
        }
    }
    return sum
}

private fun partTwo(input: String): Int {
    var sum = 0
    val mid = input.length / 2
    for (i in input.indices) {
        if (input[i] == input[(i + mid) % input.length]) {
            sum += input[i].digitToInt()
        }
    }
    return sum
}

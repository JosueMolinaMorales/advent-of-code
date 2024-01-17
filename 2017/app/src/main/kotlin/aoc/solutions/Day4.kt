package aoc.solutions

import aoc.utils.FileLoader

fun solveDayFour() {
    val input = FileLoader("day4.txt").readLines()
    println("Day 4 part 1: ${partOne(input)}")
    println("Day 4 part 2: ${partTwo(input)}")
}

private fun partOne(input: List<String>): Int {
    return input.count { line ->
        val words = line.split(" ")
        words.distinct().size == words.size
    }
}

private fun partTwo(input: List<String>): Int {
    return input.count { line ->
        val words =
            line.split(" ").map { word ->
                word.groupingBy { it }.eachCount()
            }
        words.none { word -> words.count { it == word } > 1 }
    }
}

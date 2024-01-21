package aoc.solutions

import aoc.utils.FileLoader

fun solveDaySixteen() {
    val input = FileLoader("day16.txt").readFile()
    println("Day 16 part 1: ${partOne(input)}")
    println("Day 16 part 2: ${partTwo(input)}")
}

private fun partOne(input: String): String {
    val moves = input.split(",")
    val programs = ('a'..'p').toMutableList()
    for (move in moves) {
        movePrograms(programs, move)
    }

    return programs.joinToString("")
}

private fun partTwo(input: String): String {
    val moves = input.split(",")
    val programs = ('a'..'p').toMutableList()
    val seen = mutableSetOf<String>()
    for (i in 0 until 1000000000) {
        if (programs.joinToString("") in seen) {
            return seen.elementAt(1000000000 % i)
        }
        seen.add(programs.joinToString(""))
        for (move in moves) {
            movePrograms(programs, move)
        }
    }

    return programs.joinToString("")
}

private fun movePrograms(
    programs: MutableList<Char>,
    move: String,
) {
    when (move[0]) {
        's' -> {
            // Spin
            val amount = move.substring(1).toInt()
            val subList = programs.takeLast(amount)
            programs.removeAll(programs.takeLast(amount))
            programs.addAll(0, subList)
        }

        'x' -> {
            // Exchange
            val (a, b) = move.substring(1).split("/").map { it.toInt() }
            val temp = programs[a]
            programs[a] = programs[b]
            programs[b] = temp
        }

        'p' -> {
            // Partner
            val (a, b) = move.substring(1).split("/")
            val indexA = programs.indexOf(a[0])
            val indexB = programs.indexOf(b[0])
            val temp = programs[indexA]
            programs[indexA] = programs[indexB]
            programs[indexB] = temp
        }
    }
}

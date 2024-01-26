package aoc.solutions

import aoc.utils.FileLoader

fun solveDayTwentyOne() {
    val input = FileLoader("day21.txt").readLines()
    println("Day 21 part 1: ${partOne(input)}")
    println("Day 21 part 2: ${partTwo(input)}")
}

private fun solve(
    input: List<String>,
    iterations: Int,
): Int {
    var grid =
        listOf(
            listOf('.', '#', '.'),
            listOf('.', '.', '#'),
            listOf('#', '#', '#'),
        )
    val rules = mutableMapOf<String, List<List<Char>>>()
    input.forEach {
        val parts = it.split(" => ")
        var rotated = parse(parts[0])
        // Add all rotations
        repeat(4) {
            rules[key(rotated)] = parse(parts[1])
            rotated = rotate(rotated)
        }
        // Add all flips
        rotated = flip(rotated)
        repeat(4) {
            rules[key(rotated)] = parse(parts[1])
            rotated = rotate(rotated)
        }
    }

    repeat(iterations) {
        grid = iterate(grid, rules)
    }

    return grid.sumOf { it.count { it == '#' } }
}

private fun partTwo(input: List<String>): Int {
    return solve(input, 18)
}

private fun partOne(input: List<String>): Int {
    return solve(input, 5)
}

/**
 * Breaks the grid into 2x2 or 3x3 squares and applies the rules to each square
 * then reassembles the grid
 */
private fun iterate(
    grid: List<List<Char>>,
    rules: Map<String, List<List<Char>>>,
): List<List<Char>> {
    val squares = mutableListOf<List<List<Char>>>()
    val squareSize = if (grid.size % 2 == 0) 2 else 3

    for (i in grid.indices step squareSize) {
        for (j in grid[i].indices step squareSize) {
            // Break the grid into squares
            val square = mutableListOf<MutableList<Char>>()
            for (k in i until i + squareSize) {
                square.add(mutableListOf())
                for (l in j until j + squareSize) {
                    square.last().add(grid[k][l])
                }
            }
            // Apply the rules
            val newSquare = rules[key(square)]!!
            squares.add(newSquare)
        }
    }

    // Reassemble the grid
    val newGrid = mutableListOf<MutableList<Char>>()
    val numSquares = grid.size / squareSize
    for (i in 0 until numSquares) {
        for (j in 0 until squareSize + 1) {
            newGrid.add(mutableListOf())
            for (k in 0 until numSquares) {
                newGrid.last().addAll(squares[i * numSquares + k][j])
            }
        }
    }
    return newGrid
}

private fun key(grid: List<List<Char>>): String {
    return grid.joinToString("/") { it.joinToString("") }
}

private fun parse(gridStr: String): List<List<Char>> {
    return gridStr.split("/").map { it.toList() }
}

private fun rotate(grid: List<List<Char>>): List<List<Char>> {
    return flip(symmetric(grid))
}

/*
inverts x and y
 */
private fun symmetric(grid: List<List<Char>>): List<List<Char>> {
    val newGrid = mutableListOf<MutableList<Char>>()
    for (i in grid.indices) {
        newGrid.add(mutableListOf())
        for (j in grid[i].indices) {
            newGrid[i].add(grid[j][i])
        }
    }
    return newGrid
}

/*
inverts y
 */
private fun flip(grid: List<List<Char>>): List<List<Char>> {
    val newGrid = mutableListOf<MutableList<Char>>()
    for (i in grid.indices) {
        newGrid.add(mutableListOf())
        for (j in grid[i].indices) {
            newGrid[i].add(grid[i][grid[i].size - 1 - j])
        }
    }
    return newGrid
}

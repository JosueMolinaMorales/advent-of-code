package aoc.solutions

import aoc.utils.FileLoader
import java.util.*

fun solveDayTwelve() {
    val input = FileLoader("day12.txt").readLines()
    println("Day 12 part 1: ${partOne(input)}")
    println("Day 12 part 2: ${partTwo(input)}")
}

private fun partOne(input: List<String>): Int {
    val nodes = buildGraph(input)
    // From every node, dfs to 0
    return nodes.keys.map { dfs(it, nodes, true).contains("0") }.count { it }
}

private fun partTwo(input: List<String>): Int {
    val nodes = buildGraph(input)
    return nodes.keys.map {
        dfs(it, nodes).sorted()
    }.distinct().count()
}

private fun buildGraph(input: List<String>): HashMap<String, MutableList<String>> {
    val nodes = HashMap<String, MutableList<String>>()
    input.forEach {
        val parts = it.split(" <-> ")
        if (parts[0] !in nodes) {
            nodes[parts[0]] = mutableListOf()
        }
        for (child in parts[1].split(", ")) {
            if (child !in nodes) {
                nodes[child] = mutableListOf()
            }
            nodes[parts[0]]!!.add(child)
        }
    }

    return nodes
}

private fun dfs(
    start: String,
    nodes: HashMap<String, MutableList<String>>,
    containsZero: Boolean = false,
): List<String> {
    val stack = Stack<String>()
    val visited = HashSet<String>()
    val path = mutableListOf<String>()
    stack.push(start)
    visited.add(start)
    while (!stack.isEmpty()) {
        // Pop the stack
        val curr = stack.pop()
        // Add to path
        path.add(curr)
        if (containsZero && curr == "0") {
            return path
        }
        // Visit the adj nodes
        for (child in nodes[curr]!!) {
            if (child in visited) {
                continue
            }
            visited.add(child)
            stack.push(child)
        }
    }

    return path.toList()
}

package aoc.solutions

import aoc.utils.FileLoader

fun solveDayTwentyFour() {
    val input = FileLoader("day24.txt").readLines()
    println("Day 24 part 1: ${partOne(input)}")
    println("Day 24 part 2: ${partTwo(input)}")
}

private fun partTwo(input: List<String>): Int {
    val components = input.map { Component.fromLine(it) }
    val bridges = mutableListOf<List<Component>>()
    val used = mutableSetOf<Component>()
    buildBridges(0, components, bridges, used)
    val max = bridges.maxOf { it.size }
    return bridges.filter { it.size == max }.maxOf { it.sumOf { it.a + it.b } }
}

private fun partOne(input: List<String>): Int {
    val components = input.map { Component.fromLine(it) }
    val bridges = mutableListOf<List<Component>>()
    val used = mutableSetOf<Component>()
    buildBridges(0, components, bridges, used)
    return bridges.maxOf { it.sumOf { it.a + it.b } }
}

private fun buildBridges(
    port: Int,
    components: List<Component>,
    bridges: MutableList<List<Component>>,
    used: MutableSet<Component>,
    currBridge: List<Component> = listOf(),
) {
    val possibleComponents = components.filter { it.a == port || it.b == port }
    for (component in possibleComponents) {
        if (component in used) {
            continue
        }
        val newBridge = currBridge + component
        used.add(component)
        buildBridges(if (component.a == port) component.b else component.a, components, bridges, used, newBridge)
        used.remove(component)
    }

    bridges.add(currBridge)
}

private data class Component(val a: Int, val b: Int) {
    companion object {
        fun fromLine(line: String): Component {
            val (a, b) = line.split("/").map { it.toInt() }
            return Component(a, b)
        }
    }
}

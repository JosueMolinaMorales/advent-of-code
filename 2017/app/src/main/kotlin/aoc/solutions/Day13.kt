package aoc.solutions

import aoc.utils.Direction
import aoc.utils.FileLoader

fun solveDayThirteen() {
    val input = FileLoader("day13.txt").readLines()
    println("Day 13 part 1: ${partOne(input)}")
    println("Day 13 part 2: ${partTwo(input)}")
}

private fun partOne(input: List<String>): Int {
    val firewall = mutableMapOf<Int, Layer>()
    input.forEach {
        val layer = Layer.fromLine(it)
        firewall[layer.id] = layer
    }
    var currLayer = -1
    val maxFirewall = firewall.maxBy { it.key }.key
    var res = 0
    while (currLayer <= maxFirewall) {
        // Move to next layer
        currLayer += 1
        // Check to see if there is a scanner
        if (firewall[currLayer]?.scanner == 0) {
            res += (firewall[currLayer]!!.id * firewall[currLayer]!!.depth)
        }
        // Move all scanners
        for (layer in firewall.values) {
            layer.moveScanner()
        }
    }

    return res
}

private fun partTwo(input: List<String>): Int {
    val firewall = mutableMapOf<Int, Layer>()
    input.forEach {
        val layer = Layer.fromLine(it)
        firewall[layer.id] = layer
    }

    var delay = 0
    while (true) {
        var caught = false
        for (layer in firewall.values) {
            if ((layer.id + delay) % ((layer.depth - 1) * 2) == 0) {
                caught = true
                break
            }
        }
        if (!caught) {
            return delay
        }
        delay += 1
    }
}

private class Layer(val id: Int, val depth: Int) {
    var scanner = 0
    var direction = Direction.DOWN

    companion object {
        fun fromLine(line: String): Layer {
            val parts = line.split(": ")
            return Layer(parts[0].toInt(), parts[1].toInt())
        }
    }

    fun moveScanner() {
        if (scanner == 0) {
            direction = Direction.DOWN
        } else if (scanner == depth - 1) {
            direction = Direction.UP
        }
        scanner +=
            if (direction == Direction.UP) {
                -1
            } else {
                1
            }
    }
}

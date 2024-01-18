package aoc.solutions

import aoc.utils.FileLoader
import kotlin.math.ceil

fun solveDaySix() {
    val input = FileLoader("day6.txt").readFile()
    println("Day 6 part 1: ${partOne(input)}")
    println("Day 6 part 2: ${partTwo(input)}")
}

private fun partOne(input: String): Int {
    val states = HashSet<List<Int>>()

    val buckets = input.split("\\s+".toRegex()).map { it.toInt() }.toMutableList()
    var count = 0

    while (buckets !in states) {
        states.add(buckets.toList())
        val maxIdx = buckets.indexOf(buckets.max())
        redistributeBuckets(buckets, maxIdx)
        count++
    }

    return count
}

private fun redistributeBuckets(
    buckets: MutableList<Int>,
    maxIdx: Int,
) {
    var toDistribute = buckets[maxIdx]
    val distributeAmt = ceil((toDistribute / buckets.size.toDouble())).toInt()
    buckets[maxIdx] = 0

    var currentIdx = (maxIdx + 1) % buckets.size
    while (toDistribute > 0) {
        buckets[currentIdx] += distributeAmt
        toDistribute -= distributeAmt
        currentIdx = (currentIdx + 1) % buckets.size
    }
}

private fun partTwo(input: String): Int {
    val states = HashMap<List<Int>, Int>()
    val buckets = input.split("\\s+".toRegex()).map { it.toInt() }.toMutableList()
    var count = 0
    var seen = false

    while (true) {
        val state = buckets.toList()
        if (state in states) {
            if (states[state]!! > 1) {
                return count
            }
            states[state] = states[state]!! + 1
            seen = true
        } else {
            states[state] = 1
        }

        val maxIdx = buckets.indexOf(buckets.max())
        redistributeBuckets(buckets, maxIdx)

        if (seen) {
            count++
        }
    }
}

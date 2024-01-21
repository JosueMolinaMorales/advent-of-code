package aoc.solutions

fun solveDayFifteen() {
    println("Day 15 part 1: ${partOne()}")
    println("Day 15 part 2: ${partTwo()}")
}

private fun partOne(): Int {
    val generatorA = Generator(16807, 783)
    val generatorB = Generator(48271, 325)
    var matches = 0
    for (i in 0 until 40_000_000) {
        val a = generatorA.next()
        val b = generatorB.next()
        if (a and 0xFFFF == b and 0xFFFF) {
            matches += 1
        }
    }
    return matches
}

private fun partTwo(): Int {
    val generatorA = Generator(16807, 783, 4)
    val generatorB = Generator(48271, 325, 8)
    var matches = 0
    for (i in 0 until 5_000_000) {
        val a = generatorA.next()
        val b = generatorB.next()
        if (a and 0xFFFF == b and 0xFFFF) {
            matches += 1
        }
    }
    return matches
}

private class Generator(
    val factor: Int,
    start: Int,
    val muliple: Int = 1,
) {
    var curr = start.toLong()

    fun next(): Long {
        do {
            curr = (curr * factor) % 2147483647
        } while (curr % muliple != 0L)
        return curr
    }
}

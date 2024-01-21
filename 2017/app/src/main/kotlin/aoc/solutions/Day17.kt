package aoc.solutions

fun solveDaySeventeen() {
    val steps = 356
    println("Day 17 part 1: ${partOne(steps)}")
    println("Day 17 part 2: ${partTwo(steps)}")
}

private fun partOne(steps: Int): Int {
    val list = mutableListOf<Int>(0)
    var currPos = 0
    (1..2017).forEach {
        // Take steps
        currPos = ((currPos + steps) % list.size) + 1
        // Insert next value
        list.add(currPos, it)
    }
    return list[(currPos + 1) % list.size]
}

private fun partTwo(steps: Int): Int {
    return 0
}

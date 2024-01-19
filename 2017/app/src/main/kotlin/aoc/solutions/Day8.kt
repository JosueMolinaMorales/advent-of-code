package aoc.solutions

import aoc.utils.FileLoader

fun solveDayEight() {
    val input = FileLoader("day8.txt").readLines()

    println("Day 8 part 1: ${partOne(input)}")
    println("Day 8 part 2: ${partTwo(input)}")
}

private fun partTwo(input: List<String>): Int {
    return solve(input, true)
}

private fun partOne(input: List<String>): Int {
    return solve(input, false)
}

private fun solve(
    input: List<String>,
    getMaxAllTime: Boolean,
): Int {
    val instructions = input.map { parseInstruction(it) }
    val registers = instructions.flatMap { listOf(it.register, it.r1) }.distinct().associateWith { 0 }.toMutableMap()

    var maxAllTime = 0
    for (inst in instructions) {
        if (inst.test(registers)) {
            val offset =
                when (inst.action) {
                    "inc" -> inst.diff
                    "dec" -> -inst.diff
                    else -> 0
                }
            registers[inst.register] = registers.getValue(inst.register) + offset
            val currMax = registers.values.maxOrNull()
            if (currMax != null && currMax > maxAllTime) {
                maxAllTime = currMax
            }
        }
    }
    return if (getMaxAllTime) maxAllTime else registers.values.maxOrNull() ?: 0
}

private fun parseInstruction(line: String): Instruction {
    val parts = line.split(" ")
    return Instruction(
        register = parts[0],
        action = parts[1],
        diff = parts[2].toInt(),
        r1 = parts[4],
        cond = Condition.fromStr(parts[5]),
        num = parts[6].toInt(),
    )
}

private enum class Condition {
    LessThan,
    GreaterThan,
    LessThanEqual,
    GreaterThanEqual,
    NotEqual,
    Equal,
    ;

    companion object {
        fun fromStr(str: String): Condition {
            return when (str) {
                "<" -> LessThan
                ">" -> GreaterThan
                ">=" -> GreaterThanEqual
                "<=" -> LessThanEqual
                "==" -> Equal
                "!=" -> NotEqual
                else -> throw IllegalArgumentException("IDK")
            }
        }
    }
}

private class Instruction(
    val register: String,
    val action: String,
    val diff: Int,
    val r1: String,
    val cond: Condition,
    val num: Int,
) {
    fun test(registers: Map<String, Int>): Boolean {
        return when (this.cond) {
            Condition.LessThan -> registers[this.r1]!! < this.num
            Condition.LessThanEqual -> registers[this.r1]!! <= this.num
            Condition.NotEqual -> registers[this.r1]!! != this.num
            Condition.Equal -> registers[this.r1]!! == this.num
            Condition.GreaterThanEqual -> registers[this.r1]!! >= this.num
            Condition.GreaterThan -> registers[this.r1]!! > this.num
        }
    }
}

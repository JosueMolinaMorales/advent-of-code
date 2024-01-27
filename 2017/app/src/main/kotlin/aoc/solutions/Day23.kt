package aoc.solutions

import aoc.utils.FileLoader

fun solveDayTwentyThree() {
    val input = FileLoader("day23.txt").readLines()
    println("Day 23 part 1: ${partOne(input)}")
    println("Day 23 part 2: ${partTwo(input)}")
}

private fun partTwo(input: List<String>): Int {
    // The program is counting the number of non-primes between b and c, in increments of 17
    // Get the values of b and c
    val instructions = input.map { InstructionDay23.fromLine(it) }
    val registers = mutableMapOf<String, Long>()
    registers["a"] = 1
    var i = 0
    while (i in instructions.indices) {
        if (i == 8) {
            break
        }
        val instruction = instructions[i]
        when (instruction.action) {
            "set" -> registers[instruction.x] = eval(instruction.y, registers)
            "sub" -> registers[instruction.x] = eval(instruction.x, registers) - eval(instruction.y, registers)
            "mul" -> registers[instruction.x] = eval(instruction.x, registers) * eval(instruction.y, registers)
            "jnz" -> {
                if (eval(instruction.x, registers) != 0L) {
                    i += eval(instruction.y, registers).toInt()
                    continue
                }
            }
        }
        i += 1
    }
    // Now count the number of non-primes between b and c, in increments of 17
    var count = 0
    for (n in registers["b"]!!..registers["c"]!! step 17) {
        if (!isPrime(n)) {
            count += 1
        }
    }
    return count
}

private fun isPrime(n: Long): Boolean {
    if (n <= 1) {
        return false
    }
    if (n <= 3) {
        return true
    }
    if (n % 2 == 0L || n % 3 == 0L) {
        return false
    }
    var i = 5L
    while (i * i <= n) {
        if (n % i == 0L || n % (i + 2) == 0L) {
            return false
        }
        i += 6
    }
    return true
}

private fun partOne(input: List<String>): Int {
    val instructions = input.map { InstructionDay23.fromLine(it) }
    val registers = mutableMapOf<String, Long>()
    var i = 0
    var mulCount = 0
    while (i in instructions.indices) {
        val instruction = instructions[i]
        when (instruction.action) {
            "set" -> registers[instruction.x] = eval(instruction.y, registers)
            "sub" -> registers[instruction.x] = eval(instruction.x, registers) - eval(instruction.y, registers)
            "mul" -> {
                registers[instruction.x] = eval(instruction.x, registers) * eval(instruction.y, registers)
                mulCount += 1
            }

            "jnz" -> {
                if (eval(instruction.x, registers) != 0L) {
                    i += eval(instruction.y, registers).toInt()
                    continue
                }
            }
        }
        i += 1
    }
    return mulCount
}

private fun eval(
    key: String,
    registers: Map<String, Long>,
): Long {
    if (key.toIntOrNull() == null) {
        return registers.getOrDefault(key, 0)
    }
    return key.toLong()
}

private data class InstructionDay23(
    val action: String,
    val x: String,
    val y: String,
) {
    companion object {
        fun fromLine(line: String): InstructionDay23 {
            val split = line.split(" ")
            return InstructionDay23(split[0], split[1], split[2])
        }
    }
}

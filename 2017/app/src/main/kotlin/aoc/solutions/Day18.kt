package aoc.solutions

import aoc.utils.FileLoader

fun solveDayEightteen() {
    val input = FileLoader("day18.txt").readLines()
    println("Day 18 part 1: ${partOne(input)}")
    println("Day 18 part 2: ${partTwo(input)}")
}

private fun partTwo(input: List<String>): Long {
    val instructions = input.map { InstructionDay18.fromLine(it) }
    val registers =
        listOf(
            mutableMapOf<String, Long>(),
            mutableMapOf(),
        )
    registers[0]["p"] = 0
    registers[1]["p"] = 1
    val queues =
        listOf(
            mutableListOf<Long>(),
            mutableListOf(),
        )
    val eval: ((key: String, id: Int) -> Long) = { key, id ->
        if (key.toIntOrNull() == null) {
            registers[id].getOrPut(key) { 0 }
        } else {
            key.toLong()
        }
    }
    val i = mutableListOf(0, 0)
    var currId = 0
    var res = 0L
    while (i[currId] in instructions.indices) {
        val instruction = instructions[i[currId]]
        when (instruction.action) {
            "snd" -> {
                queues[1 - currId].add(eval(instruction.x, currId))
                if (currId == 1) {
                    res += 1
                }
            }

            "set" -> registers[currId][instruction.x] = eval(instruction.y, currId)
            "add" -> registers[currId][instruction.x] = eval(instruction.x, currId) + eval(instruction.y, currId)
            "mul" -> registers[currId][instruction.x] = eval(instruction.x, currId) * eval(instruction.y, currId)
            "mod" -> registers[currId][instruction.x] = eval(instruction.x, currId) % eval(instruction.y, currId)
            "rcv" -> {
                if (queues[0].isEmpty() && queues[1].isEmpty()) {
                    return res
                }
                if (queues[currId].isEmpty()) {
                    currId = 1 - currId
                    continue
                } else {
                    registers[currId][instruction.x] = queues[currId].removeAt(0)
                }
            }

            "jgz" ->
                if (eval(instruction.x, currId) > 0L) {
                    i[currId] += eval(instruction.y, currId).toInt()
                    continue
                }
        }
        i[currId] += 1
    }
    return res
}

private fun partOne(input: List<String>): Long {
    val instructions = input.map { InstructionDay18.fromLine(it) }
    val registers = mutableMapOf<String, Long>()
    val eval: ((key: String) -> Long) = {
        if (it.toIntOrNull() == null) {
            registers.getOrPut(it) { 0 }
        } else {
            it.toLong()
        }
    }
    var i = 0
    var freq = 0L
    while (i in instructions.indices) {
        val instruction = instructions[i]
        when (instruction.action) {
            "snd" -> freq = eval(instruction.x)
            "set" -> registers[instruction.x] = eval(instruction.y)
            "add" -> registers[instruction.x] = eval(instruction.x) + eval(instruction.y)
            "mul" -> registers[instruction.x] = eval(instruction.x) * eval(instruction.y)
            "mod" -> registers[instruction.x] = eval(instruction.x) % eval(instruction.y)
            "rcv" ->
                if (eval(instruction.x) != 0L) {
                    return freq
                }

            "jgz" ->
                if (eval(instruction.x) > 0L) {
                    i += eval(instruction.y).toInt()
                    continue
                }
        }
        i += 1
    }
    return freq
}

private class InstructionDay18(val action: String, val x: String, val y: String = "") {
    companion object {
        fun fromLine(line: String): InstructionDay18 {
            val parts = line.split(" ")
            return InstructionDay18(parts[0], parts[1], parts.getOrElse(2) { "" })
        }
    }
}

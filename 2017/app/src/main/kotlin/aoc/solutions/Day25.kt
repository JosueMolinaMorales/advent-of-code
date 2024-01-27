package aoc.solutions

import aoc.utils.FileLoader

fun solveDayTwentyFive() {
    val input = FileLoader("day25.txt").readFile()
    println("Day 25 part 1: ${partOne(input)}")
}

private fun partOne(input: String): Int {
    val instructions = input.split("\n\n")
    val states = mutableMapOf<String, StateAction>()
    val checksumAfter = instructions[0].split("\n")[1].split("after ")[1].split(" ")[0].toInt()
    val startState = instructions[0].split("\n")[0].split("state ")[1].split(".")[0]
    var i = 1
    while (i in instructions.indices) {
        val stateParts = instructions[i].split("\n")
        val state = stateParts[0].split("state ")[1].split(":")[0]

        val write0 = stateParts[2].split("value ")[1].split(".")[0].toInt()
        val move0 = if (stateParts[3].split("one slot")[1].contains("left")) -1 else 1
        val nextState0 = stateParts[4].split("state ")[1].split(".")[0]

        val write1 = stateParts[6].split("value ")[1].split(".")[0].toInt()
        val move1 = if (stateParts[7].split("one slot")[1].contains("left")) -1 else 1
        val nextState1 = stateParts[8].split("state ")[1].split(".")[0]
        states[state + "0"] = StateAction(write0, move0, nextState0)
        states[state + "1"] = StateAction(write1, move1, nextState1)
        i += 1
    }
    val tape = mutableMapOf<Int, Int>()
    var currState = startState
    var currPos = 0
    repeat(checksumAfter) {
        val currVal = tape.getOrPut(currPos) { 0 }
        val action = states[currState + currVal]!!
        tape[currPos] = action.write
        currPos += action.move
        currState = action.nextState
    }
    return tape.values.sum()
}

private data class StateAction(val write: Int, val move: Int, val nextState: String)

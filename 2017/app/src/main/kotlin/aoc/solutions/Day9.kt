@file:Suppress("ktlint:standard:no-wildcard-imports")

package aoc.solutions

import aoc.utils.FileLoader
import java.util.*

fun solveDayNine() {
    val input = FileLoader("day9.txt").readFile()
    println("Day 9 part 1: ${partOne(input)}")
    println("Day 9 part 2: ${partTwo(input)}")
}

private fun partTwo(input: String): Int {
    var res = 0
    var i = 0
    var isInGarbage = false
    while (i in input.indices) {
        val currChar = input[i]
        when {
            (currChar == '!') -> {
                // Skip the next character
                i += 2
                continue
            }

            (currChar == '>' && isInGarbage) -> isInGarbage = false
            (isInGarbage) -> {
                res += 1
            }

            currChar == '<' -> isInGarbage = true

            else -> {}
        }
        i += 1
    }

    return res
}

private fun partOne(input: String): Int {
    val stack = Stack<Char>()
    var res = 0
    var currScore = 1

    var i = 0
    while (i in input.indices) {
        // important character: {}, <>, !
        val isInGarbage = !stack.isEmpty() && stack.peek() == '<'
        val currChar = input[i]
        when {
            (currChar == '!') -> {
                // Skip the next character
                i += 2
                continue
            }

            (currChar == '{' && !isInGarbage) -> {
                // If the previous was '{' we moved one in, inc score
                if (stack.size > 0 && stack.peek() == '{') {
                    currScore += 1
                }
                stack.push(currChar)
            }

            (currChar == '}' && !isInGarbage) -> {
                res += currScore
                stack.pop()
                currScore -= 1
            }

            currChar == '<' -> {
                if (stack.peek() != '<') {
                    // If the last important char was a < we are in garbage,
                    // Ignore proceeding '<'s
                    stack.push(currChar)
                }
            }

            currChar == '>' -> stack.pop()
            else -> {}
        }
        i += 1
    }

    return res
}

package aoc.solutions

import aoc.utils.FileLoader
import kotlin.math.abs

fun solveDaySeven() {
    val programs = FileLoader("day7.txt").readLines()
    println("Day 7 part 1: ${partOne(programs)}")
    println("Day 7 part 2: ${partTwo(programs)}")
}

private fun partTwo(input: List<String>): Int {
    val result = getParentAndPrograms(input)
    val root = result!!.root
    val programs = result.programs
    var found = false
    var res = 0
    getSumForChildren(root, programs) { innerSums ->
        val sums = innerSums.values.distinct()
        if (!found && sums.size > 1) {
            found = true
            // Find which child has the sum that isnt the same
            val minSum = innerSums.values.groupingBy { it }.eachCount().minBy { it.value }.key
            for ((id, s) in innerSums) {
                if (s == minSum) {
                    res = programs[id]!!.weight -
                        abs(
                            (sums.max() - sums.min()),
                        )
                }
            }
        }
    }

    return res
}

private fun partOne(input: List<String>): String {
    return getParentAndPrograms(input)!!.root
}

private fun getSumForChildren(
    root: String,
    programs: HashMap<String, Program>,
    visit: (sums: HashMap<String, Int>) -> Unit,
): Int {
    val sums = HashMap<String, Int>()
    for (child in programs[root]!!.children) {
        val sum = getSumForChildren(child, programs, visit)
        sums[child] = sum
    }

    visit(sums)

    var sum = programs[root]!!.weight
    for (s in sums.values) {
        sum += s
    }
    return sum
}

private fun getParentAndPrograms(input: List<String>): Result? {
    val programs = HashMap<String, Program>()
    for (line in input) {
        val parts = line.split(" ")
        val id = parts[0]
        val weight = parts[1].substring(1, parts[1].length - 1).toInt()
        val children = HashSet<String>()
        if (line.contains("->")) {
            children.addAll(line.split(" -> ")[1].split(", ").toHashSet())
        }
        // Have already seen this program
        programs[id] =
            Program(
                id, weight, children,
                if (id in programs) {
                    programs[id]!!.parent
                } else {
                    null
                },
            )
        // If the program has children
        if (children.size == 0) {
            continue
        }
        // For each child, set its parent
        for (child in children) {
            if (programs.contains(child)) {
                programs[child]!!.parent = id
            } else {
                programs[child] = Program(child, 0, HashSet(), id)
            }
        }
    }

    for ((id, program) in programs) {
        if (program.parent == null) {
            return Result(programs, id)
        }
    }

    return null
}

private data class Result(val programs: HashMap<String, Program>, val root: String)

private data class Program(val id: String, val weight: Int, val children: HashSet<String>, var parent: String? = null)

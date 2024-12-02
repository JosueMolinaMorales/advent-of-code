def sign(a):
    if a > 0:
        return 1
    elif a < 0:
        return -1
    else:
        return 0


def check(report):
    a, b = 0, 1

    last_sign = 0
    while a < b < len(report):
        delta = report[a] - report[b]
        curr_sign = sign(delta)

        # we've had a change of direction.
        if last_sign and curr_sign != last_sign:
            return False
        last_sign = curr_sign

        d = abs(delta)
        if 0 < d < 4:
            a += 1
            b += 1
            continue
        # not a safe inc/dec amount

        return False

    return True


def check_remove_one(report):
    if check(report):
        return True

    for i in range(len(report)):
        part = report[:i] + report[i+1:]
        if check(part):
            return True

    return False


reports = []

with open("./inputs/day_2.txt", "r") as file:
    for report in file:
        reports.append([int(part) for part in report.rstrip().split(' ')])

count = 0
for level in reports:
    if check_remove_one(level):
        count += 1

print("safe: {}".format(count))

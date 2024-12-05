import re
from itertools import chain

with open("day3/day3.txt", "r") as file:
    machine = file.read().strip()


def enabled_sum():
    data = list(chain.from_iterable(
        map(lambda x: re.findall(
            r"mul\((\d+),(\d+)\)", x.split("don't")[0]),
            re.split(r"do(?!n\'t)", machine))))
    return sum(map(lambda x: int(x[0]) * int(x[1]), data))


def main():
    print(enabled_sum())


if __name__ == "__main__":
    main()

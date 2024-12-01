left_numbers: list[int] = []
right_numbers: list[int] = []

with open("input.txt", encoding="utf-8") as input:
    for line in input:
        left, right = line.split("   ")
        left_numbers.append(int(left))
        right_numbers.append(int(right))

left_numbers.sort()

sum: int = 0

for left in left_numbers:
    sum += left * right_numbers.count(left)

print(sum)


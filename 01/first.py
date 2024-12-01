left_numbers: list[int] = []
right_numbers: list[int] = []

with open("input.txt", encoding="utf-8") as input:
    for line in input:
        left, right = line.split("   ")
        left_numbers.append(int(left))
        right_numbers.append(int(right))

left_numbers.sort()
right_numbers.sort()

sum: int = 0

for left, right in zip(left_numbers, right_numbers):
    sum += abs(left - right)

print(sum)


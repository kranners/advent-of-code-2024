const { readFileSync } = require("fs");

const input = readFileSync("./input.txt").toString();

const COMMAND_PATTERN = /mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)/g;

const commands = [...input.match(COMMAND_PATTERN)];

const parseMulCommand = (command) => {
  if (!doPerformMultiplication) {
    return 0;
  }

  const [multiplicand, multiplier] = [...command.matchAll(/[0-9]+/g)].map(Number);
  return multiplicand * multiplier;
}

let doPerformMultiplication = true;

const parseCommand = (command) => {
  if (command === "do()") {
    doPerformMultiplication = true;
    return 0;
  }

  if (command === "don't()") {
    doPerformMultiplication = false;
    return 0;
  }

  return parseMulCommand(command);
}

const sum = commands
  .map(parseCommand)
  .reduce((acc, result) => acc + result, 0);

console.log(sum);


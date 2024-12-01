"use strict";
function Reduce(operation) {
    return (numbers) => {
        return numbers.length == 0 ? 0
            : numbers.length == 1 ? numbers[0]
                : operation(numbers[0], Reduce(operation)(numbers.slice(1)));
    };
}
console.log(Reduce((x, y) => x + y)([1, 2, 3, 4, 5]));

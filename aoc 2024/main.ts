function Reduce(operation: (n: number, a: number) => number): (numbers: number[]) => number {
    return (numbers: number[]) => {
        return numbers.length == 0 ? 0
        : numbers.length == 1 ? numbers[0] 
        : operation(numbers[0], Reduce(operation)(numbers.slice(1)));
    };
}

console.log(Reduce((x: number, y: number) => x + y)([1,2,3,4,5]));
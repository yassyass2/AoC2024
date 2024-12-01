import * as fs from 'fs';

namespace day1{

    function Fold<T, R>(fxy: (acc: R, current: T) => R, seed: R){
        return function(arr: T[]): R{
            if (arr.length <= 0) return seed
            let Result = seed;
            for(let i = 0; i < arr.length; i++){
                Result = fxy(Result, arr[i]);
            }
            return Result;
        }
    }
    console.log(Fold<number, number>((x: number, y: number) => x + y, 0)([1,1,1,1,1]))

    
    function loadFileLines(filePath: string): string[] {
        const fileContent = fs.readFileSync(filePath, 'utf-8');
        return fileContent.split('\n');
    }

    const lines = loadFileLines("./dag1.txt");
    const cleanedLines = lines.map(line => line.replace(/\r/g, ''));
    console.log(cleanedLines);

    const LeftNums: number[] = [];
    const RightNums: number[] = [];

    cleanedLines.forEach(line => {
        const [num1, num2] = line.split('   ').map(Number);
        LeftNums.push(num1);
        RightNums.push(num2);
    });

    
    const combined: number[] = [];

    const sortedLeft = LeftNums.sort()
    const sortedRight = RightNums.sort()

    for(let i = 0;i < sortedLeft.length;i++){
        combined.push(Math.abs(sortedRight[i]-sortedLeft[i]))
    }

    console.log(Fold<number, number>((x: number, y: number) => x+y, 0)(combined))


    function similarity(left: number[], right: number[]){
        let scores: number = 0

        for (const num of left) {
            let count: number = 0
            for (const num2 of right) {
                if (num2 === num) {
                    count++;
                }
            }
            scores += num*count
        }
        return scores
    }
    console.log(similarity(sortedLeft, sortedRight))
}

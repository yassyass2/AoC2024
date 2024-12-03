import * as fs from 'fs';

namespace day3{
    export function Fold<T, R>(fxy: (acc: R, current: T) => R, seed: R){
        return function(arr: T[]): R{
            if (arr.length <= 0) return seed
            let Result = seed;
            for(let i = 0; i < arr.length; i++){
                Result = fxy(Result, arr[i]);
            }
            return Result;
        }
    }

    export function loadFileLines(filePath: string): string {
        const fileContent = fs.readFileSync(filePath, 'utf-8');
        return fileContent.split('\n')[0];
    }
    const input = loadFileLines("./day3.txt")

    function findAllMulPatterns(text: string): number {
        const pattern = /mul\((\d{1,3}),(\d{1,3})\)/g;
        const results: number[] = [];
    
        for (const match of text.matchAll(pattern)) {
            const num1 = parseInt(match[1], 10);
            //console.log(num1)
            const num2 = parseInt(match[2], 10);
            //console.log(num2)
            results.push(num1*num2);
        }
    
        return Fold<number, number>((x: number, y: number) => x+y, 0)(results);
    }
    console.log(findAllMulPatterns(input))
}
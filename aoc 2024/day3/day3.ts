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
        const pattern: RegExp = /mul\((\d{1,3}),(\d{1,3})\)/g;
        const doPattern: RegExp = /do\(\)/g
        const dontPattern: RegExp = /don't\(\)/g
        const results: number[] = [];
    
        let enabled = true
        const matches = text.matchAll(/mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)/g)
        for (const match of matches) {
            const [fullMatch, num1, num2] = match;
            if (dontPattern.test(fullMatch)) {
                enabled = false;
            } else if (doPattern.test(fullMatch)) {
                enabled = true;
            } else if (pattern.test(fullMatch) && enabled){
                const n1 = parseInt(num1, 10);
                const n2 = parseInt(num2, 10);
                results.push(n1*n2);
            }
        }
    
        return Fold<number, number>((x: number, y: number) => x+y, 0)(results);
    }
    console.log(findAllMulPatterns(input))
}
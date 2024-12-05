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
            if (pattern.test(fullMatch) && enabled){
                console.log([fullMatch, num1, num2])
                const n1 = parseInt(num1);
                const n2 = parseInt(num2);
                results.push(n1*n2);
            } 
            else if (pattern.test(fullMatch) && !enabled){
                ;
            }
            else if (dontPattern.test(fullMatch)) {
                enabled = !enabled;
            }
            else if (doPattern.test(fullMatch)) {
                enabled = !enabled;
            }
        }
    
        return Fold<number, number>((x: number, y: number) => x+y, 0)(results);
    }
    console.log(findAllMulPatterns(input))
}
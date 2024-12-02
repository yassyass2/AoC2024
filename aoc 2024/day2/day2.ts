import * as fs from 'fs';

namespace day2{
    function canBeValidByRemovingOne(arr: number[]): boolean {
        for (let i = 0; i < arr.length; i++) {
            const newArr = arr.slice(0, i).concat(arr.slice(i + 1));
            
            if (isValidArray(newArr, 1)) {
                return true;
            }
        }
        return false;
    }
    
    function loadFileLines(filePath: string): string[] {
        const fileContent = fs.readFileSync(filePath, 'utf-8');
        return fileContent.split('\n');
    }
    const input = loadFileLines("./day2.txt")

    function isValidArray(arr: number[], depth: number): boolean {
        let operator = true
        for (let i = 1; i < arr.length; i++) {
            if (i == 1){
                operator = arr[i] > arr[i-1]
            }
            let diff = arr[i] - arr[i - 1];
            if ((operator && diff < 0) || (!operator && diff > 0)){
                return depth == 0 ? canBeValidByRemovingOne(arr) : false
            }
            diff = Math.abs(diff)
            if (diff > 3 || diff == 0) {
                return depth == 0 ? canBeValidByRemovingOne(arr) : false
            }
        }
        return true;
    }

    function isSafe(options: string[]): number{
        let count = 0
        for(let i = 0; i < options.length;i++){
            let step = options[i]
            let nums: number[] = step.split(' ').map(Number)
            count += isValidArray(nums, 0) ? 1 : 0;
        }
        return count
    }

    console.log(isSafe(input))
}
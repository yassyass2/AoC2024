"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
const fs = __importStar(require("fs"));
var day1;
(function (day1) {
    function Fold(fxy, seed) {
        return function (arr) {
            if (arr.length <= 0)
                return seed;
            let Result = seed;
            for (let i = 0; i < arr.length; i++) {
                Result = fxy(Result, arr[i]);
            }
            return Result;
        };
    }
    console.log(Fold((x, y) => x + y, 0)([1, 1, 1, 1, 1]));
    function loadFileLines(filePath) {
        const fileContent = fs.readFileSync(filePath, 'utf-8');
        return fileContent.split('\n');
    }
    const lines = loadFileLines("./dag1.txt");
    const cleanedLines = lines.map(line => line.replace(/\r/g, ''));
    console.log(cleanedLines);
    const LeftNums = [];
    const RightNums = [];
    cleanedLines.forEach(line => {
        const [num1, num2] = line.split('   ').map(Number);
        LeftNums.push(num1);
        RightNums.push(num2);
    });
    const combined = [];
    const sortedLeft = LeftNums.sort();
    const sortedRight = RightNums.sort();
    for (let i = 0; i < sortedLeft.length; i++) {
        combined.push(Math.abs(sortedRight[i] - sortedLeft[i]));
    }
    console.log(Fold((x, y) => x + y, 0)(combined));
    function similarity(left, right) {
        let scores = 0;
        for (const num of left) {
            let count = 0;
            for (const num2 of right) {
                if (num2 === num) {
                    count++;
                }
            }
            scores += num * count;
        }
        return scores;
    }
    console.log(similarity(sortedLeft, sortedRight));
})(day1 || (day1 = {}));

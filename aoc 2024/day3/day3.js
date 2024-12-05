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
var day3;
(function (day3) {
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
    day3.Fold = Fold;
    function loadFileLines(filePath) {
        const fileContent = fs.readFileSync(filePath, 'utf-8');
        return fileContent.split('\n')[0];
    }
    day3.loadFileLines = loadFileLines;
    const input = loadFileLines("./day3.txt");
    function findAllMulPatterns(text) {
        const pattern = /mul\((\d{1,3}),(\d{1,3})\)/g;
        const doPattern = /do\(\)/g;
        const dontPattern = /don't\(\)/g;
        const results = [];
        let enabled = true;
        const matches = text.matchAll(/mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)/g);
        for (const match of matches) {
            const [fullMatch, num1, num2] = match;
            if (pattern.test(fullMatch) && enabled) {
                console.log([fullMatch, num1, num2]);
                const n1 = parseInt(num1);
                const n2 = parseInt(num2);
                results.push(n1 * n2);
            }
            else if (pattern.test(fullMatch) && !enabled) {
                ;
            }
            else if (dontPattern.test(fullMatch)) {
                enabled = !enabled;
            }
            else if (doPattern.test(fullMatch)) {
                enabled = !enabled;
            }
        }
        return Fold((x, y) => x + y, 0)(results);
    }
    console.log(findAllMulPatterns(input));
})(day3 || (day3 = {}));

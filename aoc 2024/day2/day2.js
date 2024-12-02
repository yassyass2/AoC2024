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
var day2;
(function (day2) {
    function canBeValidByRemovingOne(arr) {
        for (let i = 0; i < arr.length; i++) {
            const newArr = arr.slice(0, i).concat(arr.slice(i + 1));
            if (isValidArray(newArr, 1)) {
                return true;
            }
        }
        return false;
    }
    function loadFileLines(filePath) {
        const fileContent = fs.readFileSync(filePath, 'utf-8');
        return fileContent.split('\n');
    }
    const input = loadFileLines("./day2.txt");
    function isValidArray(arr, depth) {
        let operator = true;
        for (let i = 1; i < arr.length; i++) {
            if (i == 1) {
                operator = arr[i] > arr[i - 1];
            }
            let diff = arr[i] - arr[i - 1];
            if ((operator && diff < 0) || (!operator && diff > 0)) {
                return depth == 0 ? canBeValidByRemovingOne(arr) : false;
            }
            diff = Math.abs(diff);
            if (diff > 3 || diff == 0) {
                return depth == 0 ? canBeValidByRemovingOne(arr) : false;
            }
        }
        return true;
    }
    function isSafe(options) {
        let count = 0;
        for (let i = 0; i < options.length; i++) {
            let step = options[i];
            let nums = step.split(' ').map(Number);
            count += isValidArray(nums, 0) ? 1 : 0;
        }
        return count;
    }
    console.log(isSafe(input));
})(day2 || (day2 = {}));

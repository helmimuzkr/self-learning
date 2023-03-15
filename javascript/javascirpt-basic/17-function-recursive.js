// function recursive

// factorial dari 5
// 5*4*3*2*1 = 120

function factorial(n) {
    if (n == 1) {
        console.log(1);
        return n;
    }
    console.log(n);

    return n * factorial(n - 1);
}

console.log(factorial(5));

// n * factorial(n-1)
// 5 * factorial(4)
// 4 * factorial(3)
// 3 * factorial(2)
// 2 * facotrial(1)
// console.log(1)
// 120

// do while adalah perulangan yang mirip dengan while loop. bedanya, pengecekan kondisinya ada di belakang atau di akhir.
let i = 1;

do {
    console.log(`Perulangan ke-${i}`);
    i++;
} while (i < 5);

console.log("////////////////////////////////////////");

// block code masih bisa dieksekusi, karena condition statement ada di akhir
let j = 100;
do {
    console.log(`Perulangan ke-${j}`);
    i++;
} while (j < 5);

// for loop adalah perulangan yang akan mengeksekusi block program terus menerus sampai kondisi terpenuhi

/* // struktur for loop
for (init_statement; condition_statement; post_statement) {
  // block code
} */

// init statement hanya dieksekusi 1x

/* // init, condition, dan post tidak wajib diisii.
// Jika statement tidak diisi, terutama condition statement. maka  looping akan dieksekusi tanpa henti
for (;;) {
  console.log(`for loop tanpa statement`);
} */

/* // Jika hanya diberikan condition dalam statement
let i = 1;
for (; i < 10; ) {
  console.log(`Perulangan ke-${i}`);
  i++;
} */

/* // Menambahkan init statement
for (let i = 1; i <= 10; ) {
  console.log(`Perulangan ke-${i}`);
  i++;
} */

// Memnambahkan post statement
for (let i = 1; i <= 10; i++) {
    console.log(`Perulangan ke-${i}`);
}

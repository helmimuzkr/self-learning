// For In merupakan perulangan for yang digunakan untuk mengitarisi/mendapatkan data property di object atau index di array

// For In biasa di gunakan pada Object

// Deklarasi variabel "buah" dan assign nilai Objek
const buah = {
    nama: "Apel",
    ukuran: "Sedang",
};

// for(namaBebas in namaObject){ block code }
for (property in buah) {
    console.log(`${property}: ${buah[property]}`);
}

console.log("\n");

// Pada Array biasanya menggunakan For Of
let sayur = ["bayam", "tomat", "cabai"];

// for(namaBebas in Array){ block code}
for (nama of sayur) {
    console.log(nama);
}

// For of tidak bisa digunakan pada Objek, karena Objek bukan itterable

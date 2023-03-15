// Arrow function digunakan untuk mempersingkat pembuatan function.

// Struktur
let namaVar = (params) => {
    // body
};

// Contoh
let volCube = (s) => {
    return s * s * s;
};

console.log(volCube(2));

// Contoh arrow function sebagai value dan 1 baris code
let test = (nama) => console.log(`Nama saya adalah ${nama}`);

let name = test; // function test disimpan dalam variabel name

name("Helmi");

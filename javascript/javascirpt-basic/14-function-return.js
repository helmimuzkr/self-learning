// return pada function adalah sebuah keyword untuk mengembalikan nilai atau data ke function
// function secara default tidak mempunyai value.

// biasanya retrun digunakan untuk menyelesaikan eksekusi pada function
// jadi, ketika proses mencapai "return" maka proses eksekusi function akan berhenti.

// struktur dan implementasi
function namaFunction() {
    return "mengembalikan nilai\n"; // setelah "return" lalu di ikuti value/data.
    // bisa dalam bentuk tipe data string, number, atau bahkan function lagi.
    // return bebas mengembalikan nilai apapun.
}
console.log(namaFunction());

// contoh tanpa return
function tanpaReturn(a, b) {
    hasil = a + b;
}
console.log(tanpaReturn(2, 2) + "\n"); // hasilnya akan undefined.
// karena tidak mengembalikan value yang ada di dalam function ke luar.

// contoh dengan return
function denganReturn(a, b) {
    hasil = a + b;
    return hasil;
}
console.log(denganReturn(2, 2)); // ada valuenya.

// menurut saya, return itu untuk memberikan function sebuah value.

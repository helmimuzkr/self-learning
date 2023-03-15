// Rest Parameter adalah variabel parameter yang menampung data array.
// jadi, ketika kita mengirimkan data, akan secara otomatis dikonversikan menjadi array.

/* Untuk membuat Rest Parameter ada ketentuannya
    1. Hanya boleh ada 1 rest parameter di dalam function
    2. diawali dengan 3 titik pada penamaan variabelnya
    3. Posisi Rest Parameter harus berada di akhir parameter. Tidak boleh ditengah, maupun didepan  */

// Contoh
function buah(warna, ...restParameter) {
    console.log(warna);
    console.table(restParameter);
}

buah("Kuning", "jeruk", "pisang", "lemon");

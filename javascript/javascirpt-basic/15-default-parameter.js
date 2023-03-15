// default parameter itu digunakan ketika tidak mengirimkan data ke parameter atau mengirim data undefined,
// maka akan diisi oleh default value

// contoh
function namaFunction(a, b = "defaultParameter") {
    console.log(`${a} dan ${b}`);
}

namaFunction("a"); // tidak memberikan argumen pada parameter ke dua, yaitu parameter "b".

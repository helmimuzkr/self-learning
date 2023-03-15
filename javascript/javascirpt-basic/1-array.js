// wadah array
let arrayKosong = [];

// arrow function
const hasil = (a) => {
    if (a == 1) {
        return "benar";
    }
    if (a != 1) {
        return "bukan";
    }
};

// array
let arrayNama = ["helmi", "muzakir", "muhammad"];

// menambah/mengubah array ke arrayKosong
arrayKosong.push(arrayNama, "hallo", "hallo2");

// menambah/mengubah array ke arrayKosong
arrayKosong.push(hasil(2));

console.table(arrayKosong);

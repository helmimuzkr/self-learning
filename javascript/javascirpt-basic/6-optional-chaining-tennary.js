// Optional Chaining tennary
let person = {
    helmi: {
        name: "helmi",
    },
    muzakir: {},
};

// variabel "test" = objek "person" tidak nullish lanjut, property "muzakir" tidak nullish lanjut, cek property "name".
// kalau gaada property name di dalam property "muzakir", outputnya akan undifined
let test = person?.muzakir?.name;
console.log(test);

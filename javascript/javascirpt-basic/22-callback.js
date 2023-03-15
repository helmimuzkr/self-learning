/* const arr = [
    {
        nama : 'helmi',
        npm : 444
    },
    {
        nama : 'muzakir',
        npm : 555
    }
]; */

const mhs = (nama, npm, kelas) => {
    console.log(nama, npm, kelas);
};

const testCallback = (callback) => {
    const nama = "helmi";
    const npm = 544;
    const kelas = "4ia";

    callback(nama, npm, kelas);
};

testCallback(mhs);

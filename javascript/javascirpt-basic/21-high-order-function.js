// Ambil data video dari element list dengan atribut data-duration
const videos = Array.from(document.querySelectorAll("[data-duration]"));

// filter video yang berjudul "lanjutan"
const lanjutan = videos.filter((video) => video.innerHTML.includes("Lanjutan"));

// Ambil data durasi video yang sudah difilter
const durasiVid = lanjutan.map((item) => item.dataset.duration);

// Ubah durasi menjadi detik, merubahnya menjadi float
const durasiDetik = durasiVid
    .map((durasi) => {
        // Split durasi video ":", menjadi array
        let hasilSplit = durasi
            .split(":") // XX : XX -> [XX, XX]
            .map((item) => parseFloat(item)); // Convert string to float
        // index[0] * 60 = detik + index[1]
        return hasilSplit[0] * 60 + hasilSplit[1]; // [XXXX, XXXX, ...]
    }) // Convert durasiDetik jadi float
    .map((item) => parseFloat(item));

// Menjumlahkan detik
const jmlDetik = durasiDetik.reduce((total, detik) => total + detik);

// ubah ke jam, menit, detik
const jam = Math.round(jmlDetik / 3600);
const menit = Math.round((jmlDetik % 3600) / 60);
const detik = Math.round(jmlDetik % 60);

// Simpan di DOM
const jmlVideo = document.querySelector(".jumlah-video");

jmlVideo.innerHTML = `${lanjutan.length} Video`;

const jmlDurasi = document.querySelector(".total-durasi");

jmlDurasi.innerHTML = `${jam} : ${menit} : ${detik}`;

console.log(jmlVideo);
console.log(jmlDurasi);

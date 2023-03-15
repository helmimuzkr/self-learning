/*  Function juga bisa sebagai value, artinya function dapat disimpan dalam variabel,
    dan bisa juga dikirim melalui argumen ke function lainnya. */
// Contoh
function contohDua(params) {
    console.log(params);
}
// function "contohDua" disimpan dalam variable dengan nama "contohDuaVar"
let contohDuaVar = contohDua;
// mengeksekusi function "contohDua" dengan memanggil variable "contohDuaVar"
contohDuaVar("Function Sebagai Value");

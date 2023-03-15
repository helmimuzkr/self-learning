const root = document.querySelector('#root'); // memilih element berdasarkan id root

// Cara Biasa
function caraSatu() {
  alert('halo 1');
}

// Menggunakan Argumen, dengan menambahkan parameter msg
function CaraDua(msg) {
  alert(msg);
}

console.log(CaraDua.bind(this, 'halo'));

const element = (
  <div>
    {
      // jsx inline style, menggunakan 2 curly braces.
      // curly braces pertama sebagai emmed js, kedua sebagai object
    }
    <button style={{ marginRight: '20px' }} onClick={caraSatu}>
      Tombol 1
    </button>

    {
      // .bind(this, argumen)
      // .memproduksi function baru(dengan, argumen ini)
    }
    <button onClick={CaraDua.bind(this, 'halo')}>Tombol 2</button>
  </div>
);

ReactDOM.render(element, root);

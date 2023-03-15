const root = document.querySelector('#root');

// Membuat component dengan nama App
function App() {
  // Deklarasi sebuah variabel state baru, dimana akan dinamakan "count"
  const [count, setCount] = React.useState(0);

  // Deklarassi sebuah variabel state baru, dengan nama "klik"
  const [klik, setKlik] = React.useState(false);

  // Membuat fungsi tambah
  function add() {
    setCount(count + 1);
  }

  // Membuat fungsi pengurangan
  function min() {
    setCount(count - 1);
  }

  // Fungsi Efek pada react
  // React.useEffect(params function(){}, [params condition]);
  React.useEffect(
    function () {
      console.log('Effect!');

      // Membuat fungsi destroy
      // Fungsi destroy ini dieksekusi pada saat re-rendered
      return function () {
        console.log('Destroy effect');
      };
    },
    [count] // Ketika variable state "count" berhasil di update, maka function akan dieksekusi ulang.
    // Kalau valuenya empty array maka eksekusi hanya saat komponen akan dieksekusi ketika komponen render pertama kali. contoh empty array "[]"
    // Kalau tidak ditambahkan parameter kondisi atau hanya parameter fungsi, maka fungsi akan dieksekusi setiap ada perubahan.
  );

  // Element
  return (
    <>
      <div>
        <button onClick={min}>-</button>
        <span>{count}</span>
        <button onClick={add}>+</button>
      </div>
      <div>
        <button
          onClick={function () {
            setKlik(console.log('Tombol diklik!!'));
          }}>
          Klik di sini
        </button>
      </div>
    </>
  );
}

// Render
ReactDOM.render(<App />, root);

const root = document.querySelector('#root');

// Membuat component dengan nama App
function App() {
  /*   // Cara lama
  // Inisialisasi State, value awal 0. value awal bebas, sesuai kebutuhan.
   // 
  const state = React.useState(0);

  // Inisialisasi variable untuk mengambil nilai awal. index 0
  const getCount = state[0];

  // Inisialisasi varable untuk mengupdate nilai awal, yang ada di getCount. index 1
  const setCount = state[1]; */

  // Cara simple yang biasa digunakan
  // const [0, 1] = React.useState(value);
  const [getCount, setCount] = React.useState(0);

  // Membuat fungsi tambah
  function add() {
    setCount(getCount + 1);
  }

  // Membuat fungsi pengurangan
  function min() {
    setCount(getCount - 1);
  }

  // Element
  return (
    <>
      <button onClick={min}>-</button>
      <span>{getCount}</span>
      <button onClick={add}>+</button>
    </>
  );
}

// Render
ReactDOM.render(<App />, root);

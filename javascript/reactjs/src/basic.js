// Belajar ulang
/* const root = document.querySelector('#root');
const element = document.createElement('h1');
element.innerHTML = 'TEST';

root.appendChild(element); */

// REACT
const root = document.querySelector('#root');

// Coba 1
const coba1 = React.createElement('h1', { children: 'asd' });

// Coba 2
const coba2 = React.createElement(
  'h1',
  {
    className: 'heading',
  },
  'test123'
);

// List
let list = React.createElement(
  'ul',
  {
    className: 'list',
  },
  React.createElement('h3', null, 'LIST'),
  React.createElement('li', null, 'list'),
  React.createElement('li', null, 'list'),
  React.createElement('li', null, 'list')
);

// JSX
let nama = 'HELMI';

function namaOrang(nama) {
  return <u>{nama}</u>;
}

const jsx = (
  <>
    <h2>Nama saya adalah {namaOrang(nama)}</h2>
    <h2>Nama saya adalah Mamad</h2>
    <h2>Nama saya adalah Muza</h2>
  </>
);

// Component
function Tech(props) {
  return (
    <li>
      dengan menggunakan <u>{props.name}</u>
    </li>
  );
}

const komponen = (
  <ul>
    <Tech name="ReactJS" />
    <Tech name="JSX" />
    <Tech name="asd" />
  </ul>
);

// RENDER
const element = React.createElement(
  React.Fragment,
  null,
  coba1,
  coba2,
  list,
  jsx,
  komponen
);

ReactDOM.render(element, root);

// VANILLA JS
const rootvanilla = document.querySelector('#rootvanilla');
const elementVanilla = document.createElement('h1');
// Object.Property = value;
elementVanilla.textContent = 'Created using vanilla JS!';
elementVanilla.className = 'heading-1'; // Memberikan Kelas pada element H1

rootvanilla.appendChild(elementVanilla); // Render dari javascript ke html


// React JS
// Cara 1
const rootJSX = document.querySelector('#rootREACTJS');
// Membuat React Element
// c1 adalah objek
const c1 = React.createElement('h1', {
  // Parameter Property
  children: 'Normal React JS 1', // property children bisa jadi array misalnya children: ['cara1', 'cara2'];
  className: 'heading-1',
});

// Cara 2
const rootJSX = document.querySelector('#rootREACTJS');
const c2 = React.createElement(
  'h1', // parameter element h1
  { className: 'heading-1' }, // paremeter property
  'Normal React JS 2' // child 1
);

const cara = React.createElement(
  React.Fragment, // Parameter empty element, tanpa parent node
  null, // paremeter tanpa property atau null
  c1, // child 1
  c2 // child 2
);

// Membuat List
const ul = React.createElement(
  'ul',
  { className: 'list' },
  React.createElement('h1', null, 'LIST'),
  React.createElement('li', null, 'VANILLA'),
  React.createElement('li', null, 'REACT'),
  React.createElement('li', null, 'NODE')
);

// Membuat variabel untuk render >1 Objek
const elementReact = React.createElement(React.Fragment, null, cara, ul); 

// Render dengan parameter elementReact, dan selector rootREACTJS
ReactDOM.render(elementReact, rootREACTJS);


// React JS with JSX
const JSX = document.querySelector('#rootJSX');

const ulJSX = (
  // Cuma bisa 1 parent element
  <ul>
    <h1>LIST DENGAN JSX</h1>
    <li>VANILLA</li>
    <li>REACT</li>
    <li>NODE</li>
  </ul>
);

// Try emmed Var
// Membuat function hurufCapital
function hurufCapital(params) {
  return params.toUpperCase();
}

const name = 'REACT'; // create variable nama
const kelas = 'heading-1'; // create variable kelas
const emmedVar = (
  <h1 className={kelas}>
    Belajar Embedding atau emmed pada {hurufCapital(name)}
  </h1>
);

// React Component

// Saat menggunakan function Component, nama functionya harus diawali dengan huruf Capital, 
// begitu pula pemanggilannya pun juga.

// parameter change to property
function Tech(props) {
  return (
    <li>
      Dengan <u>{props.name}</u>.
    </li>
  );
}

const component = (
  <ul>
    <h1>Mencoba Component</h1>
    <Tech name="ReactJs" />
    <Tech name="JSX" />
    <Tech name="Babel" />
  </ul>
);

// RENDERING
const element = React.createElement(
  React.Fragment,
  null,
  ulJSX,
  emmedVar,
  component
);

ReactDOM.render(element, JSX);

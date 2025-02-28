import { createEffect, createSignal, type Component } from 'solid-js';

import logo from './logo.svg';
import styles from './App.module.css';

const App: Component = () => {
  console.log("rendering component App...")

  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <img src={logo} class={styles.logo} alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          class={styles.link}
          href="https://github.com/solidjs/solid"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn Solid
        </a>
        <Button1 />
        <Button2 />
        <Button3 />
      </header>
    </div>
  );
};

function Button1() {
  console.log("rending button 1");

  const [count, setCount] = createSignal(0);

  // not tracked - only runs once during initialization.
  console.log(`count1: ${count()} - doesnt update the singal. because the signal not tracked`);

  const increament = () => setCount((prev) => prev + 1);

  return (
    <button class={styles.margin} type="button" onClick={increament}>
      Count 1
    </button>
  )
}

function Button2() {
  console.log("rending button 2")

  const [count, setCount] = createSignal(0)

  // subscribe to the count signal
  // the `createEffect` will trigger the console log every time signal `count` changes.
  createEffect(() => {
    console.log(`count2: ${count()}`);
  });

  return (
    <button class={styles.margin} type="button" onClick={() => setCount((prev) => prev + 1)}>
      Count 2
    </button>
  )
}

function Button3() {
  console.log("rending button 3")

  const [count, setCount] = createSignal(0)

  // JSX creates a tracking scope behind the scenes,
  // which allows signals to be tracked within the return statement of a component.
  //
  // Components, much like other functions, will only run once
  // This means that if a signal is accessed outside of the return statement,
  // it will run on initialization, but any updates to the signal will not trigger an update.
  //
  // the 'Count 3' value will be updated everytime signal `count` changes
  return (
    <button class={styles.margin} type="button" onClick={() => setCount((prev) => prev + 1)}>
      Count 3: {count()}
    </button >
  )
}

export default App;

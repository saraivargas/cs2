import { useState } from "react";

function App() {
  return (
    <div>
      <h1 className="text-3xl font-bold underline text-center text-red-500">
        Hello world!
      </h1>
      <MyButton />
      <div className="flex">
        <MyPizza />
        <MyPizza />
      </div>
    </div>
  );
}

function MyPizza() {
  return (
    <div>
      <img
        src={
          "https://thumbs.dreamstime.com/b/pizza-slice-above-cartoon-illustration-vector-pepperoni-paprika-mushroom-cheese-tomato-sauce-white-background-135026080.jpg"
        }
        width={300}
        height={300}
      />
      <p className="text-center pl-2">I love pizza!</p>
    </div>
  );
}

function MyButton() {
  return <button>I'm a button</button>;
}

export default App;

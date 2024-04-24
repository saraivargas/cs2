import { useState } from "react";
import Logo from "./S.png";

function App() {
  return (
    <div>
      {/* top bar */}
      <div className="flex justify-between items-center px-4 bg-slate-950 text-white border-b-2 border-red-500">
        <img src={Logo} className=" w-10 h-10" />
        <div class="relative flex">
          <input
            type="search"
            class="relative m-0 -me-0.5 block flex-auto rounded-s border border-solid border-neutral-200 bg-transparent bg-clip-padding px-3 py-[0.25rem] text-base font-normal leading-[1.6] text-surface outline-none transition duration-200 ease-in-out placeholder:text-neutral-500 focus:z-[3] focus:border-primary focus:shadow-inset focus:outline-none motion-reduce:transition-none dark:border-white/10 dark:text-white dark:placeholder:text-neutral-200 dark:autofill:shadow-autofill dark:focus:border-primary"
            placeholder="Search"
            aria-label="Search"
            id="exampleFormControlInput3"
            aria-describedby="button-addon3"
          />
          <button
            class="z-[2] inline-block rounded-e border-2 border-primary px-6 pb-[6px] pt-2 text-xs font-medium uppercase leading-normal text-primary transition duration-150 ease-in-out hover:border-primary-accent-300 hover:bg-primary-50/50 hover:text-primary-accent-300 focus:border-primary-600 focus:bg-primary-50/50 focus:text-primary-600 focus:outline-none focus:ring-0 active:border-primary-700 active:text-primary-700 dark:text-primary-500 dark:hover:bg-blue-950 dark:focus:bg-blue-950"
            data-twe-ripple-init
            data-twe-ripple-color="white"
            type="button"
            id="button-addon3"
          >
            Search
          </button>
        </div>
        <div class="grid grid-cols-4 gap-4 py-2 w-2/5">
          <div class="h-10 bg-slate-700 rounded-md flex justify-center items-center">Home</div>
          <div class="h-10 bg-slate-700 rounded-md flex justify-center items-center">Favorites</div>
          <div class="h-10 bg-slate-700 rounded-md flex justify-center items-center">Profile</div>
          <div class="h-10 bg-slate-700 rounded-md flex justify-center items-center">Settings</div>
        </div>
      </div>

      <h1 className="text-3xl font-bold underline text-center text-red-500">Hello world!</h1>
    </div>
  );
}

export default App;

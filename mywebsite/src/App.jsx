import { useState } from "react";

function App() {
  return (
    <div>
      {/* top bar */}
      <div className="flex justify-between items-center px-4 bg-slate-950 text-white border-b-2 border-red-500">
        <div className="w-25">Sarai</div>
        <div className="w-60">Search Bar</div>
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

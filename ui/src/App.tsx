import './App.css';

function App() {
  return (
    <div className="bg-orange-50 w-screen h-screen">
      <header className="h-16 shadow bg-teal-800 flex items-center justify-between px-4 py-2 text-zinc-50 font-semibold">
        <div className='flex items-end'>
          <h3 className="text-xl mr-4 font-bold">LineUPs</h3>
          <div className="flex items-end text-sm">
            <p className="nav-link">Page1</p>
            <p className="nav-link">Page2</p>
            <p className="nav-link">Page3</p>
          </div>
        </div>
      </header>
      <body className="p-4 flex justify-around">
        <button className="btn btn-norm transition">Normal-Btn</button>
        <button className="btn btn-light transition">Light-Btn</button>
        <button className="btn btn-action transition">Action-Btn</button>

      </body>
    </div>
  );
}

export default App;

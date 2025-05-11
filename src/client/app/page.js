// "use client";
// import { useEffect, useState } from "react";
// import ForceGraph from "./graph";

// export default function Home() {
//   const [message, setMessage] = useState("Loading...");

//   useEffect(() => {
//     fetch("http://localhost:8080/api/hello")
//       .then((res) => res.json())
//       .then((data) => setMessage(data.text))
//       .catch(() => setMessage("Failed to fetch"));
//   }, []);

//   return (
//     <main>
//       <h1>{message}</h1>
//       <ForceGraph />
//     </main>
//   );
// }

"use client";

import RecipeForm from "../components/RecipeForm"; // Import RecipeForm component

export default function Home() {
  return (
    <main className="min-h-screen bg-gray-100 px-8 p-6">
      {/* Header dengan gambar */}
      <header className="relative mb-8">
        <img
          src="/images/header.png"  // Gambar diambil dari folder public/images/
          alt="Header Image"
          className="w-full h-64 object-cover rounded-t-lg"  // Menyesuaikan ukuran dan styling gambar
        />
      </header>

      {/* Judul di bawah gambar */}
      <div className="text-center mb-8">
        <h2 className="text-4xl font-bold" style={{color : "#FAA620"}}>Little Alchemy 2 Recipe Finder</h2>
      </div>

      {/* Form untuk pencarian resep */}
      <div className="max-w-2xl mx-auto bg-white p-6 rounded shadow">
        <RecipeForm />
      </div>
    </main>
  );
}




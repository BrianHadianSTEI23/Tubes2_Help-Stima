"use client";
import RecipeForm from "../../components/RecipeForm"; // Import RecipeForm component

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
        <h2 className="text-4xl font-bold blinking-text" style={{color : "#FAA620"}}>
          Little Alchemy 2 Recipe Finder
        </h2>
      </div>

      {/* Form untuk pencarian resep */}
      <div className="w-full mx-auto bg-white p-6 rounded shadow">
        <RecipeForm />
      </div>
    </main>
  );
}




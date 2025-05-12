"use client";

import { useState } from "react";
import RecipeTree from "./RecipeTree";
import axios from "axios"; // Import axios untuk melakukan HTTP request
import templateTree from "./template.json"; // import langsung dari JSON file

function normalizeTreeData(node) {
  if (!node) return null;

  return {
    name: node.Name,
    children: node.Children?.map(normalizeTreeData) || []
  };
}

export default function RecipeForm() {
  const [target, setTarget] = useState("");
  const [algorithm, setAlgorithm] = useState("1");
  const [mode, setMode] = useState("1");
  const [maxRecipes, setMaxRecipes] = useState(1);
  const [result, setResult] = useState(null);

  const handleBack = () => {
    window.history.back(); // Navigasi ke halaman sebelumnya
  };

  const handleRefresh = () => {
    window.location.reload(); // Me-refresh halaman saat ini
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Membuat payload untuk dikirim
    const payload = {
      Target : target,
      Algorithm: parseInt(algorithm),
      Mode: parseInt(mode),
      MaxRecipes: parseInt(maxRecipes),
    };

    // Console log untuk memastikan payload sudah disimpan dengan benar
    console.log("Payload yang dikirim:", payload);

    try {
      // Mengirimkan data ke server menggunakan axios
      const response = await axios.post("http://localhost:8080/api/post-recipe", payload);
      const { Data: val, NumOfRecipe } = response.data;

      console.log(val)
      // Menyimpan hasil yang diterima dari server ke dalam result
      setResult({
        status: "success",
        message: "Data berhasil diterima!",
        tree:val, // Asumsi response dari server berisi data pohon
      });
    } catch (err) {
      console.error("Gagal fetch ke backend:", err.message);
      setResult({ status: "error", message: "Terjadi kesalahan saat menghubungi server." });
    }
  };

  return (
    <div className="text-black">
      <h2 className="text-2xl font-bold mb-4">Pencarian Resep Elemen</h2>

      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block font-semibold mb-1">Elemen Tujuan</label>
          <input
            type="text"
            value={target}
            onChange={(e) => setTarget(e.target.value)}
            placeholder="Contoh: Brick"
            className="border border-gray-400 rounded px-3 py-2 w-full"
            required
          />
        </div>

        <div>
          <label className="block font-semibold mb-1">Algoritma Pencarian</label>
          <select
            value={algorithm}
            onChange={(e) => setAlgorithm(e.target.value)}
            className="border border-gray-400 rounded px-3 py-2 w-full"
          >
            <option value="1">DFS (Depth-First Search)</option>
            <option value="2">BFS (Breadth-First Search)</option>
            <option value="3">Bidirectional (Bonus)</option>
          </select>
        </div>

        <div>
          <label className="block font-semibold mb-1">Mode Pencarian</label>
          <select
            value={mode}
            onChange={(e) => setMode(e.target.value)}
            className="border border-gray-400 rounded px-3 py-2 w-full"
          >
            <option value="1">Shortest Path</option>
            <option value="2">Multiple Recipe</option>
          </select>
        </div>

        {mode === "2" && (
          <div>
            <label className="block font-semibold mb-1">Jumlah Maksimal Recipe</label>
            <input
              type="number"
              value={maxRecipes}
              onChange={(e) => setMaxRecipes(e.target.value)}
              className="border border-gray-400 rounded px-3 py-2 w-full"
              min="1"
            />
          </div>
        )}

        <div className="flex justify-between mt-4">
          <button type="submit" className="bg-[#260027] text-white font-semibold px-4 py-2 rounded">
            Cari
          </button>
          <div className="flex space-x-4">
            <button
              type="button"
              onClick={handleRefresh}
              className="bg-gray-400 text-[#260027] font-semibold px-4 py-2 rounded"
            >
              Refresh
            </button>
            <button
              type="button"
              onClick={handleBack}
              className="bg-[#FAA620] text-[#260027] font-semibold px-4 py-2 rounded"
            >
              Kembali
            </button>
          </div>
        </div>
      </form>

      {result && (
        <div className="mt-10">
          {/* Menampilkan visualisasi tree jika hasil tersedia */}
          {result.tree && (
            <>
              <h4 className="text-lg font-semibold mt-6 mb-2">Visualisasi Tree</h4>
              <RecipeTree data={normalizeTreeData(result.tree)} />
            </>
          )}
        </div>
      )}
    </div>
  );
}

'use client';
import { useState } from 'react';
import Link from 'next/link';

export default function Home() {
  // State untuk mengatur apakah tombol "On/Off" aktif
  const [isOn, setIsOn] = useState(false);

  const handleToggle = () => {
    setIsOn(!isOn); // Toggle state On/Off
  };

  return (
    <main className="min-h-screen bg-[#40042c] p-6">
      {/* Header dengan gambar */}
      <header className="relative mb-8">
        <div className="text-center mb-8">
          <h2 className="text-4xl blinking-text font-bold text-[#FAA620]">Little Alchemy 2</h2>
          <p className="text-xl fading-text text-slate-50 mt-4">Gabungkan elemen untuk menemukan lebih dari 700 item baru!</p>
        </div>
      </header>

      {/* Deskripsi tentang game */}
      <section className="max-w-2xl mx-auto bg-white p-6 rounded-xl shadow">
        <h3 className="text-xl font-semibold text-center text-[#40042c]">Tentang Game</h3>
        <p className="mt-4 text-gray-800">
          Little Alchemy 2 adalah game crafting edukatif yang dikembangkan oleh Recloak. Dalam game ini, pemain memulai dengan empat elemen dasar: udara, api, air, dan tanah. Melalui eksperimen dan logika, pemain dapat menggabungkan elemen-elemen tersebut untuk menciptakan lebih dari 700 item baru, mulai dari benda sehari-hari hingga fenomena alam dan makhluk mitologi.
        </p>

        <h4 className="text-xl font-semibold mt-6 text-center text-[#40042c]">Fitur Utama</h4>
        <ul className="list-disc ml-6 mt-2 text-gray-800">
          <li>Eksperimen Tanpa Batas: Gabungkan dua elemen untuk menemukan item baru.</li>
          <li>Desain Minimalis: Antarmuka yang sederhana dan estetis, cocok untuk sesi bermain singkat maupun panjang.</li>
          <li>Encyclopedia In-Game: Setiap item yang ditemukan dilengkapi dengan deskripsi lucu dan informatif.</li>
          <li>Mode Offline: Dapat dimainkan tanpa koneksi internet.</li>
        </ul>


        <div className="flex justify-center mt-8 space-x-4">
          {/* Tombol Mainkan Sekarang */}
            <div className="flex items-center w-full justify-center">
                <Link
                    href={isOn ? "/RecipePage" : "https://littlealchemy2.com"}
                    className="text-white px-8 py-4 font-semibold rounded-lg 
                                bg-gradient-to-r from-[#FAA620] to-[#FF7C00] 
                                hover:from-[#FF7C00] hover:to-[#FAA620] 
                                shadow-md hover:shadow-xl transition-all duration-300 transform hover:scale-105"
                    style={{ color: "#40042c" }}
                >
                    Mainkan Sekarang
                </Link>
            </div>

            <div className="flex flex-col items-center w-full">
                {/* Slider Switch */}
                <label htmlFor="toggle" className="flex items-center cursor-pointer">
                    <span className="mr-3 text-white font-semibold">{isOn ? 'On' : 'Off'}</span>
                    <div
                    className={`relative w-16 h-6 ${isOn ? 'bg-[#FF7C00]' : 'bg-gray-300'} rounded-full hover:bg-[#FF7C00] hover:scale-110 transition-all ease-in-out duration-300`}
                    onClick={handleToggle}
                    >
                    <div
                        className={`absolute w-6 h-6 bg-white rounded-full transition-transform duration-300 ease-in-out transform ${
                        isOn ? 'translate-x-10' : 'translate-x-0'
                        }`}
                    ></div>
                    </div>
                </label>

                <div className="flex flex-col text-center mt-4">
                    <p className="text-sm text-gray-700">
                        {isOn ? 'Klik tombol untuk melihat detail resep.' : 'Mainkan game untuk mulai menggabungkan'}
                    </p>
                    <p className="text-sm text-gray-700">
                        {isOn ? 'dan mencoba eksperimen lebih lanjut.' : 'elemen dan menemukan item baru!'}
                    </p>
                </div>
            </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-gray-800 text-white text-center py-4 mt-12">
        <p>&copy; 2025 Little Alchemy 2 | All Rights Reserved</p>
      </footer>
    </main>
  );
}
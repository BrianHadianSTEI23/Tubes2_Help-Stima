# This is a placeholder


## Penjelasan singkat algoritma DFS dan BFS yang diimplementasikan
### Algoritma Depth First Search (DFS)
DFS adalah algoritma pencarian graf yang mengeksplorasi setiap simpul secara mendalam sebelum mundur dan mencari jalur lainnya. Algoritma ini menggunakan struktur data stack untuk menyimpan simpul yang sedang diproses. Dalam penerapannya, DFS akan mengunjungi simpul pertama, lalu berlanjut ke simpul tetangga yang belum dikunjungi sampai kedalaman maksimal dicapai, kemudian mundur (backtrack) ke simpul sebelumnya untuk mengeksplorasi jalur lainnya.
Langkah-langkah DFS:
1. Masukkan simpul awal ke dalam stack.
2. Selama stack tidak kosong:
3. Ambil simpul teratas.
4. Jika simpul tersebut adalah simpul yang dicari, berhenti.
5. Tandai simpul sebagai dikunjungi.
6. Masukkan semua tetangga yang belum dikunjungi ke dalam stack.
DFS tidak selalu menjamin jalur terpendek, tetapi sangat efisien untuk menemukan solusi dalam graf yang dalam.

### Algoritma Breadth First Search (BFS)
BFS adalah algoritma pencarian graf yang mengeksplorasi semua simpul pada satu level terlebih dahulu sebelum melanjutkan ke level berikutnya. Algoritma ini menggunakan struktur data queue (FIFO) untuk menyimpan simpul yang sedang diproses. BFS ideal untuk mencari jalur terpendek dalam graf yang tidak berbobot.
Langkah-langkah BFS:
1. Masukkan simpul awal ke dalam queue.
2. Selama queue tidak kosong:
3. Ambil simpul terdepan dari queue.
4. Jika simpul tersebut adalah simpul yang dicari, berhenti.
5. Tandai simpul sebagai dikunjungi.
6. Masukkan semua tetangga yang belum dikunjungi ke dalam queue.
BFS memastikan bahwa jalur terpendek ditemukan pada graf yang tidak berbobot, karena mengeksplorasi simpul lebih luas sebelum menyelam lebih dalam.

## Requirement Program dan Instalasi
## Instalasi program
1. Install Go dan Node Js sebelum menjalankan backend dan frontend
2. Jalankan perintah berikut untuk mengunduh dependencies frontend
   ```bash
   npm install
   ```

## Cara Menjalankan Program
1. Masuk ke folder src/backend untuk menjalankan backend
  ```bash
  go run main.go
  ```
2. Masuk ke folder src/frontend untuk menjalankan frontend
  ```bash
  npm run dev
  ```

# Deskripsi Program
Program ini dapat membantu penentuan rute resep pembuatan setiap elemen dalam permainan Little Alchemist 2.

## Penulis Progam
- Brian Albar Hadian 13523048
- M. Izzat Jundy 13523092
- Andrew Isra Putra DB 13523110

'use client';

import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css'; // Mengimpor CSS Swiper

export default function Carousel() {
  return (
    <section className="max-w-4xl mx-auto mb-8">
      <Swiper
        spaceBetween={30} // Jarak antar slide
        slidesPerView={1} // Menampilkan 2 gambar dalam satu slide
        navigation
        loop={true} // Gambar akan bergulir terus
        autoplay={{ delay: 2000 }} // Waktu peralihan gambar dalam 3 detik
      >
        <SwiperSlide>
          <div className="flex">
            <img
              src="/images/image1.jpg" // Ganti dengan gambar yang ada di folder public/images/
              alt="Image 1"
              className="w-1/2 h-64 object-cover rounded-xl"
            />
            <img
              src="/images/image2.jpg" // Ganti dengan gambar yang ada di folder public/images/
              alt="Image 2"
              className="w-1/2 h-64 object-cover rounded-xl"
            />
          </div>
        </SwiperSlide>
        <SwiperSlide>
          <div className="flex">
            <img
              src="/images/image3.jpg" // Ganti dengan gambar yang ada di folder public/images/
              alt="Image 3"
              className="w-1/2 h-64 object-cover rounded-xl"
            />
            <img
              src="/images/image4.jpg" // Ganti dengan gambar yang ada di folder public/images/
              alt="Image 4"
              className="w-1/2 h-64 object-cover rounded-xl"
            />
          </div>
        </SwiperSlide>
        <SwiperSlide>
          <div className="flex">
            <img
              src="/images/image5.jpg" // Ganti dengan gambar yang ada di folder public/images/
              alt="Image 3"
              className="w-1/2 h-64 object-cover rounded-xl"
            />
            <img
              src="/images/image6.jpg" // Ganti dengan gambar yang ada di folder public/images/
              alt="Image 4"
              className="w-1/2 h-64 object-cover rounded-xl"
            />
          </div>
        </SwiperSlide>
        {/* Tambahkan lebih banyak slide jika diperlukan */}
      </Swiper>
    </section>
  );
}

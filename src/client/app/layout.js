import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

// Mengonfigurasi font Geist dan Geist_Mono
const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

// Metadata yang akan digunakan oleh aplikasi
export const metadata = {
  title: "Finder - Little Alchemy 2",
  description: "find your recipe",
  icons: {
    icon: "icon/logo_web.png",
    shortcut: "/icon/logo_web.png",
    apple: "public/icon/logo_web.png",
  },
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        {children}
      </body>
    </html>
  );
}

"use client"
import { useEffect, useState } from 'react';

export default function Home() {
  const [message, setMessage] = useState('Loading...');

  useEffect(() => {
    fetch('http://localhost:8080/api/hello')
      .then((res) => res.json())
      .then((data) => setMessage(data.text))
      .catch(() => setMessage('Failed to fetch'));
  }, []);

  return <h1>{message}</h1>;
}

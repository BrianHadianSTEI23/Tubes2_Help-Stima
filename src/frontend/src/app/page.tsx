"use client"
import { useEffect, useState } from 'react';

export default function Home() {
  const [response, setResponse] = useState('Loading');

  const handleSubmit = async () => {
    const res = await fetch('http://localhost:8080/api/post-recipe', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        target: 'Iron',
        algorithm: 1,
        mode: 2,
        maxRecipes: 5
      })
    })

    const data = await res.json()
    setResponse(data.message)
  }

    return <main>
      <h1>{response}</h1>
    </main>;
}

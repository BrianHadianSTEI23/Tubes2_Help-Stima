// 'use client';
// import React from 'react';
// import ReactECharts from 'echarts-for-react';

// export default function ForceGraph() {
//   const option = {
//     title: {
//       text: 'Force Graph Example',
//     },
//     tooltip: {},
//     series: [
//       {
//         type: 'graph',
//         layout: 'force',
//         roam: true,
//         label: {
//           show: true,
//           position: 'right',
//         },
//         force: {
//           repulsion: 200,
//           edgeLength: [50, 100],
//         },
//         data: [
//           { id: 'A', name: 'A', symbolSize: 50 },
//           { id: 'B', name: 'B', symbolSize: 50 },
//           { id: 'combineAB', name: '', symbolSize: 10, itemStyle: { opacity: 0 } }, // hidden node
//           { id: 'C', name: 'C', symbolSize: 60 },
//         ],
//         links: [
//           { source: 'A', target: 'combineAB' },
//           { source: 'B', target: 'combineAB' },
//           { source: 'combineAB', target: 'C', symbol: ['none', 'arrow'] },
//         ],
//       },
//     ],
//   };

//   return <ReactECharts option={option} style={{ height: '600px', width: '100%' }} />;
// }

"use client";
import { useEffect, useState } from "react";
import axios from "axios";
import ReactECharts from "echarts-for-react";

export default function ForceGraph({graphData}) {
  const [graphData, setGraphData] = useState(null);
  const [message, setMessage] = useState("Loading...");

  // Mengambil data dari API saat komponen dimuat
  useEffect(() => {
    // Mengirimkan request menggunakan Axios
    axios
      .get("http://localhost:8080/api/graphdata")  // Endpoint API yang mengirimkan data graf
      .then((response) => {
        setGraphData(response.data);  // Menyimpan data yang diterima dari API
        setMessage("Graph data loaded");
      })
      .catch((error) => {
        console.error("Error fetching graph data:", error);
        setMessage("Failed to load graph data");
      });
  }, []); // Hanya dijalankan sekali saat komponen pertama kali dimuat

  // Opsi untuk menampilkan graf menggunakan ECharts
  const chartOptions = {
    tooltip: {},
    animation: true,
    series: [
      {
        type: "graph",
        layout: "force",
        data: graphData ? graphData.nodes : [],
        links: graphData ? graphData.links : [],
        roam: true,
        label: {
          show: true,
          position: "right",
          formatter: "{b}",
        },
        force: {
          repulsion: 100,
          edgeLength: [50, 200],
        },
      },
    ],
  };

  if (!graphData) {
    return <p>{message}</p>;  // Menampilkan pesan loading jika data belum ada
  }

  return (
    <div>
      <h2>Force Graph</h2>
      <ReactECharts option={chartOptions} style={{ height: "500px" }} />
    </div>
  );
}

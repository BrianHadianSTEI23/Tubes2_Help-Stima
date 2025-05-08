'use client';
import React from 'react';
import ReactECharts from 'echarts-for-react';

export default function ForceGraph() {
  const option = {
    title: {
      text: 'Force Graph Example',
    },
    tooltip: {},
    series: [
      {
        type: 'graph',
        layout: 'force',
        roam: true,
        label: {
          show: true,
          position: 'right',
        },
        force: {
          repulsion: 200,
          edgeLength: [50, 100],
        },
        data: [
          { id: 'A', name: 'A', symbolSize: 50 },
          { id: 'B', name: 'B', symbolSize: 50 },
          { id: 'combineAB', name: '', symbolSize: 10, itemStyle: { opacity: 0 } }, // hidden node
          { id: 'C', name: 'C', symbolSize: 60 },
        ],
        links: [
          { source: 'A', target: 'combineAB' },
          { source: 'B', target: 'combineAB' },
          { source: 'combineAB', target: 'C', symbol: ['none', 'arrow'] },
        ],
      },
    ],
  };

  return <ReactECharts option={option} style={{ height: '600px', width: '100%' }} />;
}

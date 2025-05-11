"use client";

import dynamic from "next/dynamic";
import { useEffect, useRef, useState } from "react";

const Tree = dynamic(() => import("react-d3-tree").then((mod) => mod.Tree), {
  ssr: false,
});

export default function RecipeTree({ data }) {
  const treeRef = useRef(null);
  const [dimensions, setDimensions] = useState({ width: 600, height: 600 });

  useEffect(() => {
    if (treeRef.current) {
      const { width, height } = treeRef.current.getBoundingClientRect();
      setDimensions({ width, height });
    }
  }, []);

  const getColor = (name) => {
    const lower = name.toLowerCase();
    if (lower === "water") return "#3B82F6";
    if (lower === "earth") return "#10B981";
    if (lower === "fire") return "#B22200";
    if (lower === "air") return "#6B7280";
    return "#E0E0E0";
  };

  const renderCustomNode = ({ nodeDatum }) => {
    const name = nodeDatum.name;
    const paddingX = 20;
    const charWidth = 8;
    const textWidth = name.length * charWidth;
    const boxWidth = textWidth + paddingX;
    const boxHeight = 40;

    const fillColor = getColor(name);

    return (
      <g transform="rotate(180)">
        <rect
          x={-boxWidth / 2}
          y={-boxHeight / 2}
          width={boxWidth}
          height={boxHeight}
          fill={fillColor}
          stroke="#191919"
          strokeWidth="1.5"
          rx={8}
        />
       <text
            fill="#000000"
            x={0}
            y={4} // teks lebih jauh di bawah node
            textAnchor="middle"
            style={{
                fontSize: "14px",
                fontFamily: "Roboto",
                fontWeight: "normal",

            }}
            >
            {name}
        </text>
      </g>
    );
  };

  return (
    <div
      ref={treeRef}
      style={{
        width: "100%",
        height: "600px",
        transform: "rotate(180deg)",
        background: "#F9FAFB",
        padding: "1rem",
        borderRadius: "8px",
      }}
    >
      <Tree
        data={data}
        orientation="vertical"
        translate={{ x: dimensions.width / 2, y: dimensions.height - 50 }}
        pathFunc="elbow"
        depthFactor={150}
        collapsible={false}
        nodeSize={{ x: 200, y: 100 }}
        renderCustomNodeElement={renderCustomNode}
      />
    </div>
  );
}

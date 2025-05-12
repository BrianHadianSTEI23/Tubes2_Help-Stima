"use client";

import dynamic from "next/dynamic";
import { useEffect, useRef, useState } from "react";

// Dynamically import the Tree component from react-d3-tree
const Tree = dynamic(() => import("react-d3-tree").then((mod) => mod.Tree), {
  ssr: false,
});

export default function RecipeTree({ data }) {
  const treeRef = useRef(null);
  const [dimensions, setDimensions] = useState({ width: 600, height: 600 });

  // Calculate tree dimensions on component mount
  useEffect(() => {
    if (treeRef.current) {
      const { width, height } = treeRef.current.getBoundingClientRect();
      setDimensions({ width, height });
    }
  }, []);

  // Unified color for all nodes
  const nodeColor = "#B0B0B0";  // Set to a pleasant green color

  // Memoize the custom node renderer to prevent unnecessary re-renders
  const renderCustomNode = ({ nodeDatum }) => {
    const name = String(nodeDatum.name);
    const paddingX = 20;
    const charWidth = 8;
    const textWidth = name.length * charWidth;
    const boxWidth = textWidth + paddingX;
    const boxHeight = 40;

    // Use the same color for all nodes
    const fillColor = nodeColor;

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
          fill="#000000" // Set the text color to white for better contrast
          x={0}
          y={5} // Move the text slightly lower inside the node
          textAnchor="middle"
          stroke = "none"
          style={{
            fontSize: "16px", // Slightly larger text for better readability
            fontFamily: "Roboto, sans-serif",
            fontWeight: "800",  // A bit of boldness for clarity
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
        padding: "1.5rem",  // Increased padding for better spacing
        borderRadius: "8px",
        boxSizing: "border-box",
        overflow: "hidden",
      }}
    >
      <Tree
        data={data}
        orientation="vertical"
        translate={{ x: dimensions.width/2, y: dimensions.height/2 }}
        pathFunc="elbow"
        depthFactor={120}
        collapsible={false}
        nodeSize={{ x: 200, y: 100 }}
        renderCustomNodeElement={renderCustomNode}
      />
    </div>
  );
}
